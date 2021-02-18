package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/iods/go-eddie/internal/env"
	"github.com/iods/go-eddie/internal/helpers/error"
)

var eddieHome string

// setupHomeDir
func setupHomeDir() string {

	eddieHome = env.Get("EDDIE_HOME")
	if len(eddieHome) == 0 {
		if runtime.GOOS == "windows" {
			eddieHome = filepath.Join(os.Getenv("APPDATA"), "eddie")
		} else {
			eddieHome = filepath.Join(os.Getenv("HOME"), ".config/eddie")
		}
	}
	if _, err := os.Stat(filepath.Join(eddieHome, "config")); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("there is no eddie project yet.")
		} else {
			fmt.Println("there is one, or err out.")
		}
	}
	return eddieHome
}

// setupRecordFile
func setupRecordFile(eddieHome string) {
	filename := filepath.Join(eddieHome, "/config")
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		if _, err := os.Stat(eddieHome); os.IsNotExist(err) {
			err = os.MkdirAll(eddieHome, os.ModePerm)
		}
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		error.Handle("unable to open file", err)
		f.WriteString("[default]")
		f.Close()
	}
}

// setupLogFile
func setupLogFile(eddieHome string) {
	filename := filepath.Join(eddieHome, "/log")
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		if _, err := os.Stat(eddieHome); os.IsNotExist(err) {
			err = os.MkdirAll(eddieHome, os.ModePerm)
		}
		_, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
		error.Handle("unable to open file", err)
	}
}

// setupConfigFile
func setupConfigFile(eddieHome string) {
	filename := filepath.Join(eddieHome, "/records.db")
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		if _, err := os.Stat(eddieHome); os.IsNotExist(err) {
			err = os.MkdirAll(eddieHome, os.ModePerm)
		}
		_, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
		error.Handle("unable to open file", err)
	}
}

func InstallEddie() {
	dir := setupHomeDir()
	setupConfigFile(dir)
	setupLogFile(dir)
	setupRecordFile(dir)
}