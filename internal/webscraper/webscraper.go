package webscraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/miscord-dev/annict-epgstation-connector/internal/syncer"
	// "golang.org/x/exp/slog" // slog can be removed if not used elsewhere in this file after this change
)

// GetVodServices scrapes the Annict work page and returns a list of VOD services.
func GetVodServices(workID string) ([]syncer.VodService, error) {
	url := fmt.Sprintf("https://annict.com/works/%s", workID)
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
	// Selector for the links under the VOD list
	selector := "body > div > div.l-default__main.d-flex.flex-column > div.l-default__content > div.c-work-header.pt-3 > div.container > div > div.col.mt-3.mt-sm-0 > ul.list-inline.mt-2 > li > a"
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		serviceName := strings.TrimSpace(s.Text())
		if serviceName != "" {
			vodServices = append(vodServices, syncer.VodService{Name: serviceName})
		}
		// If serviceName is empty, the <a> tag is simply ignored.
	})

	// Deduplication logic remains useful.
	if len(vodServices) > 0 {
		seen := make(map[string]bool)
		uniqueServices := []syncer.VodService{}
		for _, service := range vodServices {
            name := strings.TrimSpace(service.Name)
            if name == "" {
                continue
            }
			if _, ok := seen[name]; !ok {
				seen[name] = true
				uniqueServices = append(uniqueServices, syncer.VodService{Name: name})
			}
		}
		vodServices = uniqueServices
	}
	return vodServices, nil
}
