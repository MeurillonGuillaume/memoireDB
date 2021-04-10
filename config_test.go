package main

import (
	"testing"

	"github.com/MeurillonGuillaume/memoireDB/shared"
	"github.com/koding/multiconfig"
	"github.com/stretchr/testify/assert"
)

const (
	_myTestNodeName    = "my-node-name"
	_myTestClusterName = "my-cluster"
)

// TestConfigLoading will check if configuration loading works as intended.
func TestConfigLoading(t *testing.T) {
	var cfg Config
	configLoader := multiconfig.New()

	err := shared.SetEnvMap(shared.EnvMap{
		"CONFIG_ROLE_NODENAME":        _myTestNodeName,
		"CONFIG_ROLE_CLUSTERNAME":     _myTestClusterName,
		"CONFIG_ROLE_INITIALNODEROLE": "leader",
	})
	assert.NoError(t, err)
	assert.Nil(t, configLoader.Load(&cfg))
	assert.Nil(t, configLoader.Validate(&cfg))
	assert.Equal(t, _myTestNodeName, cfg.Role.NodeName)
}
