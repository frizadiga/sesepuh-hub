package main

import "os"

// GetModelToUse returns the model name to use based on environment variables
func GetModelToUse(envVarName, defaultModel string) string {
	modelName := defaultModel

	if envModel := os.Getenv(envVarName); envModel != "" {
		modelName = envModel
	}

	return modelName
}

// GetEnvWithDefault returns environment variable value or default if not set
func GetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}