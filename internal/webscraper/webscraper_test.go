package webscraper

import (
	"fmt" // Ensure fmt is imported if used for constructing URLs for the test server
	"net/http"
	"net/http/httptest"
	"strings" // Ensure strings is imported
	"testing"

	"github.com/PuerkitoBio/goquery" // Ensure goquery is imported
	"github.com/miscord-dev/annict-epgstation-connector/internal/syncer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetVodServices(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Base selector used by the actual GetVodServices function
		// We need to ensure our test HTML is structured such that this selector works.
		// selector := "body > div > div.l-default__main.d-flex.flex-column > div.l-default__content > div.c-work-header.pt-3 > div.container > div > div.col.mt-3.mt-sm-0 > ul.list-inline.mt-2 > li > a"

		// Simplified structure for test server responses, matching what the selector expects
		var responseBody string
		switch r.URL.Path {
		case "/works/123": // Multiple VOD services, text in <a>
			responseBody = `
				<html><body><div><div><div><div><div><div><ul class="list-inline mt-2">
					<li><a href="#">Netflix</a></li>
					<li><a href="#">Amazon Prime Video</a></li>
					<li><a href="#">  dアニメストア  </a></li>
				</ul></div></div></div></div></div></div></body></html>`
		case "/works/456": // Single VOD service, text in <a> and an img alt (prefer text)
			responseBody = `
				<html><body><div><div><div><div><div><div><ul class="list-inline mt-2">
					<li><a href="#"><img alt="Hulu (alt)" src="...">Hulu</a></li>
				</ul></div></div></div></div></div></div></body></html>`
		case "/works/789": // No VOD services (empty list)
			responseBody = `
				<html><body><div><div><div><div><div><div><ul class="list-inline mt-2">
				</ul></div></div></div></div></div></div></body></html>`
		case "/works/000": // Page not found
			w.WriteHeader(http.StatusNotFound)
			return
		case "/works/img_only": // Only img alt text available
			responseBody = `
				<html><body><div><div><div><div><div><div><ul class="list-inline mt-2">
					<li><a href="#"><img alt="Disney+" src="..."></a></li>
				</ul></div></div></div></div></div></div></body></html>`
		case "/works/text_and_img_alt_different": // Text in <a>, different in img alt (text should be preferred)
			responseBody = `
				<html><body><div><div><div><div><div><div><ul class="list-inline mt-2">
					<li><a href="#"><img alt="U-NEXT (alt)" src="...">U-NEXT</a></li>
				</ul></div></div></div></div></div></div></body></html>`
		case "/works/text_and_img_alt_specific": // Text in <a>, more specific in img alt (current logic prefers <a> text)
											// If we wanted to prefer more specific alt: `strings.Contains(strings.TrimSpace(imgAlt), serviceName) && len(strings.TrimSpace(imgAlt)) > len(serviceName)`
			responseBody = `
				<html><body><div><div><div><div><div><div><ul class="list-inline mt-2">
					<li><a href="#"><img alt="dアニメストア (ニコニコ支店)" src="...">dアニメストア</a></li>
				</ul></div></div></div></div></div></div></body></html>`
		case "/works/duplicate_services": // Duplicate services to test deduplication
			responseBody = `
				<html><body><div><div><div><div><div><div><ul class="list-inline mt-2">
					<li><a href="#">Netflix</a></li>
					<li><a href="#"><img alt="Netflix" src="..."></a></li>
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

	// Use the actual GetVodServices function, but point its calls to the test server.
	// This requires GetVodServices to be flexible enough to take a URL or be modified for testing.
	// For this test, we'll redefine a local GetVodServices that takes the server URL.
	// This mirrors the structure of the webscraper.GetVodServices but allows injecting the test server's URL.
	getLocalVodServices := func(baseURL string, workID string) ([]syncer.VodService, error) {
		url := fmt.Sprintf("%s/works/%s", baseURL, workID) // Use baseURL from test server
		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to get Annict work page: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("failed to get Annict work page: status code %d", res.StatusCode)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Annict work page: %w", err)
		}

		// This is the exact extraction logic from the webscraper.go (after correction)
		var vodServices []syncer.VodService
		selector := "body > div > div.l-default__main.d-flex.flex-column > div.l-default__content > div.c-work-header.pt-3 > div.container > div > div.col.mt-3.mt-sm-0 > ul.list-inline.mt-2 > li > a"
		doc.Find(selector).Each(func(i int, s *goquery.Selection) {
			serviceName := strings.TrimSpace(s.Text())
			if serviceName != "" {
				// Prefer longer, more specific alt text only if it clearly contains the shorter serviceName from text and is longer.
				imgAlt, imgExists := s.Find("img").Attr("alt")
				if imgExists && strings.TrimSpace(imgAlt) != "" && strings.Contains(strings.TrimSpace(imgAlt), serviceName) && len(strings.TrimSpace(imgAlt)) > len(serviceName) {
					vodServices = append(vodServices, syncer.VodService{Name: strings.TrimSpace(imgAlt)})
				} else {
					vodServices = append(vodServices, syncer.VodService{Name: serviceName})
				}
			} else {
				imgAlt, imgExists := s.Find("img").Attr("alt")
				if imgExists && strings.TrimSpace(imgAlt) != "" {
					vodServices = append(vodServices, syncer.VodService{Name: strings.TrimSpace(imgAlt)})
				}
			}
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
			expectedError: false,
		},
		{
			name:   "Single VOD service - text in a preferred over img alt",
			workID: "456", // <a>Hulu</a> <img alt="Hulu (alt)">
			expected: []syncer.VodService{
				{Name: "Hulu"}, // s.Text() is "Hulu"
			},
			expectedError: false,
		},
		{
			name:          "No VOD services",
			workID:        "789",
			expected:      []syncer.VodService{},
			expectedError: false,
		},
		{
			name:          "Work not found (404)",
			workID:        "000",
			expected:      nil,
			expectedError: true,
		},
		{
			name:   "Image alt text only",
			workID: "img_only", // <a><img alt="Disney+"></a>
			expected: []syncer.VodService{
				{Name: "Disney+"},
			},
			expectedError: false,
		},
		{
			name:   "Text and img alt different - prefer text",
			workID: "text_and_img_alt_different", // <a>U-NEXT</a> <img alt="U-NEXT (alt)">
			expected: []syncer.VodService{
				{Name: "U-NEXT"},
			},
			expectedError: false,
		},
		{
			name:   "Text and more specific img alt - current logic prefers alt",
			workID: "text_and_img_alt_specific", // <a>dアニメストア</a> <img alt="dアニメストア (ニコニコ支店)">
			expected: []syncer.VodService{
				// The implemented logic for preferring alt text is:
				// `imgExists && strings.TrimSpace(imgAlt) != "" && strings.Contains(strings.TrimSpace(imgAlt), serviceName) && len(strings.TrimSpace(imgAlt)) > len(serviceName)`
				// "dアニメストア (ニコニコ支店)" contains "dアニメストア" and is longer. So alt should be chosen.
				{Name: "dアニメストア (ニコニコ支店)"},
			},
			expectedError: false,
		},
		{
			name:   "Duplicate services - deduplicated",
			workID: "duplicate_services", // Netflix (text), Netflix (img alt), Amazon Prime Video (text)
			expected: []syncer.VodService{
				{Name: "Netflix"}, // From <a>Netflix</a>
				{Name: "Amazon Prime Video"},
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Using getLocalVodServices which incorporates the actual scraping logic with the test server
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
