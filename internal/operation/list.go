package operation

import (
	"github.com/MeurillonGuillaume/memoireDB/external/communication/model"
	"github.com/MeurillonGuillaume/memoireDB/internal/datastore"
	"github.com/google/uuid"
)

type listKeysOperation struct {
	data model.ListKeysModel
	done chan struct{}
	id   uuid.UUID

	result interface{}
	err    error
}

var _ Operation = (*listKeysOperation)(nil)

func NewListOperation(id uuid.UUID, m model.ListKeysModel) Operation {
	return &listKeysOperation{
		data: m,
		id:   id,
		done: make(chan struct{}, 1),
	}
}

func (lkop *listKeysOperation) Start(ds datastore.Store) {
	defer close(lkop.done)

	if result, err := ds.ListKeys(lkop.data); err != nil {
		lkop.err = err
	} else {
		lkop.result = result
	}
}

func (lkop *listKeysOperation) Wait() { <-lkop.done }

func (lkop *listKeysOperation) Result() (interface{}, error) { return lkop.result, lkop.err }

func (lkop *listKeysOperation) String() string { return _listKeysOperationName }
