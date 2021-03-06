package communication

// grpcCommunicator is a NodeCommunicator which uses the gRPC technology as communication backbone
type grpcCommunicator struct{}

var _ NodeCommunicator = (*grpcCommunicator)(nil)

func (gc *grpcCommunicator) PingAllNodes() error {
	return nil
}
