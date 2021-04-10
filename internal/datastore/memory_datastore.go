package datastore

import "sync"

type memoryDatastore struct {
	store map[string]interface{}
	mux   sync.RWMutex
}

var _ Store = (*memoryDatastore)(nil)

func newMemoryDatastore() Store {
	return &memoryDatastore{
		store: make(map[string]interface{}),
		mux:   sync.RWMutex{},
	}
}

func (md *memoryDatastore) LoadKeyValue(key string) (interface{}, error) {
	md.mux.RLock()
	defer md.mux.RUnlock()

	if r, ok := md.store[key]; !ok {
		return nil, ErrNoSuchKey
	} else {
		return r, nil
	}
}

func (md *memoryDatastore) StoreKeyValue(key string, value interface{}) (interface{}, error) {
	md.mux.Lock()
	defer md.mux.Unlock()

	md.store[key] = value
	return value, nil
}
