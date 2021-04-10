package shared

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
)

type (
	// EnvMap is a type of string-string map to simplify the applying of multiple key:value pairs for ENV.
	EnvMap map[string]string
)

// SetEnvMap will assign multiple key-value pairs to the ENV variables.
func SetEnvMap(m EnvMap) (err error) {
	if m == nil || len(m) < 1 {
		return fmt.Errorf("could not apply EnvMap, is nil or has no content")
	}

	for envKey, envValue := range m {
		if err := os.Setenv(envKey, envValue); err != nil {
			return err
		}
	}
	return nil
}

// NewRandomString will generate a random string using A-Z a-z 0-9 with given length.
func NewRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
