package config

import (
	"log"
	"os"
)

func Get(key EnvKey) string {
	value, ok := os.LookupEnv(string(key))

	if !ok {
		log.Panicf("environment variable \"%s\" is empty", key)
	}

	return value
}

func GetOrDefault(key EnvKey, def string) string {
	value, ok := os.LookupEnv(string(key))

	if !ok {
		return def
	}

	return value
}
