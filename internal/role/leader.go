package role

// leader is a type of NodeRole which is assigned to orchestrate what happens in the cluster
type leader struct {
	cfg         *Config
	currentRole RoleType
}

var _ NodeRole = (*leader)(nil)

func newLeaderNode(cfg *Config) NodeRole {
	l := leader{
		cfg: cfg,
	}
	l.init()
	return &l
}

func (l *leader) init() {
	l.currentRole = RoleType(l.cfg.InitialNodeRole)
}

func (l *leader) GetRole() RoleType {
	return l.currentRole
}

func (l *leader) GetName() string {
	return l.cfg.NodeName
}

func (l *leader) GetCluster() string {
	return l.cfg.ClusterName
}
