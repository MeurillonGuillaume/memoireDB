package communication

// Config contains configuration parameters for the internal node communicator(s).
type Config struct {
	Channel string   `default:"gRPC" help:"Which communication method should be used to let nodes communicate with eachother"`
	Peers   []string `help:"Which other peers are there to connect with and form a cluster"`
}
