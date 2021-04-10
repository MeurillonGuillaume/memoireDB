package role

// RoleType is a String-type specifically used for Node RoleTypes.
type RoleType string

// Public consts.
const (
	// ClusterLeader is a type of NodeRole defining that this node should be used as leader under ideal circumstances.
	ClusterLeader = RoleType("leader")
	// ClusterFollower is a type of NodeRole defining that this node should be used as follower under ideal circumstances.
	ClusterFollower = RoleType("follower")
)

// Local consts.
const (
	_systemName         = "memoireDB"
	_nameSep            = "-"
	_default            = "default"
	_node               = "node"
	_defaultClusterName = _systemName + _nameSep + _default
)
