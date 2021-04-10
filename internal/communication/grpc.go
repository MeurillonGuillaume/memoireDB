package communication

// grpcCommunicator is a NodeCommunicator which uses the gRPC technology as communication backbone.
type grpcCommunicator struct {
	cfg *Config
}

var _ NodeCommunicator = (*grpcCommunicator)(nil)

func (gc *grpcCommunicator) PingAllNodes() error {
	return nil
}

func newGrpcCommunicator(cfg *Config) NodeCommunicator {
	return &grpcCommunicator{
		cfg: cfg,
	}
}
