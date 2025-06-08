package syncer

import (
	"context"
	"errors"
	"io"
	"context"
	"errors"
	"io"
	// "net/http" // No longer needed for this specific mock
	// "net/http/httptest" // No longer needed for this specific mock
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/cockroachdb/pebble"
	"github.com/miscord-dev/annict-epgstation-connector/epgstation"
	"github.com/miscord-dev/annict-epgstation-connector/internal/webscraper"
	// "github.com/PuerkitoBio/goquery" // No longer needed for this specific mock
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slog"
)

// Mock Annict and EPGStation clients for testing
type mockAnnictClient struct {
	graphql.Client
}

type mockEPGStationClient struct {
	*epgstation.Client
	postRulesFunc func(ctx context.Context, body epgstation.PostRulesJSONRequestBody) (*http.Response, error)
	getRulesKeywordFunc func(ctx context.Context, params *epgstation.GetRulesKeywordParams) (*http.Response, error)
}

func (m *mockEPGStationClient) PostRules(ctx context.Context, body epgstation.PostRulesJSONRequestBody) (*http.Response, error) {
	if m.postRulesFunc != nil {
		return m.postRulesFunc(ctx, body)
	}
	// Default mock response
	rec := httptest.NewRecorder() // Use a new recorder for default
	rec.WriteHeader(http.StatusCreated)
	_, _ = rec.WriteString(`{"ruleId": 0}`) // Default ruleId
	return rec.Result(), nil
}

func (m *mockEPGStationClient) GetRulesKeyword(ctx context.Context, params *epgstation.GetRulesKeywordParams) (*http.Response, error) {
	if m.getRulesKeywordFunc != nil {
		return m.getRulesKeywordFunc(ctx, params)
	}
	// Default mock response (no rules found)
	rec := httptest.NewRecorder()
	rec.WriteHeader(http.StatusOK)
	_, _ = rec.WriteString(`{"items":[]}`)
	return rec.Result(), nil
}


func TestRegisterRuleToEPGStation_VODFilter(t *testing.T) {
	originalGetVodServices := webscraper.GetVodServices
	webscraper.GetVodServices = func(workID string) ([]VodService, error) { // This VodService is syncer.VodService
		switch workID {
		case "vod_available_subscribed_netflix":
			return []VodService{{Name: "Netflix"}}, nil
		case "vod_available_subscribed_amazon":
			return []VodService{{Name: "Amazon Prime Video"}}, nil
		case "vod_available_not_subscribed_hulu":
			return []VodService{{Name: "Hulu"}}, nil
		case "vod_not_available":
			return []VodService{}, nil
		case "vod_scraping_error":
			return nil, errors.New("simulated scraping error")
		case "existing_annict_id_no_vod_conflict":
		    return []VodService{}, nil
		case "new_annict_id_existing_keyword_no_vod_conflict":
		    return []VodService{}, nil
		default:
			return []VodService{}, nil
		}
	}
	defer func() { webscraper.GetVodServices = originalGetVodServices }()

	tempDir := t.TempDir()
	db, err := pebble.Open(tempDir, &pebble.Options{})
	require.NoError(t, err)
	defer db.Close()

	tests := []struct {
		name                   string
		work                   annictWork
		// subscribedServices is hardcoded in syncer.go as []string{"Netflix", "Amazon Prime Video"}
		mockPostRulesFunc      func(ctx context.Context, body epgstation.PostRulesJSONRequestBody) (*http.Response, error)
		mockGetRulesKeywordFunc func(ctx context.Context, params *epgstation.GetRulesKeywordParams) (*http.Response, error)
		expectRuleRegistration bool
		expectedError          bool
	}{
		{
			name: "Work available on subscribed VOD (Netflix) - Skip registration",
			work: annictWork{ID: "vod_available_subscribed_netflix", Title: "Test Anime Netflix"},
			expectRuleRegistration: false,
			expectedError:          false,
		},
		{
			name: "Work available on subscribed VOD (Amazon) - Skip registration",
			work: annictWork{ID: "vod_available_subscribed_amazon", Title: "Test Anime Amazon"},
			expectRuleRegistration: false,
			expectedError:          false,
		},
		{
			name: "Work available on non-subscribed VOD (Hulu) - Register rule",
			work: annictWork{ID: "vod_available_not_subscribed_hulu", Title: "Test Anime Hulu"},
			mockPostRulesFunc: func(ctx context.Context, body epgstation.PostRulesJSONRequestBody) (*http.Response, error) {
				rec := httptest.NewRecorder()
				rec.WriteHeader(http.StatusCreated)
				_, _ = rec.WriteString(`{"ruleId": 101}`)
				return rec.Result(), nil
			},
			expectRuleRegistration: true,
			expectedError:          false,
		},
		{
			name: "Work not available on VOD - Register rule",
			work: annictWork{ID: "vod_not_available", Title: "Test Anime No VOD"},
			mockPostRulesFunc: func(ctx context.Context, body epgstation.PostRulesJSONRequestBody) (*http.Response, error) {
				rec := httptest.NewRecorder()
				rec.WriteHeader(http.StatusCreated)
				_, _ = rec.WriteString(`{"ruleId": 102}`)
				return rec.Result(), nil
			},
			expectRuleRegistration: true,
			expectedError:          false,
		},
		{
			name: "VOD scraping error - Register rule (graceful fallback)",
			work: annictWork{ID: "vod_scraping_error", Title: "Test Anime Scraping Error"},
			mockPostRulesFunc: func(ctx context.Context, body epgstation.PostRulesJSONRequestBody) (*http.Response, error) {
				rec := httptest.NewRecorder()
				rec.WriteHeader(http.StatusCreated)
				_, _ = rec.WriteString(`{"ruleId": 103}`)
				return rec.Result(), nil
			},
			expectRuleRegistration: true,
			expectedError:          false,
		},
		{
			name: "Rule already exists by Annict ID (no VOD conflict) - Skip registration",
			work: annictWork{ID: "existing_annict_id_no_vod_conflict", Title: "Test Anime Existing Annict"},
			mockGetRulesKeywordFunc: func(ctx context.Context, params *epgstation.GetRulesKeywordParams) (*http.Response, error) {
				rec := httptest.NewRecorder()
				rec.WriteHeader(http.StatusOK)
				_, _ = rec.WriteString(`{"items":[]}`)
				return rec.Result(), nil
			},
			expectRuleRegistration: false,
			expectedError:          false,
		},
		{
			name: "Rule already exists by Keyword (no VOD conflict) - Skip registration",
			work: annictWork{ID: "new_annict_id_existing_keyword_no_vod_conflict", Title: "Test Anime Existing Keyword"},
			mockGetRulesKeywordFunc: func(ctx context.Context, params *epgstation.GetRulesKeywordParams) (*http.Response, error) {
				rec := httptest.NewRecorder()
				rec.WriteHeader(http.StatusOK)
				_, _ = rec.WriteString(`{"items":[{"id": 201, "keyword": "Test Anime Existing Keyword"}]}`)
				return rec.Result(), nil
			},
			expectRuleRegistration: false,
			expectedError:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset DB for Annict ID check
			if tt.work.ID == "existing_annict_id_no_vod_conflict" {
				_ = db.Delete([]byte(tt.work.ID), pebble.Sync)
				err := db.Set([]byte(tt.work.ID), []byte(`[12345]`), pebble.Sync)
				require.NoError(t, err)
			} else {
				// Ensure Annict ID is not in DB for other tests unless specifically testing that
				err := db.Delete([]byte(tt.work.ID), pebble.Sync)
				if err != nil && !errors.Is(err, pebble.ErrNotFound) {
					require.NoError(t, err) // Fail if it's any error other than not found
				}
			}


			mockESClient := &mockEPGStationClient{
				postRulesFunc:      tt.mockPostRulesFunc,
				getRulesKeywordFunc: tt.mockGetRulesKeywordFunc,
			}

			esClientWithResponses, err := epgstation.NewClientWithResponses("http://localhost", epgstation.WithHTTPClient(mockESClient))
			require.NoError(t, err)

			s := &syncer{
				annictClient: &mockAnnictClient{},
				esClient:     esClientWithResponses,
				db:           db,
			}
            // Suppress slog output during tests
			originalHandler := slog.Default().Handler()
			slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))
			defer slog.SetDefault(slog.New(originalHandler))


			postRulesCalled := false
			if tt.expectRuleRegistration {
				// Override postRulesFunc to track calls for positive cases
				originalPostFunc := mockESClient.postRulesFunc
				mockESClient.postRulesFunc = func(ctx context.Context, body epgstation.PostRulesJSONRequestBody) (*http.Response, error) {
					postRulesCalled = true
					if originalPostFunc != nil {
						return originalPostFunc(ctx, body)
					}
					// Default successful registration if no specific mock provided but registration expected
					rec := httptest.NewRecorder()
					rec.WriteHeader(http.StatusCreated)
					_, _ = rec.WriteString(`{"ruleId": ` + strconv.Itoa(int(time.Now().UnixNano())) + `}`)
					return rec.Result(), nil
				}
			}


			err = s.registerRuleToEPGStation(context.Background(), tt.work)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectRuleRegistration, postRulesCalled, "PostRules call expectation mismatch")
			}
		})
	}
}

// http.Response and httptest are still needed for mockEPGStationClient's method signatures and default responses.
