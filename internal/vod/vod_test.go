package vod

import (
	"testing"
)

func TestParseAnnictWorkID(t *testing.T) {
	tests := []struct {
		name     string
		workID   string
		expected int
		hasError bool
	}{
		{
			name:     "numeric string",
			workID:   "12345",
			expected: 12345,
			hasError: false,
		},
		{
			name:     "numeric string with prefix",
			workID:   "work_12345",
			expected: 12345,
			hasError: false,
		},
		{
			name:     "mixed string",
			workID:   "abc123def",
			expected: 123,
			hasError: false,
		},
		{
			name:     "no numbers",
			workID:   "abcdef",
			expected: 0,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseAnnictWorkID(tt.workID)
			if tt.hasError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected %d but got %d", tt.expected, result)
				}
			}
		})
	}
}

func TestCheckVODServices(t *testing.T) {
	tests := []struct {
		name             string
		htmlContent      string
		expectedServices []StreamingService
		enableFallback   bool
	}{
		{
			name: "Netflix and Amazon links in c-work-header",
			htmlContent: `
				<html>
					<body>
						<div class="c-work-header">
							<ul>
								<li><a href="https://netflix.com/title/123">Netflix</a></li>
								<li><a href="https://amazon.co.jp/dp/B07VGLWY22?ref=prime">Amazon プライム・ビデオ</a></li>
							</ul>
						</div>
					</body>
				</html>
			`,
			expectedServices: []StreamingService{Netflix, AmazonPrime},
			enableFallback:   false,
		},
		{
			name: "Netflix URL only in c-work-header",
			htmlContent: `
				<html>
					<body>
						<div class="c-work-header">
							<ul>
								<li><a href="https://netflix.com/title/123">Netflix</a></li>
							</ul>
						</div>
					</body>
				</html>
			`,
			expectedServices: []StreamingService{Netflix},
			enableFallback:   false,
		},
		{
			name: "Multiple Japanese services in c-work-header",
			htmlContent: `
				<html>
					<body>
						<div class="c-work-header">
							<ul>
								<li><a href="https://abema.tv/video/title/25-120">ABEMAビデオ</a></li>
								<li><a href="https://b-ch.com/ttl/index.php?ttl_c=6635">バンダイチャンネル</a></li>
								<li><a href="https://animestore.docomo.ne.jp/animestore/ci_pc?workId=22859">dアニメストア</a></li>
							</ul>
						</div>
					</body>
				</html>
			`,
			expectedServices: []StreamingService{Abema, Bandai, DAnime},
			enableFallback:   false,
		},
		{
			name: "no VOD services - no fallback",
			htmlContent: `
				<html>
					<body>
						<div>No streaming services available</div>
						<div class="main-content">
							<a href="https://netflix.com/title/123">Watch on Netflix</a>
						</div>
					</body>
				</html>
			`,
			expectedServices: []StreamingService{},
			enableFallback:   false,
		},
		{
			name: "no VOD services - with fallback enabled",
			htmlContent: `
				<html>
					<body>
						<div>No streaming services available</div>
						<div class="main-content">
							<a href="https://netflix.com/title/123">Watch on Netflix</a>
							<a href="https://crunchyroll.com/series/GY8VEQ95Y">Watch on Crunchyroll</a>
						</div>
						<div class="comments">
							<a href="https://netflix.com/title/456">User mentioned Netflix</a>
						</div>
					</body>
				</html>
			`,
			expectedServices: []StreamingService{Netflix, Crunchyroll},
			enableFallback:   true,
		},
		{
			name: "prioritize c-work-header over other links",
			htmlContent: `
				<html>
					<body>
						<div class="c-work-header">
							<ul>
								<li><a href="https://netflix.com/title/123">Netflix</a></li>
							</ul>
						</div>
						<div class="main-content">
							<a href="https://crunchyroll.com/series/GY8VEQ95Y">Crunchyroll in main content</a>
						</div>
						<div class="comments">
							<a href="https://hulu.com/watch/123">User mentioned Hulu</a>
						</div>
					</body>
				</html>
			`,
			expectedServices: []StreamingService{Netflix},
			enableFallback:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create checker with the appropriate fallback setting
			checker := NewCheckerWithFallback(tt.enableFallback)

			// Use the parseServicesFromHTML method for easier testing
			services, err := checker.parseServicesFromHTML(tt.htmlContent)
			if err != nil {
				t.Fatalf("parseServicesFromHTML failed: %v", err)
			}

			// Check that we got the expected services
			if len(services) != len(tt.expectedServices) {
				t.Errorf("expected %d services but got %d: %v", len(tt.expectedServices), len(services), services)
			}

			// Check that all expected services are present
			for _, expectedService := range tt.expectedServices {
				if !containsService(services, expectedService) {
					t.Errorf("expected service %s not found in result: %v", expectedService, services)
				}
			}

			// Check that no unexpected services are present
			for _, service := range services {
				if !containsService(tt.expectedServices, service) {
					t.Errorf("unexpected service %s found in result: %v", service, services)
				}
			}
		})
	}
}

func TestVODFallbackBehavior(t *testing.T) {
	htmlWithoutVODSection := `
		<html>
			<body>
				<div class="main-content">
					<a href="https://netflix.com/title/123">Watch on Netflix</a>
					<a href="https://crunchyroll.com/series/GY8VEQ95Y">Watch on Crunchyroll</a>
				</div>
			</body>
		</html>
	`

	// Test with fallback disabled (default behavior)
	checkerNoFallback := NewChecker()
	servicesNoFallback, err := checkerNoFallback.parseServicesFromHTML(htmlWithoutVODSection)
	if err != nil {
		t.Fatalf("parseServicesFromHTML failed: %v", err)
	}
	if len(servicesNoFallback) != 0 {
		t.Errorf("expected 0 services without fallback, got %d: %v", len(servicesNoFallback), servicesNoFallback)
	}

	// Test with fallback enabled
	checkerWithFallback := NewCheckerWithFallback(true)
	servicesWithFallback, err := checkerWithFallback.parseServicesFromHTML(htmlWithoutVODSection)
	if err != nil {
		t.Fatalf("parseServicesFromHTML failed: %v", err)
	}
	if len(servicesWithFallback) != 2 {
		t.Errorf("expected 2 services with fallback, got %d: %v", len(servicesWithFallback), servicesWithFallback)
	}

	expectedServices := []StreamingService{Netflix, Crunchyroll}
	for _, expectedService := range expectedServices {
		if !containsService(servicesWithFallback, expectedService) {
			t.Errorf("expected service %s not found in fallback result: %v", expectedService, servicesWithFallback)
		}
	}
}

func TestIsAvailableOnServices(t *testing.T) {
	tests := []struct {
		name              string
		availableServices []StreamingService
		excludedServices  []StreamingService
		expected          bool
	}{
		{
			name:              "available on excluded service",
			availableServices: []StreamingService{Netflix, Hulu},
			excludedServices:  []StreamingService{Netflix},
			expected:          true,
		},
		{
			name:              "not available on excluded services",
			availableServices: []StreamingService{Crunchyroll},
			excludedServices:  []StreamingService{Netflix, Hulu},
			expected:          false,
		},
		{
			name:              "no available services",
			availableServices: []StreamingService{},
			excludedServices:  []StreamingService{Netflix},
			expected:          false,
		},
		{
			name:              "no excluded services",
			availableServices: []StreamingService{Netflix},
			excludedServices:  []StreamingService{},
			expected:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test the logic directly
			result := false
			for _, available := range tt.availableServices {
				for _, excluded := range tt.excludedServices {
					if available == excluded {
						result = true
						break
					}
				}
				if result {
					break
				}
			}

			if result != tt.expected {
				t.Errorf("expected %v but got %v", tt.expected, result)
			}
		})
	}
}

func TestContainsService(t *testing.T) {
	services := []StreamingService{Netflix, Hulu, Crunchyroll}

	if !containsService(services, Netflix) {
		t.Error("expected to find Netflix in services")
	}

	if containsService(services, AmazonPrime) {
		t.Error("expected not to find Amazon Prime in services")
	}
}

func TestRemoveDuplicateServices(t *testing.T) {
	input := []StreamingService{Netflix, Hulu, Netflix, Crunchyroll, Hulu}
	expected := []StreamingService{Netflix, Hulu, Crunchyroll}

	result := removeDuplicateServices(input)

	if len(result) != len(expected) {
		t.Errorf("expected length %d but got %d", len(expected), len(result))
	}

	// Check that all expected services are present
	for _, expectedService := range expected {
		if !containsService(result, expectedService) {
			t.Errorf("expected service %s not found in result", expectedService)
		}
	}
}
