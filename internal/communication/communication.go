package communication

// NodeCommunicator is an interface defining which methods should be implemented to create
// basic communication between nodes using a certain technology
type NodeCommunicator interface {
	// PingAllNodes makes a ping request to every known node to verify that all nodes are available
	PingAllNodes() error
}

// NewNodeCommunicator will create a new communicator to regulate communication between nodes
func NewNodeCommunicator(cfg *Config) (NodeCommunicator, error) {
	return nil, nil
}
