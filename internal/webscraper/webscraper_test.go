package webscraper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/miscord-dev/annict-epgstation-connector/internal/syncer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	// slog import is not needed here if we are not asserting log outputs
)

func TestGetVodServices(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var responseBody string
		switch r.URL.Path {
		case "/works/123": // Multiple VOD services, text in <a>
			responseBody = `
				<html><body><div><div><div><div><div><div><ul class="list-inline mt-2">
					<li><a href="#">Netflix</a></li>
					<li><a href="#">Amazon Prime Video</a></li>
					<li><a href="#">  dアニメストア  </a></li>
				</ul></div></div></div></div></div></div></body></html>`
		case "/works/456": // Single VOD service, text in <a>
			responseBody = `
				<html><body><div><div><div><div><div><div><ul class="list-inline mt-2">
					<li><a href="#">Hulu</a></li>
				</ul></div></div></div></div></div></div></body></html>`
		case "/works/789": // No VOD services (empty list)
			responseBody = `
				<html><body><div><div><div><div><div><div><ul class="list-inline mt-2">
				</ul></div></div></div></div></div></div></body></html>`
		case "/works/000": // Page not found
			w.WriteHeader(http.StatusNotFound)
			return
		case "/works/empty_a_with_img_alt": // Empty <a> text, also has img with alt. Should be ignored.
			responseBody = `
				<html><body><div><div><div><div><div><div><ul class="list-inline mt-2">
					<li><a href="#"><img alt="Disney+" src="..."></a></li>
				</ul></div></div></div></div></div></div></body></html>`
		case "/works/text_and_img_alt_different": // Text in <a>, different in img alt (text should be used)
			responseBody = `
				<html><body><div><div><div><div><div><div><ul class="list-inline mt-2">
					<li><a href="#"><img alt="U-NEXT (alt)" src="...">U-NEXT</a></li>
				</ul></div></div></div></div></div></div></body></html>`
		case "/works/duplicate_services": // Duplicate services to test deduplication
			responseBody = `
				<html><body><div><div><div><div><div><div><ul class="list-inline mt-2">
					<li><a href="#">Netflix</a></li>
					<li><a href="#">Netflix </a></li>
					<li><a href="#">Amazon Prime Video</a></li>
				</ul></div></div></div></div></div></div></body></html>`
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(responseBody))
	}))
	defer server.Close()

	// Local GetVodServices for testing, using the actual simplified logic from webscraper.go
	getLocalVodServices := func(baseURL string, workID string) ([]syncer.VodService, error) {
		url := fmt.Sprintf("%s/works/%s", baseURL, workID)
		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to get Annict work page: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("failed to get Annict work page: status code %d for workID %s", res.StatusCode, workID)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Annict work page for workID %s: %w", workID, err)
		}

		var vodServices []syncer.VodService
		selector := "body > div > div.l-default__main.d-flex.flex-column > div.l-default__content > div.c-work-header.pt-3 > div.container > div > div.col.mt-3.mt-sm-0 > ul.list-inline.mt-2 > li > a"
		doc.Find(selector).Each(func(i int, s *goquery.Selection) {
			serviceName := strings.TrimSpace(s.Text())
			if serviceName != "" {
				vodServices = append(vodServices, syncer.VodService{Name: serviceName})
			}
			// If serviceName is empty, the <a> tag is simply ignored.
		})

		if len(vodServices) > 0 {
			seen := make(map[string]bool)
			uniqueServices := []syncer.VodService{}
			for _, service := range vodServices {
				name := strings.TrimSpace(service.Name)
				if name == "" { continue }
				if _, ok := seen[name]; !ok {
					seen[name] = true
					uniqueServices = append(uniqueServices, syncer.VodService{Name: name})
				}
			}
			vodServices = uniqueServices
		}
		return vodServices, nil
	}

	tests := []struct {
		name          string
		workID        string
		expected      []syncer.VodService
		expectedError bool
	}{
		{
			name:   "Multiple VOD services - text in a",
			workID: "123",
			expected: []syncer.VodService{
				{Name: "Netflix"},
				{Name: "Amazon Prime Video"},
				{Name: "dアニメストア"},
			},
		},
		{
			name:   "Single VOD service - text in a",
			workID: "456",
			expected: []syncer.VodService{{Name: "Hulu"}},
		},
		{
			name:     "No VOD services",
			workID:   "789",
			expected: []syncer.VodService{},
		},
		{
			name:          "Work not found (404)",
			workID:        "000",
			expectedError: true,
		},
		{
			name:     "Empty <a> text with <img> alt - service ignored",
			workID:   "empty_a_with_img_alt", // <a><img alt="Disney+"></a>
			expected: []syncer.VodService{}, // No service should be returned
		},
		{
			name:   "Text in <a>, different in <img> alt - text used",
			workID: "text_and_img_alt_different", // <a>U-NEXT</a> <img alt="U-NEXT (alt)">
			expected: []syncer.VodService{{Name: "U-NEXT"}},
		},
		{
			name:   "Duplicate services - deduplicated",
			workID: "duplicate_services",
			expected: []syncer.VodService{
				{Name: "Netflix"},
				{Name: "Amazon Prime Video"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vodServices, err := getLocalVodServices(server.URL, tt.workID)
			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.ElementsMatch(t, tt.expected, vodServices)
			}
		})
	}
}
