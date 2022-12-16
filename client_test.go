package paddle

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient(Settings{
		URL:            SandboxBaseURL,
		CheckoutURL:    SandboxCheckoutBaseURL,
		VendorID:       "123",
		VendorAuthCode: "123abc",
	})
	require.NoError(t, err)
	require.NotNil(t, c)
	assert.NotNil(t, c.client)
	assert.NotNil(t, c.BaseURL)
	assert.NotNil(t, c.CheckoutBaseURL)
	assert.Equal(t, userAgent, c.UserAgent)
	assert.Equal(t, "123", c.vendorID)
	assert.Equal(t, "123abc", c.vendorAuthCode)
}

func TestNewClientEmptySettings(t *testing.T) {
	c, err := NewClient(Settings{})
	require.NoError(t, err)
	require.NotNil(t, c)
	assert.NotNil(t, c.client)
	assert.NotNil(t, c.BaseURL)
	assert.NotNil(t, c.CheckoutBaseURL)
	assert.Equal(t, userAgent, c.UserAgent)
}

func TestNewClientInvalidURL(t *testing.T) {
	_, err := NewClient(Settings{
		URL: "http\\:a",
	})
	require.Error(t, err)
}

func newSandboxClient(t *testing.T) *Client {
	t.Helper()
	vendorID := os.Getenv("TEST_PADDLE_VENDOR_ID")
	authCode := os.Getenv("TEST_PADDLE_AUTH_CODE")
	if vendorID == "" || authCode == "" {
		t.Skip("vendor_id or auth_code not set")
	}
	c, err := NewClient(Settings{
		URL:            SandboxBaseURL,
		CheckoutURL:    SandboxCheckoutBaseURL,
		VendorID:       vendorID,
		VendorAuthCode: authCode,
	})
	require.NoError(t, err)
	return c
}

func getEnv(t *testing.T, name string) string {
	t.Helper()
	v := os.Getenv(name)
	if v == "" {
		t.Skipf("%s not set", name)
	}
	return v
}
