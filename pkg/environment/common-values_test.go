package environment

import (
	"os"
	"testing"
)

func TestGetLang(t *testing.T) {
	expected := "en_US.UTF-8"
	os.Setenv(LangEnvironmentVariableName, expected)
	defer os.Unsetenv(LangEnvironmentVariableName)

	if got := GetLang(); got != expected {
		t.Errorf("GetLang() = %v, want %v", got, expected)
	}
}

func TestGetLocale(t *testing.T) {
	tests := []struct {
		lang     string
		expected string
	}{
		{"en_US.UTF-8", "en_US"},
		{"fr_FR.UTF-8", "fr_FR"},
		{"", ""},
	}

	for _, tt := range tests {
		os.Setenv(LangEnvironmentVariableName, tt.lang)
		if got := GetLocale(); got != tt.expected {
			t.Errorf("GetLocale() with LANG=%v, got %v, want %v", tt.lang, got, tt.expected)
		}
		os.Unsetenv(LangEnvironmentVariableName)
	}
}

func TestGetEncoding(t *testing.T) {
	tests := []struct {
		lang     string
		expected string
	}{
		{"en_US.UTF-8", "UTF-8"},
		{"fr_FR.ISO-8859-1", "ISO-8859-1"},
		{"en_US", ""},
		{"", ""},
	}

	for _, tt := range tests {
		os.Setenv(LangEnvironmentVariableName, tt.lang)
		if got := GetEncoding(); got != tt.expected {
			t.Errorf("GetEncoding() with LANG=%v, got %v, want %v", tt.lang, got, tt.expected)
		}
		os.Unsetenv(LangEnvironmentVariableName)
	}
}
