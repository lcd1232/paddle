package paddle

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

const (
	defaultBaseURL = "https://vendors.paddle.com/api/2.0/"
	userAgent      = "lcd1232-paddle"
)

type Client struct {
	client         *http.Client
	vendorID       string
	vendorAuthCode string
	BaseURL        *url.URL
	UserAgent      string
}

func NewClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Client{
		client:    client,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}
}

func NewClientWithAuthentication(client *http.Client, vendorID string, vendorAuthCode string) *Client {
	c := NewClient(client)
	c.vendorID = vendorID
	c.vendorAuthCode = vendorAuthCode
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.Reader
	form := url.Values{}
	if c.vendorID != "" {
		form.Set("vendor_id", c.vendorID)
	}
	if c.vendorAuthCode != "" {
		form.Set("vendor_auth_code", c.vendorAuthCode)
	}
	if body != nil {
		if err := encoder.Encode(body, form); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	if len(form) > 0 {
		buf = strings.NewReader(form.Encode())
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}
