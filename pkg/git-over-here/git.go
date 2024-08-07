package git_over_here

import (
	"fmt"
	"github.com/timmattison/go-away/pkg/fylez"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// GetRepoBase Returns the location of the .git directory for the current repo by searching up the directory tree until it finds it, gets an error, or reaches the maximum number of iterations
func GetRepoBase() (string, error) {
	// Get the current working directory the user is in
	var cwd string
	var err error

	if cwd, err = os.Getwd(); err != nil {
		return "", err
	}

	// Make sure we don't end up in a loop by checking if our last path is the same as the current path
	lastPath := ""

	// Make sure we don't end up in a loop by checking to see if we do too many iterations
	maximumIterations := 50
	iterationCount := 0

	basePath := cwd

	for {
		gitDirectory := path.Join(basePath, ".git")

		if _, err = os.Stat(gitDirectory); err == nil {
			// .git exists, we are done
			// NOTE: This means this may not work as expected when using submodules or nested repos
			return gitDirectory, nil
		}

		if !os.IsPermission(err) && !os.IsNotExist(err) {
			// If permission is denied or the file doesn't exist we can just ignore it but anything else is a legit error
			return "", err
		}

		// Go up one level
		basePath = filepath.Dir(basePath)

		iterationCount++

		if iterationCount >= maximumIterations {
			// Too many iterations
			break
		}

		if lastPath == basePath {
			// Ended up in the same place we came from
			break
		}

		// Keep track of where we just were
		lastPath = basePath
	}

	return "", os.ErrNotExist
}

// GetRepoRelativePath returns an absolute path for a file relative to the directory that contains the .git directory.
// If the git repo is in `/home/user/repo/.git` and you pass in "main.go" it will return `/home/user/repo/main.go`. This
// is determined by finding the .git directory, going up one level, joining the relative path to that directory, and
// then returning the absolute path.
func GetRepoRelativePath(relativePath string) (string, error) {
	var repoBase string
	var err error

	if repoBase, err = GetRepoBase(); err != nil {
		return "", err
	}

	return path.Join(repoBase, "..", relativePath), nil
}

func RunCommandInRepoDirectoriesWithFile(file string, command []string) {
	var repoBase string
	var err error

	commandString := strings.Join(command, " ")

	if repoBase, err = GetRepoBase(); err != nil {
		log.Fatal("Couldn't find the git repo", "error", err)
	}

	repo := path.Dir(repoBase) + "/"

	directoryHandler := fylez.DirectoryHandler(func(entryPath string, entry os.DirEntry) {
		filename := path.Join(entryPath, file)

		if !fylez.FileExists(filename) {
			return
		}

		if err = runCommandInDirectory(entryPath, command); err != nil {
			log.Fatal(fmt.Sprintf("Error running %s", commandString), "directory", entryPath, "error", err)
		}
	})

	if err = filepath.WalkDir(repo, fylez.VisitWithNameChecker(nil, nil, directoryHandler)); err != nil {
		log.Fatal("Error walking path", "path", repo, "error", err)
	}
}

func runCommandInDirectory(dir string, command []string) error {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Dir = dir
	return cmd.Run()
}
