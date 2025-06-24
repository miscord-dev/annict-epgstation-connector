package vod

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"log/slog"
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

// RetryConfig contains configuration for retry logic
type RetryConfig struct {
	MaxRetries    int
	BaseDelay     time.Duration
	MaxDelay      time.Duration
	BackoffFactor float64
}

// RateLimitConfig contains configuration for rate limiting
type RateLimitConfig struct {
	RequestDelay time.Duration
}

// DefaultRetryConfig returns the default retry configuration
func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxRetries:    3,
		BaseDelay:     1 * time.Second,
		MaxDelay:      30 * time.Second,
		BackoffFactor: 2.0,
	}
}

// DefaultRateLimitConfig returns the default rate limit configuration
func DefaultRateLimitConfig() RateLimitConfig {
	return RateLimitConfig{
		RequestDelay: 1 * time.Second,
	}
}

// Checker handles VOD service availability checking
type Checker struct {
	httpClient      *http.Client
	enableFallback  bool
	retryConfig     RetryConfig
	rateLimitConfig RateLimitConfig
	lastRequestTime time.Time
}

// NewChecker creates a new VOD service checker
func NewChecker() *Checker {
	return &Checker{
		httpClient:      &http.Client{Timeout: 30 * time.Second},
		enableFallback:  false,
		retryConfig:     DefaultRetryConfig(),
		rateLimitConfig: DefaultRateLimitConfig(),
	}
}

// NewCheckerWithFallback creates a new VOD service checker with fallback enabled
func NewCheckerWithFallback(enableFallback bool) *Checker {
	return &Checker{
		httpClient:      &http.Client{Timeout: 30 * time.Second},
		enableFallback:  enableFallback,
		retryConfig:     DefaultRetryConfig(),
		rateLimitConfig: DefaultRateLimitConfig(),
	}
}

// NewCheckerWithConfig creates a new VOD service checker with custom configuration
func NewCheckerWithConfig(enableFallback bool, retryConfig RetryConfig, rateLimitConfig RateLimitConfig) *Checker {
	return &Checker{
		httpClient:      &http.Client{Timeout: 30 * time.Second},
		enableFallback:  enableFallback,
		retryConfig:     retryConfig,
		rateLimitConfig: rateLimitConfig,
	}
}

// CheckVODServices checks which streaming services have the given anime available
func (c *Checker) CheckVODServices(ctx context.Context, annictWorkID int) ([]StreamingService, error) {
	// Apply rate limiting
	if err := c.applyRateLimit(ctx); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://annict.com/works/%d", annictWorkID)

	var lastErr error
	for attempt := 0; attempt <= c.retryConfig.MaxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		// Set user agent to avoid being blocked
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("failed to fetch page: %w", err)
			if attempt < c.retryConfig.MaxRetries {
				delay := c.calculateBackoffDelay(attempt)
				slog.Warn("Request failed, retrying",
					slog.Int("annict_work_id", annictWorkID),
					slog.Int("attempt", attempt+1),
					slog.Int("max_retries", c.retryConfig.MaxRetries),
					slog.Duration("delay", delay),
					slog.String("error", err.Error()))

				select {
				case <-time.After(delay):
				case <-ctx.Done():
					return nil, ctx.Err()
				}
				continue
			}
			return nil, lastErr
		}

		defer func() {
			if err := resp.Body.Close(); err != nil {
				slog.Warn("failed to close response body", slog.String("error", err.Error()))
			}
		}()

		// Handle rate limiting (429) and server errors (5xx)
		if resp.StatusCode == http.StatusTooManyRequests || (resp.StatusCode >= 500 && resp.StatusCode < 600) {
			lastErr = fmt.Errorf("HTTP error: %d", resp.StatusCode)
			if attempt < c.retryConfig.MaxRetries {
				delay := c.calculateBackoffDelay(attempt)

				// Check for Retry-After header
				if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
					if retrySeconds, err := strconv.Atoi(retryAfter); err == nil {
						suggestedDelay := time.Duration(retrySeconds) * time.Second
						if suggestedDelay > delay {
							delay = suggestedDelay
						}
					}
				}

				slog.Warn("Rate limited or server error, retrying",
					slog.Int("annict_work_id", annictWorkID),
					slog.Int("status_code", resp.StatusCode),
					slog.Int("attempt", attempt+1),
					slog.Int("max_retries", c.retryConfig.MaxRetries),
					slog.Duration("delay", delay))

				select {
				case <-time.After(delay):
				case <-ctx.Done():
					return nil, ctx.Err()
				}
				continue
			}
			return nil, lastErr
		}

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

		// Update last request time on successful request
		c.lastRequestTime = time.Now()
		return services, nil
	}

	return nil, lastErr
}

// applyRateLimit applies rate limiting between requests
func (c *Checker) applyRateLimit(ctx context.Context) error {
	if c.rateLimitConfig.RequestDelay <= 0 {
		return nil
	}

	if !c.lastRequestTime.IsZero() {
		elapsed := time.Since(c.lastRequestTime)
		if elapsed < c.rateLimitConfig.RequestDelay {
			waitTime := c.rateLimitConfig.RequestDelay - elapsed
			slog.Debug("Rate limiting: waiting before next request",
				slog.Duration("wait_time", waitTime))

			select {
			case <-time.After(waitTime):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	return nil
}

// calculateBackoffDelay calculates the delay for the next retry attempt using exponential backoff
func (c *Checker) calculateBackoffDelay(attempt int) time.Duration {
	delay := time.Duration(float64(c.retryConfig.BaseDelay) * math.Pow(c.retryConfig.BackoffFactor, float64(attempt)))
	if delay > c.retryConfig.MaxDelay {
		delay = c.retryConfig.MaxDelay
	}
	return delay
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

	found := c.searchServicesWithSelectors(doc, selectors, &services)

	// If no specific VOD section was found and fallback is enabled,
	// fallback to searching all links
	if !found && c.enableFallback {
		c.searchServicesWithSelectors(doc, []string{"a[href]"}, &services)
	}

	// Remove duplicates
	services = removeDuplicateServices(services)

	return services
}

// searchServicesWithSelectors searches for streaming services using the given selectors
func (c *Checker) searchServicesWithSelectors(doc *goquery.Document, selectors []string, services *[]StreamingService) bool {
	found := false
	for _, selector := range selectors {
		doc.Find(selector).Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if !exists {
				return
			}

			if service := c.identifyStreamingService(strings.ToLower(href)); service != "" {
				*services = append(*services, service)
				found = true
			}
		})

		// If we found some streaming services with this selector, stop trying others
		if found && len(selectors) > 1 {
			break
		}
	}
	return found
}

// identifyStreamingService identifies the streaming service from a URL
func (c *Checker) identifyStreamingService(href string) StreamingService {
	if strings.Contains(href, "netflix.com") {
		return Netflix
	}
	if strings.Contains(href, "amazon.co.jp") && (strings.Contains(href, "prime") || strings.Contains(href, "video")) {
		return AmazonPrime
	}
	if strings.Contains(href, "hulu.jp") || strings.Contains(href, "hulu.com") {
		return Hulu
	}
	if strings.Contains(href, "disneyplus.com") || strings.Contains(href, "disney") {
		return Disney
	}
	if strings.Contains(href, "abema.tv") {
		return Abema
	}
	if strings.Contains(href, "crunchyroll.com") {
		return Crunchyroll
	}
	if strings.Contains(href, "funimation.com") {
		return Funimation
	}
	if strings.Contains(href, "dazn.com") {
		return Dazn
	}
	if strings.Contains(href, "b-ch.com") {
		return Bandai
	}
	if strings.Contains(href, "nicovideo.jp") || strings.Contains(href, "ch.nicovideo.jp") {
		return Nico
	}
	if strings.Contains(href, "animestore.docomo.ne.jp") {
		return DAnime
	}
	return ""
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
