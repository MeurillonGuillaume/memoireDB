package communication

import "github.com/MeurillonGuillaume/memoireDB/internal/operation"

type httpCommunicator struct {
	operationsChan chan operation.Operation
}

var _ ClientCommunicator = (*httpCommunicator)(nil)

func newHTTPCommunicator() ClientCommunicator {
	hc := httpCommunicator{}
	hc.init()
	return &hc
}

func (hc *httpCommunicator) init() {}

func (hc *httpCommunicator) Operations() <-chan operation.Operation {
	return hc.operationsChan
}
