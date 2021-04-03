package operation

import (
	"time"

	"github.com/MeurillonGuillaume/memoireDB/external/communication/model"
)

type (
	// Operation is an interface which declares what a database operation should have as functionality
	Operation interface {
		Start()
		Wait()
		Result() error
	}

	insertOperation struct {
		done chan struct{}
	}
)

var _ Operation = (*insertOperation)(nil)

func NewInsertOperation(m model.InsertModel) Operation {
	return &insertOperation{
		done: make(chan struct{}, 1),
	}
}

func (iop *insertOperation) Start() {
	defer close(iop.done)
	time.Sleep(time.Second)
}

func (iop *insertOperation) Wait() { <-iop.done }

func (iop *insertOperation) Result() error { return nil }
