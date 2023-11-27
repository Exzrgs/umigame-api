package utils

import (
	"os"
	"path/filepath"
)

func GetPath()(string, error){
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	path := filepath.Dir(exe)
	return path, nil
}