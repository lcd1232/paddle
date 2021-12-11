package paddle

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient(Settings{
		URL: SandboxBaseURL,
	})
	require.NoError(t, err)
	require.NotNil(t, c)
	assert.NotNil(t, c.client)
	assert.NotNil(t, c.BaseURL)
	assert.Equal(t, userAgent, c.UserAgent)
	assert.Zero(t, c.vendorID)
	assert.Zero(t, c.vendorAuthCode)
}
