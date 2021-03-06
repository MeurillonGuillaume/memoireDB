package communication

// Config contains configuration parameters for
type Config struct {
	Peers []string `flagUsage:"Which other peers are there to connect with"`
}
