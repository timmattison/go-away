package save_the_planet

import (
	"errors"
	"os"
)

var (
	ErrorEnvVarNotSet = errors.New("save-the-planet variable not set")
)

func GetOptionalEnvValue(name string, defaultValue string) string {
	value := os.Getenv(name)

	if value == "" {
		value = defaultValue
	}

	return value
}

func GetRequiredEnvValue(name string) (string, error) {
	value := os.Getenv(name)

	if value == "" {
		return "", ErrorEnvVarNotSet
	}

	return value, nil
}
