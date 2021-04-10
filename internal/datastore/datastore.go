package datastore

import "github.com/MeurillonGuillaume/memoireDB/external/communication/model"

// Store is an interface declaring the functionality of a datastore for MemoireDB
type Store interface {
	// LoadKeyValue attempts to retrieve a single key:value pair from the datastore, or returns an ErrNoSuchKey if the key is unknown
	LoadKeyValue(m model.RetrieveModel) (interface{}, error)
	// StoreKeyValue attempts to store a single key:value pair in the datastore and returns the result and/or an error
	StoreKeyValue(m model.InsertModel) (interface{}, error)
	// ListKeys attempts to list all known keys starting with a certain prefix, or returns all keys if no prefix is given. If no results are available, an ErrNoSuchKey message is returned
	ListKeys(m model.ListKeysModel) ([]string, error)
}

func NewDatastore() (Store, error) {
	return newMemoryDatastore(), nil
}
