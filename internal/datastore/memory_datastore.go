package datastore

import (
	"sort"
	"strings"
	"sync"

	"github.com/MeurillonGuillaume/memoireDB/external/communication/model"
)

type memoryDatastore struct {
	storage map[string]interface{}
	mux     sync.RWMutex
}

var _ Store = (*memoryDatastore)(nil)

func newMemoryDatastore() Store {
	return &memoryDatastore{
		storage: make(map[string]interface{}),
		mux:     sync.RWMutex{},
	}
}

func (md *memoryDatastore) LoadKeyValue(m model.RetrieveModel) (interface{}, error) {
	md.mux.RLock()
	r, ok := md.storage[m.Key]
	md.mux.RUnlock()

	if !ok {
		return nil, ErrNoSuchKey
	} else {
		return r, nil
	}
}

func (md *memoryDatastore) StoreKeyValue(m model.InsertModel) (interface{}, error) {
	md.mux.Lock()
	md.storage[m.Key] = m.Value
	md.mux.Unlock()
	return m.Value, nil
}

func (md *memoryDatastore) ListKeys(m model.ListKeysModel) ([]string, error) {
	// Check if the prefix contains nothing but whitespace
	if len(m.Prefix) > 0 && len(strings.TrimSpace(m.Prefix)) == 0 {
		return nil, ErrPrefixWhitespace
	}

	// Execute query
	var result []string
	md.mux.RLock()
	defer md.mux.RUnlock()

	if len(m.Prefix) > 0 {
		for key := range md.storage {
			if strings.HasPrefix(key, m.Prefix) {
				result = append(result, key)
			}
		}
	} else {
		for key := range md.storage {
			result = append(result, key)
		}
	}

	if len(result) < 1 {
		return nil, ErrNoSuchKey
	}
	sort.Strings(result)
	return result, nil
}

func (md *memoryDatastore) Close() error {
	return nil
}
