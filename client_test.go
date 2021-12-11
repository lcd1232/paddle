package paddle

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient(Settings{
		URL:            SandboxBaseURL,
		VendorID:       "123",
		VendorAuthCode: "123abc",
	})
	require.NoError(t, err)
	require.NotNil(t, c)
	assert.NotNil(t, c.client)
	assert.NotNil(t, c.BaseURL)
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
	assert.Equal(t, userAgent, c.UserAgent)
}

func TestNewClientInvalidURL(t *testing.T) {
	_, err := NewClient(Settings{
		URL: "http\\:a",
	})
	require.Error(t, err)
}
