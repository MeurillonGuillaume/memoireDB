package main

import (
	"testing"

	"github.com/MeurillonGuillaume/memoireDB/config"
	"github.com/MeurillonGuillaume/memoireDB/shared"
	"github.com/stretchr/testify/assert"
)

// TestConfigLoading will check if configuration loading works as intended.
func TestConfigLoading(t *testing.T) {
	var cfg BaseConfig
	err := shared.SetEnvMap(shared.EnvMap{
		"MEMOIREDB_BASECONFIG_DATASTORE_TYPE": "memory",
	})
	assert.NoError(t, err)
	assert.Nil(t, config.LoadFromEnv(config.PrefixMemoireDB, &cfg))
}
