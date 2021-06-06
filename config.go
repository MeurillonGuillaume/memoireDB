package main

import (
	"github.com/MeurillonGuillaume/memoireDB/config"
	externalCom "github.com/MeurillonGuillaume/memoireDB/external/communication"
	internalCom "github.com/MeurillonGuillaume/memoireDB/internal/communication"
	"github.com/MeurillonGuillaume/memoireDB/internal/datastore"
)

// BaseConfig contains the configuration parameters used for memoireDB.
type BaseConfig struct {
	InternalCommunication internalCom.Config
	ExternalCommunication externalCom.Config
	Datastore             datastore.Config
}

// loadConfig will either load configuration parameters from ENV or crash out.
func loadConfig() (cfg BaseConfig, err error) {
	if err := config.LoadFromEnv(config.PrefixMemoireDB, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
