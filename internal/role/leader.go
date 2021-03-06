package role

// leader is a type of NodeRole which is assigned to orchestrate what happens in the cluster
type leader struct {
	initialConfig *Config
	currentRole   RoleType
}

var _ NodeRole = (*leader)(nil)

func newLeaderNode(cfg *Config) NodeRole {
	l := leader{
		initialConfig: cfg,
	}
	l.init()
	return &l
}

func (l *leader) init() {
	l.currentRole = RoleType(l.initialConfig.InitialNodeRole)
}

func (l *leader) Role() RoleType {
	return l.currentRole
}
