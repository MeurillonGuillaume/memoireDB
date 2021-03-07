package communication

import (
	"testing"

	"github.com/MeurillonGuillaume/memoireDB/shared"
	"github.com/stretchr/testify/assert"
)

func TestNewCommunicator(t *testing.T) {
	// Create gRPC communicator
	var (
		nc  NodeCommunicator
		err error
	)

	nc, err = NewNodeCommunicator(&Config{
		Channel: CommunicatorGrpc,
		Peers:   []string{"1", "2"},
	})
	assert.Nil(t, err)
	assert.NotNil(t, nc)

	// Create a lot of invalid cases using invalid communicator types
	for i := 0; i < 10000; i++ {
		nc, err = NewNodeCommunicator(&Config{
			// Generate 10000 random strings between 5 - 25 characters
			Channel: shared.NewRandomString((i % 25) + 5),
			Peers:   []string{"1", "2"},
		})
		assert.NotNil(t, err)
		assert.Nil(t, nc)
	}
}
