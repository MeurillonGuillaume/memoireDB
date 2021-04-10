package communication

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/MeurillonGuillaume/memoireDB/external/communication/helpers"
	"github.com/MeurillonGuillaume/memoireDB/external/communication/model"
	"github.com/MeurillonGuillaume/memoireDB/internal/operation"
	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
)

type (
	// HttpCommunicatorConfig contains configuration parameters to set up a HTTP server
	HttpCommunicatorConfig struct {
		Port int `default:"8080" flagUsage:"Which port should be used to listen for incoming HTTP requests"`
	}
	httpCommunicator struct {
		cfg            *HttpCommunicatorConfig
		operationsChan chan interface{}
		server         *http.Server
		wg             sync.WaitGroup
	}
)

const (
	_statusOk  = "Operation succeeded"
	_statusNok = "Operation failed"
)

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
		{
			Name:    "Insert key-value pair",
			Path:    "/insert",
			Methods: []string{http.MethodPost},
			Handler: hc.putHandler,
		},
		{
			Name:    "Retrieve key-value pair",
			Path:    "/retrieve",
			Methods: []string{http.MethodPost},
			Handler: hc.getHandler,
		},
		{
			Name:    "List keys with optional prefix",
			Path:    "/list",
			Methods: []string{http.MethodPost},
			Handler: hc.listHandler,
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

func (hc *httpCommunicator) putHandler(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()

	hc.wg.Add(1)
	defer func() {
		hc.wg.Done()
		if err := r.Body.Close(); err != nil {
			logrus.WithError(err).Error("Could not properly close request body")
		}
	}()

	var insertRequest model.InsertModel
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&insertRequest); err != nil {
		logrus.WithError(err).Error("Could not decode insert request data")
		helpers.HTTPReplyJSON(rw, http.StatusBadRequest, model.SimpleResponse{
			Result:  _statusNok,
			Message: "Retrieve failed",
			Error:   "Could not properly decode insert model",
			Took:    time.Since(start).Nanoseconds(),
		})
		return
	}

	// Create an operation and pass it to the internal handler
	op := operation.NewInsertOperation(insertRequest)
	hc.operationsChan <- op
	op.Wait()

	if result, err := op.Result(); err != nil {
		helpers.HTTPReplyJSON(rw, http.StatusBadRequest, model.SimpleResponse{
			Result:  _statusNok,
			Message: "Insert failed",
			Error:   err.Error(),
			Took:    time.Since(start).Nanoseconds(),
		})
	} else {
		helpers.HTTPReplyJSON(rw, http.StatusOK, model.SimpleResponse{
			Result:  _statusOk,
			Message: "Insert successful",
			Value:   result,
			Took:    time.Since(start).Nanoseconds(),
		})
	}
}

func (hc *httpCommunicator) getHandler(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()

	hc.wg.Add(1)
	defer func() {
		hc.wg.Done()
		if err := r.Body.Close(); err != nil {
			logrus.WithError(err).Error("Could not properly close request body")
		}
	}()

	var retrieveRequest model.RetrieveModel
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&retrieveRequest); err != nil {
		logrus.WithError(err).Error("Could not decode retrieve request data")
		helpers.HTTPReplyJSON(rw, http.StatusBadRequest, model.SimpleResponse{
			Result:  _statusOk,
			Message: "Retrieve failed",
			Error:   "Could not properly decode retrieve model",
			Took:    time.Since(start).Nanoseconds(),
		})
		return
	}

	op := operation.NewRetrieveOperation(retrieveRequest)
	hc.operationsChan <- op
	op.Wait()

	if result, err := op.Result(); err != nil {
		helpers.HTTPReplyJSON(rw, http.StatusBadRequest, model.SimpleResponse{
			Result:  _statusNok,
			Message: "Retrieve failed",
			Error:   err.Error(),
			Took:    time.Since(start).Nanoseconds(),
		})
	} else {
		helpers.HTTPReplyJSON(rw, http.StatusOK, model.RetrieveResponse{
			Key:   retrieveRequest.Key,
			Value: result,
			Took:  time.Since(start).Nanoseconds(),
		})
	}
}

func (hc *httpCommunicator) listHandler(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()

	hc.wg.Add(1)
	defer func() {
		hc.wg.Done()
		if err := r.Body.Close(); err != nil {
			logrus.WithError(err).Error("Could not properly close request body")
		}
	}()

	var listModel model.ListKeyModel
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&listModel); err != nil && err != io.EOF {
		logrus.WithError(err).Error("Could not decode list request")
		helpers.HTTPReplyJSON(rw, http.StatusBadRequest, model.SimpleResponse{
			Result:  _statusOk,
			Message: "List keys failed",
			Error:   "Could not properly decode list key(s) model",
			Took:    time.Since(start).Nanoseconds(),
		})
		return
	}

	op := operation.NewListOperation(listModel)
	hc.operationsChan <- op
	op.Wait()

	if result, err := op.Result(); err != nil {
		helpers.HTTPReplyJSON(rw, http.StatusBadRequest, model.SimpleResponse{
			Result:  _statusNok,
			Message: "List keys failed",
			Error:   err.Error(),
			Took:    time.Since(start).Nanoseconds(),
		})
	} else {
		helpers.HTTPReplyJSON(rw, http.StatusOK, model.ListKeysResponse{
			Prefix: listModel.Prefix,
			Keys:   result,
			Took:   time.Since(start).Nanoseconds(),
		})
	}
}
