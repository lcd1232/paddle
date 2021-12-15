package paddle

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAudience(t *testing.T) {
	for _, tc := range []struct {
		name             string
		vendorID         string
		email            string
		marketingConsent bool
		responseCode     int
		responseBody     []byte
		wantErr          bool
		wantQuery        url.Values
		wantUserID       int
	}{
		{
			name:             "subscribe",
			vendorID:         "123",
			email:            "user@example.com",
			marketingConsent: true,
			responseCode:     http.StatusOK,
			responseBody:     []byte(`{"user_id": 10}`),
			wantQuery: map[string][]string{
				"email":             {"user@example.com"},
				"marketing_consent": {"1"},
			},
			wantUserID: 10,
		},
		{
			name:             "unsubscribe",
			vendorID:         "123",
			email:            "user@example.com",
			marketingConsent: false,
			responseCode:     http.StatusOK,
			responseBody:     []byte(`{"user_id": 10}`),
			wantQuery: map[string][]string{
				"email":             {"user@example.com"},
				"marketing_consent": {"1"},
			},
			wantUserID: 10,
		},
		{
			name:             "error",
			vendorID:         "123",
			email:            "user@example.com",
			marketingConsent: true,
			responseCode:     http.StatusOK,
			responseBody: []byte(`{
    "error": {
        "code": 102,
        "message": "The selected vendor id is invalid."
    },
    "success": false
}`),
			wantErr: true,
		},
		{
			name:         "invalid json",
			vendorID:     "123",
			responseCode: http.StatusOK,
			responseBody: []byte(`}`),
			wantErr:      true,
		},
		{
			name:         "502 status",
			vendorID:     "123",
			responseCode: http.StatusBadGateway,
			responseBody: []byte(`{}`),
			wantErr:      true,
		},
		{
			name:    "empty vendor id",
			wantErr: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			WithTestServer(t, tc.responseCode, tc.responseBody, func(url string, rCh <-chan *http.Request) {
				c, err := NewClient(Settings{
					CheckoutURL: url,
					VendorID:    tc.vendorID,
				})
				require.NoError(t, err)
				userID, err := c.Audience(context.Background(), tc.email, tc.marketingConsent)
				if tc.wantErr {
					require.Error(t, err)
					return
				}
				require.NoError(t, err)
				assert.Equal(t, tc.wantUserID, userID)
				r := <-rCh
				assert.Equal(t, fmt.Sprintf("api/1.0/audience/%s/add", tc.vendorID), r.URL.Path)
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, tc.wantQuery, r.URL.Query())
			})
		})
	}
}
