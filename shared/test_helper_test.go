package shared

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSetEnvMap will ensure the SetEnvMap func is working as intended
func TestSetEnvMap(t *testing.T) {
	envs := EnvMap{
		"key":       "value",
		"other-key": "other-value",
		"1":         "2",
	}
	assert.Nil(t, SetEnvMap(envs))

	for key, value := range envs {
		envVal, ok := os.LookupEnv(key)
		assert.Equal(t, true, ok)
		assert.Equal(t, value, envVal)
	}
	os.Clearenv()
}
