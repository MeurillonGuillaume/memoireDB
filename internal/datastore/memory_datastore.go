package datastore

import (
	"strings"
	"sync"

	"github.com/MeurillonGuillaume/memoireDB/external/communication/model"
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

func (md *memoryDatastore) LoadKeyValue(m model.RetrieveModel) (interface{}, error) {
	md.mux.RLock()
	defer md.mux.RUnlock()

	if r, ok := md.store[m.Key]; !ok {
		return nil, ErrNoSuchKey
	} else {
		return r, nil
	}
}

func (md *memoryDatastore) StoreKeyValue(m model.InsertModel) (interface{}, error) {
	md.mux.Lock()
	defer md.mux.Unlock()

	md.store[m.Key] = m.Value
	return m.Value, nil
}

func (md *memoryDatastore) ListKeys(m model.ListKeysModel) ([]string, error) {
	md.mux.RLock()
	defer md.mux.RUnlock()

	// Check if the prefix contains nothing but whitespace
	if len(m.Prefix) > 0 && len(strings.TrimSpace(m.Prefix)) == 0 {
		return nil, ErrPrefixWhitespace
	}

	// Execute query
	var result []string
	if len(m.Prefix) > 0 {
		for key := range md.store {
			if strings.HasPrefix(key, m.Prefix) {
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
