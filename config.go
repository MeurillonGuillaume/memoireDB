package main

import (
	externalCom "github.com/MeurillonGuillaume/memoireDB/external/communication"
	internalCom "github.com/MeurillonGuillaume/memoireDB/internal/communication"
	"github.com/MeurillonGuillaume/memoireDB/internal/role"
	"github.com/koding/multiconfig"
)

// Config contains the configuration parameters used for memoireDB
type Config struct {
	Role                  role.Config
	InternalCommunication internalCom.Config
	ExternalCommunication externalCom.Config
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

	cfg.Role.FillEmptyFields()
	return &cfg, nil
}
