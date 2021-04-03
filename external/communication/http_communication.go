package communication

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
)

// HttpCommunicatorConfig contains configuration parameters to set up a HTTP server
type HttpCommunicatorConfig struct {
	Port int `default:"8080" flagUsage:"Which port should be used to listen for incoming HTTP requests"`
}

type httpCommunicator struct {
	cfg            *HttpCommunicatorConfig
	operationsChan chan interface{}
	server         *http.Server
	wg             sync.WaitGroup
}

var _ ClientCommunicator = (*httpCommunicator)(nil)

func newHTTPCommunicator() (cc ClientCommunicator, err error) {
	var cfg HttpCommunicatorConfig
	configLoader := multiconfig.New()

	if err = configLoader.Load(&cfg); err != nil {
		return
	}

	if err = configLoader.Validate(&cfg); err != nil {
		return
	}

	server := new(http.Server)
	server.Addr = fmt.Sprintf(":%d", cfg.Port)

	cc = &httpCommunicator{
		operationsChan: make(chan interface{}, 1),
		cfg:            &cfg,
		server:         server,
	}
	return
}

func (hc *httpCommunicator) Run(ctx context.Context) {
	go func() {
		logrus.WithField("HTTP port", hc.cfg.Port).Info("HTTP Client communicator is ready to serve requests")
		if err := hc.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithError(err).Error("Can no longer serve external HTTP requests")
		} else {
			logrus.Info("Received gracefull shutdown command for external HTTP client communicator")
		}
	}()
	<-ctx.Done()
}

func (hc *httpCommunicator) Operation() <-chan interface{} { return hc.operationsChan }

func (hc *httpCommunicator) Close() error {
	if err := hc.server.Close(); err != nil {
		return err
	}

	hc.wg.Wait()
	close(hc.operationsChan)
	return nil
}
