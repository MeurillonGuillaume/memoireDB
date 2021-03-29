package communication

import (
	"net/http"
	"sync"
)

type httpCommunicator struct {
	operationsChan chan interface{}
	server         *http.Server
	wg             sync.WaitGroup
}

var _ ClientCommunicator = (*httpCommunicator)(nil)

func newHTTPCommunicator() ClientCommunicator {
	hc := httpCommunicator{
		operationsChan: make(chan interface{}, 1),
	}
	hc.init()
	return &hc
}

func (hc *httpCommunicator) init() {
	hc.server = new(http.Server)

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
