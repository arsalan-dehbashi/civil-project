package config

import (
	"os"
	"strconv"
)

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getEnvAsInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}

func getEnvAsBool(key string, def bool) bool {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.ParseBool(v); err == nil {
			return n
		}
	}
	return def
}
