package paddle

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
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
		wantForm       url.Values
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
			wantForm: map[string][]string{
				"vendor_id":        {"123"},
				"vendor_auth_code": {"12ac"},
				"product_id":       {"5"},
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
				assert.Equal(t, tc.wantForm, r.PostForm)
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
