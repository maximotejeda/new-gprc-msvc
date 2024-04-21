package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return getEnvValue("ENV")
}

func GetDataSourceURL() string {
	return getEnvValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	portStr := getEnvValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}
	return port
}
func getEnvValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatal("env variable not set: " + key)
	}
	return os.Getenv(key)
}
