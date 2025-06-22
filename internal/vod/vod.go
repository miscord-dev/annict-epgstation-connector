package vod

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/exp/slog"
)

// StreamingService represents a VOD streaming service
type StreamingService string

const (
	Netflix     StreamingService = "netflix"
	AmazonPrime StreamingService = "amazon-prime"
	Hulu        StreamingService = "hulu"
	Disney      StreamingService = "disney"
	Abema       StreamingService = "abema"
	Crunchyroll StreamingService = "crunchyroll"
	Funimation  StreamingService = "funimation"
	Dazn        StreamingService = "dazn"
	Bandai      StreamingService = "bandai"
	Nico        StreamingService = "nico"
	DAnime      StreamingService = "danime"
)

// ServiceInfo contains information about where an anime is available
type ServiceInfo struct {
	AnnictWorkID int
	Services     []StreamingService
}

// Checker handles VOD service availability checking
type Checker struct {
	httpClient     *http.Client
	enableFallback bool
}

// NewChecker creates a new VOD service checker
func NewChecker() *Checker {
	return &Checker{
		httpClient:     &http.Client{},
		enableFallback: false,
	}
}

// NewCheckerWithFallback creates a new VOD service checker with fallback enabled
func NewCheckerWithFallback(enableFallback bool) *Checker {
	return &Checker{
		httpClient:     &http.Client{},
		enableFallback: enableFallback,
	}
}

// CheckVODServices checks which streaming services have the given anime available
func (c *Checker) CheckVODServices(ctx context.Context, annictWorkID int) ([]StreamingService, error) {
	url := fmt.Sprintf("https://annict.com/works/%d", annictWorkID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set user agent to avoid being blocked
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	services, err := c.parseServicesFromResponse(resp)
	if err != nil {
		return nil, err
	}

	slog.Debug("VOD services found",
		slog.Int("annict_work_id", annictWorkID),
		slog.Int("service_count", len(services)),
		slog.Any("services", services))

	return services, nil
}

// parseServicesFromResponse extracts streaming services from an HTTP response
func (c *Checker) parseServicesFromResponse(resp *http.Response) ([]StreamingService, error) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	return c.parseServicesFromDocument(doc), nil
}

// parseServicesFromHTML extracts streaming services from HTML content (for testing)
func (c *Checker) parseServicesFromHTML(htmlContent string) ([]StreamingService, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	return c.parseServicesFromDocument(doc), nil
}

// parseServicesFromDocument extracts streaming services from a goquery document
func (c *Checker) parseServicesFromDocument(doc *goquery.Document) []StreamingService {
	var services []StreamingService

	// Look for streaming service links in the specific VOD section
	// Try multiple possible selectors for the streaming services section
	selectors := []string{
		".c-work-header ul a[href]",
		"div.c-work-header ul a[href]",
		".c-work-header a[href]",
		".streaming-services a[href]",
		"div.streaming-services a[href]",
		".vod-services a[href]",
		".watch-links a[href]",
		".streaming a[href]",
	}

	found := false
	for _, selector := range selectors {
		doc.Find(selector).Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if !exists {
				return
			}

			href = strings.ToLower(href)

			// Check for specific streaming service URLs
			if strings.Contains(href, "netflix.com") {
				services = append(services, Netflix)
				found = true
			} else if strings.Contains(href, "amazon.co.jp") && (strings.Contains(href, "prime") || strings.Contains(href, "video")) {
				services = append(services, AmazonPrime)
				found = true
			} else if strings.Contains(href, "hulu.jp") || strings.Contains(href, "hulu.com") {
				services = append(services, Hulu)
				found = true
			} else if strings.Contains(href, "disneyplus.com") || strings.Contains(href, "disney") {
				services = append(services, Disney)
				found = true
			} else if strings.Contains(href, "abema.tv") {
				services = append(services, Abema)
				found = true
			} else if strings.Contains(href, "crunchyroll.com") {
				services = append(services, Crunchyroll)
				found = true
			} else if strings.Contains(href, "funimation.com") {
				services = append(services, Funimation)
				found = true
			} else if strings.Contains(href, "dazn.com") {
				services = append(services, Dazn)
				found = true
			} else if strings.Contains(href, "b-ch.com") {
				services = append(services, Bandai)
				found = true
			} else if strings.Contains(href, "nicovideo.jp") || strings.Contains(href, "ch.nicovideo.jp") {
				services = append(services, Nico)
				found = true
			} else if strings.Contains(href, "animestore.docomo.ne.jp") {
				services = append(services, DAnime)
				found = true
			}
		})

		// If we found some streaming services with this selector, stop trying others
		if found {
			break
		}
	}

	// If no specific VOD section was found and fallback is enabled,
	// fallback to searching all links but filter out obvious non-VOD sections
	if !found && c.enableFallback {
		doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if !exists {
				return
			}

			href = strings.ToLower(href)

			// Check for specific streaming service URLs
			if strings.Contains(href, "netflix.com") {
				services = append(services, Netflix)
			} else if strings.Contains(href, "amazon.co.jp") && (strings.Contains(href, "prime") || strings.Contains(href, "video")) {
				services = append(services, AmazonPrime)
			} else if strings.Contains(href, "hulu.jp") || strings.Contains(href, "hulu.com") {
				services = append(services, Hulu)
			} else if strings.Contains(href, "disneyplus.com") || strings.Contains(href, "disney") {
				services = append(services, Disney)
			} else if strings.Contains(href, "abema.tv") {
				services = append(services, Abema)
			} else if strings.Contains(href, "crunchyroll.com") {
				services = append(services, Crunchyroll)
			} else if strings.Contains(href, "funimation.com") {
				services = append(services, Funimation)
			} else if strings.Contains(href, "dazn.com") {
				services = append(services, Dazn)
			} else if strings.Contains(href, "b-ch.com") {
				services = append(services, Bandai)
			} else if strings.Contains(href, "nicovideo.jp") || strings.Contains(href, "ch.nicovideo.jp") {
				services = append(services, Nico)
			} else if strings.Contains(href, "animestore.docomo.ne.jp") {
				services = append(services, DAnime)
			}
		})
	}

	// Remove duplicates
	services = removeDuplicateServices(services)

	return services
}

// IsAvailableOnServices checks if the anime is available on any of the given services
func (c *Checker) IsAvailableOnServices(ctx context.Context, annictWorkID int, excludedServices []StreamingService) (bool, error) {
	availableServices, err := c.CheckVODServices(ctx, annictWorkID)
	if err != nil {
		return false, err
	}

	for _, available := range availableServices {
		for _, excluded := range excludedServices {
			if available == excluded {
				return true, nil
			}
		}
	}

	return false, nil
}

func containsService(services []StreamingService, service StreamingService) bool {
	for _, s := range services {
		if s == service {
			return true
		}
	}
	return false
}

func removeDuplicateServices(services []StreamingService) []StreamingService {
	keys := make(map[StreamingService]bool)
	var result []StreamingService

	for _, service := range services {
		if !keys[service] {
			keys[service] = true
			result = append(result, service)
		}
	}

	return result
}

// ParseAnnictWorkID extracts the Annict work ID from a work ID string
func ParseAnnictWorkID(workID string) (int, error) {
	// Handle both string and numeric work IDs
	if id, err := strconv.Atoi(workID); err == nil {
		return id, nil
	}

	// Try to extract numeric ID from string using regex
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(workID)
	if match == "" {
		return 0, fmt.Errorf("no numeric ID found in work ID: %s", workID)
	}

	return strconv.Atoi(match)
}
