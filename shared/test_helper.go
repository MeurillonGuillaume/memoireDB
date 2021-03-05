package shared

import (
	"fmt"
	"os"
)

type (
	// EnvMap is a type of string-string map to simplify the applying of multiple key:value pairs for ENV
	EnvMap map[string]string
)

// SetEnvMap will assign multiple key-value pairs to the ENV variables
func SetEnvMap(m EnvMap) (err error) {
	if m == nil || len(m) < 1 {
		return fmt.Errorf("Could not apply EnvMap, is nil or has no content")
	}

	for envKey, envValue := range m {
		if err := os.Setenv(envKey, envValue); err != nil {
			return err
		}
	}
	return nil
}
