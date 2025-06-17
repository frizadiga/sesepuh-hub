package main

import "os"

// GetModelToUse returns the model name to use based on environment variables
func GetModelToUse(envVarName, defaultModel string) string {
	modelName := defaultModel

	if envModel := os.Getenv("SESEPUH_HUB_MODEL"); envModel != "" {
		modelName = envModel
		return modelName
	}

	if envModel := os.Getenv(envVarName); envModel != "" {
		modelName = envModel
	}

	return modelName
}

// GetEnv returns environment variable value or default if not set
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

// WriteToFile writes the given content to a file
func WriteToFile(filename string, content []byte) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = file.Write(content); err != nil {
		return err
	}

	return nil
}

// WriteRespToFile writes the response content to a file
func WriteRespToFile(content []byte, filename string) error {
	if filename == "" {
		filename = "./.last-response.txt"
	}

	if err := WriteToFile(filename, []byte(content)); err != nil {
		return err
	}

	return nil
}
