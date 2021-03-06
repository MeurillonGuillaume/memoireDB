package role

import "github.com/google/uuid"

// newNodeName returns a new NodeName uniquely identified with an UUID
func newNodeName() string {
	return _systemName + _nameSep + _node + _nameSep + uuid.NewString()
}
