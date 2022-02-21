package paddle

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func WithTestServer(t *testing.T, responseCode int, responseBody []byte, f func(url string, rCh <-chan *http.Request)) {
	ch := make(chan *http.Request, 1)
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err == nil {
			r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		}
		w.WriteHeader(responseCode)
		if len(responseBody) > 0 {
			_, err := w.Write(responseBody)
			assert.NoError(t, err)
		}
		ch <- r
	}))
	defer s.Close()

	f(s.URL+"/", ch)
}

func TestGeneratePayLink(t *testing.T) {
	for _, tc := range []struct {
		name           string
		vendorID       string
		vendorAuthCode string
		request        GeneratePayLinkRequest
		responseCode   int
		responseBody   []byte
		wantErr        bool
		wantForm       func(t *testing.T, values url.Values)
		wantURL        string
	}{
		{
			name:           "one parameter only",
			vendorID:       "123",
			vendorAuthCode: "12ac",
			request: GeneratePayLinkRequest{
				ProductID: 5,
			},
			responseCode: http.StatusOK,
			responseBody: []byte(`{
  "success": true,
  "response": {
    "url": "https://checkout.paddle.com/checkout/custom/eyJ0IjoiUHJvZ"
  }
}`),
			wantForm: func(t *testing.T, values url.Values) {
				assert.Equal(t, map[string][]string{
					"vendor_id":        {"123"},
					"vendor_auth_code": {"12ac"},
					"product_id":       {"5"},
				}, values)
			},
			wantURL: "https://checkout.paddle.com/checkout/custom/eyJ0IjoiUHJvZ",
		},
		{
			name:           "prices",
			vendorID:       "123",
			vendorAuthCode: "12ac",
			request: GeneratePayLinkRequest{
				ProductID: 5,
				Prices: map[string]string{
					"USD": "4.99",
					"RUB": "199.99",
				},
			},
			responseCode: http.StatusOK,
			responseBody: []byte(`{
  "success": true,
  "response": {
    "url": "https://checkout.paddle.com/checkout/custom/eyJ0IjoiUHJvZ"
  }
}`),
			wantForm: func(t *testing.T, values url.Values) {
				assert.Equal(t, "123", values.Get("vendor_id"))
				assert.Equal(t, "12ac", values.Get("vendor_auth_code"))
				assert.Equal(t, "5", values.Get("product_id"))
				if strings.Contains(values.Get("prices[0]"), "USD") {
					assert.Equal(t, "4.99", values.Get("prices[0]"))
					assert.Equal(t, "199.99", values.Get("prices[1]"))
				} else {
					assert.Equal(t, "199.99", values.Get("prices[0]"))
					assert.Equal(t, "4.99", values.Get("prices[1]"))
				}
			},
			wantURL: "https://checkout.paddle.com/checkout/custom/eyJ0IjoiUHJvZ",
		},
		{
			name:           "prices + recurring prices",
			vendorID:       "123",
			vendorAuthCode: "12ac",
			request: GeneratePayLinkRequest{
				ProductID: 5,
				Prices: map[string]string{
					"USD": "4.99",
					"RUB": "199.99",
				},
				RecurringPrices: map[string]string{
					"USD": "9.99",
					"RUB": "399.99",
				},
			},
			responseCode: http.StatusOK,
			responseBody: []byte(`{
  "success": true,
  "response": {
    "url": "https://checkout.paddle.com/checkout/custom/eyJ0IjoiUHJvZ"
  }
}`),
			wantForm: func(t *testing.T, values url.Values) {
				assert.Equal(t, "123", values.Get("vendor_id"))
				assert.Equal(t, "12ac", values.Get("vendor_auth_code"))
				assert.Equal(t, "5", values.Get("product_id"))
				if strings.Contains(values.Get("prices[0]"), "USD") {
					assert.Equal(t, "4.99", values.Get("prices[0]"))
					assert.Equal(t, "199.99", values.Get("prices[1]"))
				} else {
					assert.Equal(t, "199.99", values.Get("prices[0]"))
					assert.Equal(t, "4.99", values.Get("prices[1]"))
				}
				if strings.Contains(values.Get("recurring_prices[0]"), "USD") {
					assert.Equal(t, "9.99", values.Get("recurring_prices[0]"))
					assert.Equal(t, "399.99", values.Get("recurring_prices[1]"))
				} else {
					assert.Equal(t, "399.99", values.Get("recurring_prices[0]"))
					assert.Equal(t, "9.99", values.Get("recurring_prices[1]"))
				}
			},
			wantURL: "https://checkout.paddle.com/checkout/custom/eyJ0IjoiUHJvZ",
		},
		{
			name:           "error",
			vendorID:       "123",
			vendorAuthCode: "12ac",
			request: GeneratePayLinkRequest{
				ProductID: 5,
			},
			responseCode: http.StatusOK,
			responseBody: []byte(`{
  "success": false,
  "error": {
    "code": 130,
	"message": "The allowed uses must be a number."
  }
}`),
			wantErr: true,
		},
		{
			name:         "invalid json",
			responseCode: http.StatusOK,
			responseBody: []byte(`}`),
			wantErr:      true,
		},
		{
			name:         "502 status",
			responseCode: http.StatusBadGateway,
			responseBody: []byte(`{"success": true,
  "response": {
    "url": "https://checkout.paddle.com/checkout/custom/eyJ0IjoiUHJvZ"
  }
}`),
			wantErr: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			WithTestServer(t, tc.responseCode, tc.responseBody, func(url string, rCh <-chan *http.Request) {
				c, err := NewClient(Settings{
					URL:            url,
					VendorID:       tc.vendorID,
					VendorAuthCode: tc.vendorAuthCode,
				})
				require.NoError(t, err)
				urlStr, err := c.GeneratePayLink(context.Background(), tc.request)
				if tc.wantErr {
					require.Error(t, err)
					return
				}
				require.NoError(t, err)
				assert.Equal(t, tc.wantURL, urlStr)
				r := <-rCh
				require.NoError(t, r.ParseForm())
				tc.wantForm(t, r.PostForm)
			})
		})
	}
}

func TestGeneratePayLinkSandbox(t *testing.T) {
	client := newSandboxClient(t)
	productIDStr := getEnv(t, "TEST_PRODUCT_ID")
	productID, err := strconv.Atoi(productIDStr)
	require.NoError(t, err)
	urlStr, err := client.GeneratePayLink(context.Background(), GeneratePayLinkRequest{
		ProductID: productID,
		Prices: map[string]string{
			"USD": "9.99",
			"EUR": "8.99",
		},
	})
	require.NoError(t, err)
	assert.NotEmpty(t, urlStr)
}
