package role

import "strings"

// Config defines global configuration for NodeRoles.
type Config struct {
	InitialNodeRole string `default:"leader" flagUsage:"Defines which initial role a node has in the cluster. Can be either 'leader' or 'follower'."`
	NodeName        string `flagUsage:"The nodeName is used to identify unique nodes in a cluster setup."`
	ClusterName     string `default:"memoireDB-cluster-default" flagUsage:"The clusterName is used to identify unity between one or more nodes. A unique cluster name will result in a unique store of data."`
}

// FillEmptyFields will fill all empty fields which can be filled automatically.
func (cfg *Config) FillEmptyFields() {
	// Set default ClusterName if none
	if len(strings.TrimSpace(cfg.ClusterName)) < 1 {
		cfg.ClusterName = _defaultClusterName
	}

	// Create unique NodeName if none
	if len(strings.TrimSpace(cfg.NodeName)) < 1 {
		cfg.NodeName = newNodeName()
	}
}
