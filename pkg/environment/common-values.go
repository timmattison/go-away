package environment

import (
	"os"
	"strings"
)

const (
	LangEnvironmentVariableName = "LANG" // LangEnvironmentVariableName is the name of the environment variable that contains the locale
)

// GetLang returns the value of the LANG environment variable. If the LANG environment
// variable is not set, an empty string is returned.
func GetLang() string {
	// This should return the "LANG" in the form of "en_US.UTF-8"
	return os.Getenv(LangEnvironmentVariableName)
}

// GetLocale returns the locale part of the LANG environment variable. If the LANG
// environment variable is not set, an empty string is returned.
func GetLocale() string {
	langEnv := GetLang()

	if langEnv == "" {
		return ""
	}

	parts := strings.Split(langEnv, ".")

	return parts[0]
}

// GetEncoding returns the encoding part of the LANG environment variable. If the LANG
// environment variable is not set or does not contain an encoding, an empty string is returned.
func GetEncoding() string {
	langEnv := GetLang()

	if langEnv == "" {
		return ""
	}

	parts := strings.Split(langEnv, ".")

	if len(parts) > 1 {
		return parts[1]
	}

	return ""
}
