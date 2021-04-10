package communication

import (
	"context"
	"fmt"
	"strings"
)

// ClientCommunicator is an interface declaring the functionality of a communicator for client - server communication.
type ClientCommunicator interface {
	// Run will keep the ClientCommunicator alive as long as the context is unclosed
	Run(ctx context.Context)
	// Operation will expose the internal channel for operations
	Operation() <-chan interface{}
	// Close will close the clientcommunicator and handle all operations in state
	Close() error
}

// NewClientCommunicators will initialize all configured Client Communicators.
func NewClientCommunicators(cfg *Config) (cc []ClientCommunicator, err error) {
	if len(cfg.Methods) < 1 {
		err = fmt.Errorf("at least one external communication method is required")
		return
	}

	for _, method := range cfg.Methods {
		switch strings.ToLower(method) {
		case MethodHTTPCommunicator:
			hc, err := newHTTPCommunicator()
			if err != nil {
				return cc, err
			}
			cc = append(cc, hc)
		default:
			err = fmt.Errorf("could not configure ClientCommunicator of unknown type %s", method)
			return
		}
	}
	return
}

// GetCommunicatorChans will loop over a slice of communicators and return a slice of their Operation channels.
func GetCommunicatorChans(cc []ClientCommunicator) []<-chan interface{} {
	clientStream := make([]<-chan interface{}, 0, len(cc))
	for _, c := range cc {
		clientStream = append(clientStream, c.Operation())
	}
	return clientStream
}
