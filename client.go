package paddle

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

const (
	DefaultBaseURL = "https://vendors.paddle.com/api/"
	SandboxBaseURL = "https://sandbox-vendors.paddle.com/api/"
	userAgent      = "lcd1232-paddle"
)

type Settings struct {
	URL            string
	Client         *http.Client
	VendorID       string
	VendorAuthCode string
}

type Client struct {
	client         *http.Client
	vendorID       string
	vendorAuthCode string
	BaseURL        *url.URL
	UserAgent      string
}

func NewClient(settings Settings) (*Client, error) {
	if settings.URL == "" {
		settings.URL = DefaultBaseURL
	}
	if settings.Client == nil {
		settings.Client = http.DefaultClient
	}
	baseURL, err := url.Parse(settings.URL)
	if err != nil {
		return nil, err
	}

	return &Client{
		client:         settings.Client,
		BaseURL:        baseURL,
		UserAgent:      userAgent,
		vendorID:       settings.VendorID,
		vendorAuthCode: settings.VendorAuthCode,
	}, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
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

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
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
