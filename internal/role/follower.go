package role

// follower is a type of NodeRole which is assigned to listen and follow the leader
type follower struct {
	initialConfig *Config
	currentRole   RoleType
}

var _ NodeRole = (*follower)(nil)

func newFollowerNode(cfg *Config) NodeRole {
	f := follower{
		initialConfig: cfg,
	}
	f.init()
	return &f
}

func (f *follower) init() {
	f.currentRole = RoleType(f.initialConfig.InitialNodeRole)
}

func (f *follower) Role() RoleType {
	return f.currentRole
}
