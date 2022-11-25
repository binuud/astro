package config

// binu@dronasys.com

import (
	"fmt"
	"os"
	"strconv"
)

type EnvConfig struct {
	MongoEndpoint          string
	MongoDB                string
	MongoUser              string
	MongoPassword          string
	MongoEncryptionKey     string
	Production             string
	LogVerbosity           int
	Version                string
	AstroEphimerisEndpoint string
}

var Config *EnvConfig

func NewConfig() *EnvConfig {
	if Config != nil {
		return Config
	}

	config := EnvConfig{}

	config.MongoEndpoint = os.Getenv("MONGO_ENDPOINT")
	config.MongoDB = os.Getenv("MONGO_DB")
	config.MongoUser = os.Getenv("MONGO_USER")
	config.MongoPassword = os.Getenv("MONGO_PASSWORD")
	config.MongoEncryptionKey = os.Getenv("MONGO_ENCRYPTION_KEY")

	config.Production = getEnv("TRANCE_PRODUCTION", "false")
	config.Version = getEnv("TRANCE_VERSION", "astro-server")
	config.AstroEphimerisEndpoint = getEnv("ASTRO_EPHIMERIS_ENDPOINT", "")

	config.LogVerbosity, _ = getIntEnv("TRANCE_LOG_VERBOSITY", 1)

	Config = &config
	// logrus.Infof("Token is (%s)", config.AccessToken)
	return Config
}

func getEnv(key, defaultVal string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultVal
	}
	return value
}

func getIntEnv(key string, defaultVal int) (int, error) {

	intVal := os.Getenv(key)
	if intVal == "" {
		return defaultVal, fmt.Errorf("pack %v: ENV var is not set", key)
	} else {
		intervalInt, err := strconv.Atoi(intVal)
		if err != nil {
			return defaultVal, fmt.Errorf("pack %v: ENV var is not an integer", key)
		}
		return intervalInt, nil
	}
}
