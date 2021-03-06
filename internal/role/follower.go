package role

// follower is a type of NodeRole which is assigned to listen and follow the leader
type follower struct {
	cfg         *Config
	currentRole RoleType
}

var _ NodeRole = (*follower)(nil)

func newFollowerNode(cfg *Config) NodeRole {
	f := follower{
		cfg: cfg,
	}
	f.init()
	return &f
}

func (f *follower) init() {
	f.currentRole = RoleType(f.cfg.InitialNodeRole)
}

func (f *follower) GetRole() RoleType {
	return f.currentRole
}

func (f *follower) GetName() string {
	return f.cfg.NodeName
}

func (f *follower) GetCluster() string {
	return f.cfg.ClusterName
}
