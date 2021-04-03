package operation

// Operation is an interface which declares what a database operation should have as functionality
type Operation interface {
	Start()
	Wait()
	Result() (interface{}, error)
	String() string
}

const (
	_retrieveOperationName = "retrieve operation"
	_insertOperationName   = "insert operation"
)
