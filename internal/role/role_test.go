package role

import (
	"testing"

	"github.com/MeurillonGuillaume/memoireDB/shared"
	"github.com/stretchr/testify/assert"
)

func TestNewNodeRole(t *testing.T) {
	var (
		node NodeRole
		err  error
	)

	// Create a new leader
	node, err = NewNodeWithRole(&Config{
		InitialNodeRole: string(ClusterLeader),
	})
	assert.Nil(t, err)
	assert.NotNil(t, node)

	// Create a new follower
	node, err = NewNodeWithRole(&Config{
		InitialNodeRole: string(ClusterFollower),
	})
	assert.Nil(t, err)
	assert.NotNil(t, node)

	// Create a lot of invalid cases using random node role types
	for i := 0; i < 10000; i++ {
		node, err = NewNodeWithRole(&Config{
			// Generate 10000 random strings between 5 - 25 characters
			InitialNodeRole: shared.NewRandomString((i % 25) + 5),
		})
		assert.NotNil(t, err)
		assert.Nil(t, node)
	}
}
