package operation

import (
	"github.com/MeurillonGuillaume/memoireDB/external/communication/model"
	"github.com/MeurillonGuillaume/memoireDB/internal/datastore"
	"github.com/google/uuid"
)

type retrieveOperation struct {
	data model.RetrieveModel
	done chan struct{}
	id   uuid.UUID

	result interface{}
	err    error
}

var _ Operation = (*retrieveOperation)(nil)

func NewRetrieveOperation(id uuid.UUID, m model.RetrieveModel) Operation {
	return &retrieveOperation{
		data: m,
		id:   id,
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
