package datastore

// Store is an interface declaring the functionality of a datastore for MemoireDB
type Store interface {
	StoreKeyValue(key string, value interface{})
	LoadKeyValue(key string, value interface{})
}
