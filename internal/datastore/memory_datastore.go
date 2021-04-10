package datastore

import (
	"strings"
	"sync"
)

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

func (md *memoryDatastore) ListKeys(prefix string) ([]string, error) {
	md.mux.RLock()
	defer md.mux.RUnlock()

	// Check if the prefix contains nothing but whitespace
	if len(prefix) > 0 && len(strings.TrimSpace(prefix)) == 0 {
		return nil, ErrPrefixWhitespace
	}

	// Execute query
	var result []string
	if len(prefix) > 0 {
		for key := range md.store {
			if strings.HasPrefix(key, prefix) {
				result = append(result, key)
			}
		}
	} else {
		for key := range md.store {
			result = append(result, key)
		}
	}

	if len(result) < 1 {
		return nil, ErrNoSuchKey
	}
	return result, nil
}
