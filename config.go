package main

import (
	"github.com/google/uuid"
	"github.com/koding/multiconfig"
)

const (
	_systemName = "memoireDB"
	_nameSep    = "-"
)

// Config contains the configuration parameters used for memoireDB
type Config struct {
	NodeName    string `flagUsage:"The nodeName is used to identify unique nodes in a cluster setup."`
	ClusterName string `required:"true" flagUsage:"The clusterName is used to identify unity between one or more nodes. A unique cluster name will result in a unique store of data."`
	InitialRole string `required:"true" flagUsage:"The InitialRole defines what this node is designed to be from a cold start. A node can be guaranteed to become a follower, but not a leader. At least 1 leader is required, and the maximum is configurable."`
}

// loadConfig will either load configuration parameters from ENV or crash out
func loadConfig() (*Config, error) {
	var cfg Config
	configLoader := multiconfig.New()

	// Load config
	if err := configLoader.Load(&cfg); err != nil {
		return nil, err
	}

	// Validate flags
	if err := configLoader.Validate(&cfg); err != nil {
		return nil, err
	}

	cfg.setNodeName()
	return &cfg, nil
}

// setNodeName will create a new NodeName in the form of "memoireDB-UUID" if no nodename is given
func (cfg *Config) setNodeName() {
	if len(cfg.NodeName) < 1 {
		cfg.NodeName = _systemName + _nameSep + uuid.New().String()
	}
}
