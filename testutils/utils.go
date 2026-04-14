package testutils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// GetRootFolder determines the root folder of the project
func GetRootFolder() (string, error) {
	wd, _ := os.Getwd()
	return lookupRootFolder(wd, 0)
}

func lookupRootFolder(dir string, level int) (string, error) {
	if level > 5 {
		return "", errors.New("failed to find root folder")
	}
	// Look for go.mod first (for library projects), then main.go (for application projects)
	goModFile := fmt.Sprintf("%s/go.mod", dir)
	if fileExists(goModFile) {
		return dir, nil
	}
	mainFile := fmt.Sprintf("%s/main.go", dir)
	if fileExists(mainFile) {
		return dir, nil
	}
	nextLevel := level + 1
	parentDir := filepath.Dir(dir)
	return lookupRootFolder(parentDir, nextLevel)
}

func fileExists(file string) bool {
	if stat, err := os.Stat(file); err == nil {
		return !stat.IsDir()
	}
	return false
}
