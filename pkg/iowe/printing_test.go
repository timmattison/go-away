package iowe

import (
	"github.com/timmattison/go-away/pkg/save-the-planet"
	"os"
	"testing"

	"golang.org/x/text/language"
)

func TestGetPrinter(t *testing.T) {
	englishBase, _ := language.English.Base()
	frenchBase, _ := language.French.Base()
	germanBase, _ := language.German.Base()

	tests := []struct {
		name string
		lang string
		want language.Base
	}{
		{"English", "en_US.UTF-8", englishBase},
		{"French", "fr_FR.UTF-8", frenchBase},
		{"German", "de_DE.UTF-8", germanBase},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(save_the_planet.LangEnvironmentVariableName, tt.lang)
			defer os.Unsetenv(save_the_planet.LangEnvironmentVariableName)

			_, tag := GetPrinter()
			base, _, _ := tag.Raw()

			if base != tt.want {
				t.Errorf("GetPrinter() = %v, want %v", base, tt.want)
			}
		})
	}
}
