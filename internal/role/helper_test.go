package role

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewNodeName will assure a new nodename is never empty
func TestNewNodeName(t *testing.T) {
	for i := 0; i < 100000; i++ {
		name := newNodeName()
		assert.NotEqual(t, "", name)
		assert.NotEqual(t, "", strings.TrimSpace(name))
		assert.Equal(t, false, len(name) < 1)
		assert.Equal(t, false, len(strings.TrimSpace(name)) < 1)
	}
}
