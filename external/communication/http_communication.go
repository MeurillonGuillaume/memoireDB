package communication

type httpCommunicator struct {
	operationsChan chan interface{}
}

var _ ClientCommunicator = (*httpCommunicator)(nil)

func newHTTPCommunicator() ClientCommunicator {
	hc := httpCommunicator{}
	hc.init()
	return &hc
}

func (hc *httpCommunicator) init() {}

func (hc *httpCommunicator) Operation() <-chan interface{} {
	return hc.operationsChan
}
