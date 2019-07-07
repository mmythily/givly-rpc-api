package config

import (
	"os"
	"path/filepath"
	"strings"
)

// GetExecPath gets the Path of the executable file
func GetExecPath() (string, error) {
	dir, err := os.Executable()
	if err != nil {
		return "", err
	}
	execPath, err := filepath.EvalSymlinks(dir)
	if err != nil {
		return "", err
	}
	return execPath, nil
}

// GetPath gets the full path of the folder/file
func GetPath(foldername string) (string, error) {
	execPath, err := GetExecPath()
	if err != nil {
		return "", err
	}
	PathList := []string{filepath.Dir(execPath), foldername}
	Path := strings.Join(PathList, "")
	return Path, nil
}
