package config

import (
	"fmt"
	"os"
)

func GetEnvVarWithDefault(key, defaultValue string) string {
	val := defaultValue
	if v := os.Getenv(key); v != "" {
		val = v
	}
	return val
}

func GetReleaseMode() string {
	return GetEnvVarWithDefault("RELEASE_MODE", "debug")
}

func GetHost() string {
	return GetEnvVarWithDefault("BIND_HOST", "0.0.0.0")
}

func GetPort() string {
	return GetEnvVarWithDefault("BIND_PORT", "8080")
}

func GetBindHost() string {
	port := GetPort()
	host := GetHost()
	return fmt.Sprintf("%s:%s", host, port)
}

func GetLogColorEnabled() bool {
	enabled := true
	if v := os.Getenv("DISABLE_LOGGING_COLOR"); v != "" {
		enabled = false
	}
	return enabled
}
