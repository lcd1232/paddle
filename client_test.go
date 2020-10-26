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
}
