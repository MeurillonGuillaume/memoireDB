package operation

import "github.com/MeurillonGuillaume/memoireDB/internal/datastore"

// Operation is an interface which declares what a database operation should have as functionality
type Operation interface {
	Start(ds datastore.Store)
	Wait()
	Result() (interface{}, error)
	String() string
}

const (
	_retrieveOperationName = "retrieve operation"
	_insertOperationName   = "insert operation"
)
