package communication

import (
	"fmt"
	"strings"
)

// ClientCommunicator is an interface declaring the functionality of a communicator for client - server communication
type ClientCommunicator interface {
	Operation() <-chan interface{}
}

// NewClientCommunicators will initialize all configured Client Communicators
func NewClientCommunicators(cfg *Config) (cc []ClientCommunicator, err error) {
	if len(cfg.Methods) < 1 {
		err = fmt.Errorf("at least one external communication method is required")
		return
	}

	for _, method := range cfg.Methods {
		switch strings.ToLower(method) {
		case MethodHTTPCommunicator:
			cc = append(cc, newHTTPCommunicator())
		default:
			err = fmt.Errorf("could not configure ClientCommunicator of unknown type %s", method)
			return
		}
	}
	return
}

// GetCommunicatorChans will loop over a slice of communicators and return a slice of their Operation channels
func GetCommunicatorChans(cc []ClientCommunicator) []<-chan interface{} {
	clientStream := make([]<-chan interface{}, 0, len(cc))

	for _, c := range cc {
		clientStream = append(clientStream, c.Operation())
	}

	return clientStream
}
