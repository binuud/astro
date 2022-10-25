package config

// binu@dronasys.com

import (
	"fmt"
	"os"
	"strconv"
)

type EnvConfig struct {
	MongoEndpoint                        string
	MongoDB                              string
	MongoUser                            string
	MongoPassword                        string
	MongoEncryptionKey                   string
	Production                           string
	EventServerEndPoint                  string
	LogVerbosity                         int
	SpecTemplatePath                     string
	TranceWebURL                         string
	TRANCE_EVENTS_SERVER_PUBLIC_ENDPOINT string
	Cluster_Service_Token                string
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
	config.Production = os.Getenv("TRANCE_PRODUCTION")
	config.EventServerEndPoint = os.Getenv("TRANCE_SERVER_ENDPOINT")
	config.LogVerbosity, _ = getIntEnv("TRANCE_LOG_VERBOSITY", 1)

	config.SpecTemplatePath = os.Getenv("SPEC_TEMPLATE_PATH")
	config.TranceWebURL = os.Getenv("TRANCE_WEB_URL")
	config.TRANCE_EVENTS_SERVER_PUBLIC_ENDPOINT = os.Getenv("TRANCE_EVENTS_SERVER_PUBLIC_ENDPOINT")

	config.Cluster_Service_Token = os.Getenv("CLUSTER_SERVICE_TOKEN")
	Config = &config

	return Config
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
