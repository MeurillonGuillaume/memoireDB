package communication

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// NodeCommunicator is an interface defining which methods should be implemented to create
// basic communication between nodes using a certain technology.
type NodeCommunicator interface {
	// PingAllNodes makes a ping request to every known node to verify that all nodes are available
	PingAllNodes() error
}

// NewNodeCommunicator will create a new communicator to regulate communication between nodes.
func NewNodeCommunicator(cfg *Config) (nc NodeCommunicator, err error) {
	if len(cfg.Peers) < 1 {
		logrus.WithField("peers", cfg.Peers).Warn("Received no peer addresses, starting cluster in Single Node mode")
		nc = newSingleNodeCommunicator(cfg)
		return
	}

	switch strings.ToLower(cfg.Channel) {
	case CommunicatorGrpc:
		nc = newGrpcCommunicator(cfg)
	default:
		err = fmt.Errorf("could not create NodeCommunicator with unknown type %s", cfg.Channel)
	}
	return
}
