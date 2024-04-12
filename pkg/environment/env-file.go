package environment

import (
	"github.com/joho/godotenv"
	"github.com/timmattison/go-away/pkg/fylez"
	"github.com/timmattison/go-away/pkg/git-over-here"
	"os"
)

func LoadEnv() error {
	var envPath string
	var err error

	if envPath, err = git_over_here.GetRepoRelativePath(".env"); err != nil {
		return err
	}

	if !fylez.FileExists(envPath) {
		return os.ErrNotExist
	}

	return godotenv.Load(envPath)
}
