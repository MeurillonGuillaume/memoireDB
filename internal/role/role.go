package role

import "fmt"

// NodeRole defines which functionality a Role should implement to create a functional cluster
type NodeRole interface {
	Role() RoleType
}

// NewNodeWithRole creates a new NodeRole from given configuration
func NewNodeWithRole(cfg *Config) (nr NodeRole, err error) {
	switch RoleType(cfg.InitialNodeRole) {
	case ClusterLeader:
		nr = newLeaderNode(cfg)
	case ClusterFollower:
		nr = newFollowerNode(cfg)
	default:
		err = fmt.Errorf("could not create noderole with role %s: no such role", cfg.InitialNodeRole)
	}
	return
}
