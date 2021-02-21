package project

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/iods/go-eddie/internal/env"
	"github.com/iods/go-eddie/internal/errors"
)

var eddieHome string

// setDir Checks existence of eddie project for creation of initial files and configurations.
func setDir() string {
	eddieHome = env.Get("EDDIE_HOME")
	if len(eddieHome) == 0 {
		if runtime.GOOS == "windows" {
			eddieHome = filepath.Join(os.Getenv("APPDATA"), "eddie")
		} else {
			eddieHome = filepath.Join(os.Getenv("HOME"), ".config/eddie")
		}
	}
	return eddieHome
}

// setupConfigFile Creates a configuration file with some defaults in the eddie project directory.
func setupConfigFile(eddieHome string) {
	filename := filepath.Join(eddieHome, "/config")
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		if _, err := os.Stat(eddieHome); os.IsNotExist(err) {
			err = os.MkdirAll(eddieHome, os.ModePerm)
		}
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		errors.Handle("unable to open file", err)
		f.WriteString("[default]\n")
		f.Close()
	}
}

// setupRecordsFile Creates a database file for use with storing our records.
func setupRecordsFile(eddieHome string) {
	filename := filepath.Join(eddieHome, "/records.db")
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		if _, err := os.Stat(eddieHome); os.IsNotExist(err) {
			err = os.MkdirAll(eddieHome, os.ModePerm)
		}
		_, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
		errors.Handle("unable to open file", err)
	}
}