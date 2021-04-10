package operation

import (
	"github.com/MeurillonGuillaume/memoireDB/external/communication/model"
	"github.com/MeurillonGuillaume/memoireDB/internal/datastore"
)

type retrieveOperation struct {
	data model.RetrieveModel
	done chan struct{}

	result interface{}
	err    error
}

var _ Operation = (*retrieveOperation)(nil)

func NewRetrieveOperation(m model.RetrieveModel) Operation {
	return &retrieveOperation{
		data: m,
		done: make(chan struct{}, 1),
	}
}

func (rop *retrieveOperation) Start(ds datastore.Store) {
	defer close(rop.done)

	if value, err := ds.LoadKeyValue(rop.data); err != nil {
		rop.err = err
	} else {
		rop.result = value
	}
}

func (rop *retrieveOperation) Wait() { <-rop.done }

func (rop *retrieveOperation) Result() (interface{}, error) { return rop.result, rop.err }

func (rop *retrieveOperation) String() string { return _retrieveOperationName }
