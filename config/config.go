package config

import (
	"os"
)

// GetEnv fetches an environment variable string specified by key
// If the environment variable doesn't exist, returns fallback
func GetEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

type Config struct {
	Environment string // Used to track environment
	CommitHash  string // Used to track release errors in Sentry
}

// To be used by other packages to access the config variables
var ConfigSingleton *Config = nil

func New() *Config {
	if ConfigSingleton == nil {
		ConfigSingleton = &Config{
			Environment: GetEnv("NAMESPACE", "local"),
			CommitHash:  GetEnv("COMMIT_HASH", "unknown"),
		}
	}
	return ConfigSingleton
}
