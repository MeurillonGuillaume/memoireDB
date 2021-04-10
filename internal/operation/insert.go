package operation

import (
	"github.com/MeurillonGuillaume/memoireDB/external/communication/model"
	"github.com/MeurillonGuillaume/memoireDB/internal/datastore"
)

type insertOperation struct {
	data model.InsertModel
	done chan struct{}

	result interface{}
	err    error
}

var _ Operation = (*insertOperation)(nil)

func NewInsertOperation(m model.InsertModel) Operation {
	return &insertOperation{
		data: m,
		done: make(chan struct{}, 1),
	}
}

func (iop *insertOperation) Start(ds datastore.Store) {
	defer close(iop.done)

	if result, err := ds.StoreKeyValue(iop.data.Key, iop.data.Value); err != nil {
		iop.err = err
	} else {
		iop.result = result
	}
}

func (iop *insertOperation) Wait() { <-iop.done }

func (iop *insertOperation) Result() (interface{}, error) { return iop.result, iop.err }

func (iop *insertOperation) String() string { return _insertOperationName }
