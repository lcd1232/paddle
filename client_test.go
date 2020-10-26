package paddle

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil)
	require.NotNil(t, c)
	assert.NotNil(t, c.client)
	assert.Zero(t, c.vendorID)
	assert.Zero(t, c.vendorAuthCode)
}

func TestNewClientWithAuthentication(t *testing.T) {
	c := NewClientWithAuthentication(nil, "id1", "auth1")
	require.NotNil(t, c)
	assert.NotNil(t, c.client)
	assert.Equal(t, "id1", c.vendorID)
	assert.Equal(t, "auth1", c.vendorAuthCode)
}
