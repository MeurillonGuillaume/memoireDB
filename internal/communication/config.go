package communication

// Config contains configuration parameters for
type Config struct {
	Channel string   `default:"gRPC" flagUsage:"Which communication method should be used to let nodes communicate with eachother"`
	Peers   []string `flagUsage:"Which other peers are there to connect with and form a cluster"`
}
