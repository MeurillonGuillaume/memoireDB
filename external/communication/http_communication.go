package communication

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/MeurillonGuillaume/memoireDB/external/communication/helpers"
	"github.com/MeurillonGuillaume/memoireDB/external/communication/model"
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

	hc := httpCommunicator{
		operationsChan: make(chan interface{}, 1),
		cfg:            &cfg,
	}

	server := new(http.Server)
	server.Addr = fmt.Sprintf(":%d", cfg.Port)
	hc.server = helpers.NewHTTPServer(cfg.Port, hc.getRoutes())
	return &hc, nil
}

func (hc *httpCommunicator) Run(ctx context.Context) {
	go func() {
		logrus.WithField("HTTP port", hc.cfg.Port).Info("HTTP Client communicator is ready to serve requests")
		if err := hc.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithError(err).Error("Can no longer serve external HTTP requests")
		} else {
			logrus.Info("Received graceful shutdown command for external HTTP client communicator")
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

func (hc *httpCommunicator) getRoutes() []helpers.Route {
	return []helpers.Route{
		{
			Name:    "Cluster Status",
			Path:    "/node/status",
			Methods: []string{http.MethodGet},
			Handler: hc.statusHandler,
		},
	}
}

// statusHandler is a simple HTTP responsewriter to display the current status to a requester
func (hc *httpCommunicator) statusHandler(rw http.ResponseWriter, r *http.Request) {
	hc.wg.Add(1)
	defer hc.wg.Done()

	helpers.HTTPReplyJSON(rw, http.StatusOK, model.StatusResponse{
		Message: "I'm online!",
	})
}
