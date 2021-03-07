package communication

// Config contains configuration parameters to configure external communication.
// External communication in this case means client - server communication.
type Config struct {
	Methods []string `flagUsage:"Use sources to configure one or multiple available client-server communication protocols."`
}
