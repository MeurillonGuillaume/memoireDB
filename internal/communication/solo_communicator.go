package communication

type singleNodeCommunicator struct {
	cfg *Config
}

var _ NodeCommunicator = (*singleNodeCommunicator)(nil)

func (snc *singleNodeCommunicator) PingAllNodes() error {
	return nil
}

func newSingleNodeCommunicator(cfg *Config) NodeCommunicator {
	return &singleNodeCommunicator{
		cfg: cfg,
	}
}
