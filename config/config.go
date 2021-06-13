package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
)

const (
	PrefixMemoireDB = "memoiredb"
	_lineSepStart   = "---- START ----"
	_lineSepEnd     = "----- END -----"
	_tab            = "    "
	_childArrow     = "↳"
	_redaction      = "[ REDACTED ]"

	_fieldSecret   = "secret"
	_fieldDefault  = "default"
	_fieldRequired = "required"
	_fieldHelp     = "help"
)

// LoadFromEnv will load configuration parameters with a specific prefix from the ENV variables and attempt to Load & Validate these parameters based on the given configuration interface.
// The ENV prefix = given prefix string + "_" + the name of the configuration struct.
func LoadFromEnv(prefix string, cfg interface{}) error {
	cfgLoader := multiconfig.New()
	cfgLoader.Loader = &EnvLoader{
		Prefix: prefix + "_" + getInterfaceName(cfg),
	}

	// Load config
	if err := cfgLoader.Load(cfg); err != nil {
		return err
	}

	// Validate flags
	if err := cfgLoader.Validate(cfg); err != nil {
		return err
	}

	logConfig(prefix, cfg)
	return nil
}

// getInterfaceName will return the name of the given interface.
func getInterfaceName(i interface{}) string {
	t := reflect.TypeOf(i)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

func logConfig(prefix string, cfg interface{}) {
	logrus.Info(_lineSepStart)
	// Interface name logging
	e := reflect.ValueOf(cfg).Elem()
	logrus.Infof("%s_%s:", strings.ToUpper(prefix), strings.ToUpper(e.Type().Name()))

	// Log sublevels of this config
	logSubLevels(e, 0)
	logrus.Info(_lineSepEnd)
}

func logSubLevels(rv reflect.Value, level int) {
	t := rv.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !rv.Field(i).CanInterface() {
			continue
		}
		value := rv.Field(i).Interface()

		// Log nested config
		if field.Type.Kind() == reflect.Struct {
			logrus.Infof("%v%v%v:", strings.Repeat(_tab, level+1), _childArrow, field.Name)
			logSubLevels(reflect.ValueOf(value), level+1)
			continue
		}

		// Hide sensitive data from logging
		if secret := field.Tag.Get(_fieldSecret); secret == "true" {
			value = _redaction
		}

		// Set base metadata
		meta := fmt.Sprintf("(type: %v", field.Type)
		// Fetch required
		if value := field.Tag.Get(_fieldRequired); value != "" {
			meta += fmt.Sprintf(", required: %s", value)
		}

		// Fetch default value if any
		if value := field.Tag.Get(_fieldDefault); value != "" {
			meta += fmt.Sprintf(", default: %s", value)
		}

		// Fetch help if any
		if value := field.Tag.Get(_fieldHelp); value != "" {
			meta += fmt.Sprintf(", description: %s", value)
		}
		meta += ")"

		logrus.Infof("%s%s%s %s: %v", strings.Repeat(_tab, level+1), _childArrow, field.Name, meta, value)
	}
}
