package datastore

import "github.com/MeurillonGuillaume/memoireDB/external/communication/model"

type (
	walDatastore struct{}
)

var _ Store = (*walDatastore)(nil)

func newWalDatastore() (Store, error) {
	return nil, nil
}

// LoadKeyValue attempts to retrieve a single key:value pair from the datastore, or returns an ErrNoSuchKey if the key is unknown.
func (wsd *walDatastore) LoadKeyValue(m model.RetrieveModel) (interface{}, error) {
	return nil, nil
}

// StoreKeyValue attempts to store a single key:value pair in the datastore and returns the result and/or an error.
func (wsd *walDatastore) StoreKeyValue(m model.InsertModel) (interface{}, error) {
	return nil, nil
}

// ListKeys attempts to list all known keys starting with a certain prefix, or returns all keys if no prefix is given. If no results are available, an ErrNoSuchKey message is returned.
func (wsd *walDatastore) ListKeys(m model.ListKeysModel) ([]string, error) {
	return nil, nil
}

func (wsd *walDatastore) Close() error {
	return nil
}
