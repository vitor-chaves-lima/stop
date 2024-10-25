package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Environment Environment
	Database    MongoDBConfig
}

func LoadConfig() *Config {
	var errors []error

	mongoDbConfig, databaseEnvironmentErrors := loadDatabaseEnvironmentValues()
	if databaseEnvironmentErrors != nil {
		errors = append(errors, databaseEnvironmentErrors...)
	}

	if errors != nil {
		var errorMessage string
		for _, err := range errors {
			errorMessage += fmt.Sprintf("- %s\n", err.Error())
		}
		panic(fmt.Sprintf("Could not load config:\n%s", errorMessage))
	}

	environment, err := NewEnvironment(getEnv("APP_ENV", "dev"))
	if err != nil {
		errors = append(errors, err)
	}

	return &Config{
		Environment: environment,
		Database:    mongoDbConfig,
	}
}

func loadDatabaseEnvironmentValues() (mongodbConfig MongoDBConfig, errors []error) {
	databaseUsername, err := getEnvRequired("DB_USERNAME")
	if err != nil {
		errors = append(errors, err)
	}

	databasePassword, err := getEnvRequired("DB_PASSWORD")
	if err != nil {
		errors = append(errors, err)
	}

	databaseHost, err := getEnvRequired("DB_HOST")
	if err != nil {
		errors = append(errors, err)
	}

	databasePortString, err := getEnvRequired("DB_PORT")
	if err != nil {
		errors = append(errors, err)
	}

	var databasePort int
	if databasePassword != "" {
		databasePort, err = strconv.Atoi(databasePortString)
		if err != nil {
			errors = append(errors, err)
		}
	}

	databaseName, err := getEnvRequired("DB_NAME")
	if err != nil {
		errors = append(errors, err)
	}

	if errors != nil {
		return MongoDBConfig{}, errors
	}

	return MongoDBConfig{
		Username: databaseUsername,
		Password: databasePassword,
		Host:     databaseHost,
		Port:     databasePort,
		Database: databaseName,
	}, nil
}

func getEnvRequired(key string) (envVar string, err error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}

	return "", fmt.Errorf("the environment variable %s is not defined", key)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
