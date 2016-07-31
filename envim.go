package envim

import (
	"os"
	"strings"
)

// Set sets environment variable named key new value.
// Returns an error if occurred.
func Set(key, value string) error {
	return os.Setenv(key, value)
}

// Unset unsets an environment variable
// Returns an error if occurred.
func Unset(key string) error {
	return os.Unsetenv(key)
}

// Get gets value of environment variable named key
// Returns empty string if variable is not set
func Get(key string) string {
	return os.Getenv(key)
}

// IsSet check variable is set or not
func IsSet(key string) bool {
	_, isSet := os.LookupEnv(key)

	return isSet
}

// Clear cleans environment
func Clear() {
	os.Clearenv()
}

// Map returns a map contains environment variable key-value pair
func Map() map[string]string {
	m := make(map[string]string)

	for _, s := range os.Environ() {
		t := strings.SplitN(s, "=", 2)
		m[t[0]] = t[1]
	}

	return m
}

// MapWithPrefix like Map(), but only contains key start with prefix
func MapWithPrefix(prefix string) map[string]string {
	m := Map()
	mWithPrefix := make(map[string]string)

	for k, v := range m {
		if strings.HasPrefix(k, prefix) {
			mWithPrefix[k] = v
		}
	}

	return mWithPrefix
}

// FromMap populates environment with key-value pair from map m
// A pair with invalid key/value will be skipped
func FromMap(m map[string]string) {
	for k, v := range m {
		_ = Set(k, v)
	}
}
