package paddle

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
	responseBody := `{
  "success": true,
  "response": {
    "url": "https://checkout.paddle.com/checkout/custom/eyJ0IjoiUHJvZ"
  }
}`
	WithTestServer(t, http.StatusOK, []byte(responseBody), func(url string, rCh <-chan *http.Request) {
		c := NewTestClient(t, url, nil)
		url, err := c.GeneratePayLink(context.Background())
		require.NoError(t, err)
		assert.Equal(t, "https://checkout.paddle.com/checkout/custom/eyJ0IjoiUHJvZ", url)
		r := <-rCh
		require.NoError(t, r.ParseForm())
		assert.Equal(t, map[string][]string{
			"vendor_id":        {"123"},
			"vendor_auth_code": {"12ac"},
			"product_id":       {"5"},
		}, r.PostForm)
	})
}
