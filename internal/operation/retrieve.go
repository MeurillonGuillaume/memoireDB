package operation

import (
	"time"

	"github.com/MeurillonGuillaume/memoireDB/external/communication/model"
)

type retrieveOperation struct {
	done chan struct{}
}

var _ Operation = (*retrieveOperation)(nil)

func NewRetrieveOperation(m model.RetrieveModel) Operation {
	return &retrieveOperation{
		done: make(chan struct{}, 1),
	}
}

func (rop *retrieveOperation) Start() {
	defer close(rop.done)
	time.Sleep(time.Second)
}

func (rop *retrieveOperation) Wait() { <-rop.done }

func (rop *retrieveOperation) Result() (interface{}, error) { return nil, nil }

func (rop *retrieveOperation) String() string { return _retrieveOperationName }
