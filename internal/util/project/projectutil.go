package util

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/iods/go-eddie/internal/errors"
)

type Project struct {
	Directory string
}

// loadProject Returns an OS specific path for setting the home directory (base path) for the application.
func loadProject() Project {
	p := Project{}
	if runtime.GOOS == "windows" {
		p.Directory = filepath.Join(os.Getenv("APPDATA"), "eddie")
	} else {
		p.Directory = filepath.Join(os.Getenv("HOME"), ".config/eddie")
	}
	return p
}

// loadConfig Creates and writes some basic defaults to a config file for getting basic data like BMI, starting weight, etc.
func loadConfig(dir string)  {
	filename := filepath.Join(dir, "/config")
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
		}
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		errors.Handle("handle this error at some point", err)
		f.WriteString("eddie, the cli service dog\n\n")
		f.Close()
	}
}

// loadRecords Creates a records file for the application to write tracked records to.
func loadRecords(dir string, stub bool) {
	var filename string

	if stub == true {
		filename = filepath.Join(dir, "/records-stub.db")
	} else {
		filename = filepath.Join(dir, "/records.db")
	}
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
		}
		_, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
		errors.Handle("unable to open file", err)
	}
}

// GetProjectHome Returns the installation path for the application.
func GetProjectHome() string {
	p := loadProject()
	home := fmt.Sprintf("%s", p.Directory)
	return home
}

// CheckDirectory Returns a boolean for methods to confirm the existence of a project path.
func (p Project) CheckDirectory() bool {
	proj := loadProject()
	if s, err := os.Stat(proj.Directory); err == nil && s.IsDir() {
		return true
	}
	return false
}

// InstallDirectory Installs the application and returns a boolean for confirming the success or failure of the install.
func (p Project) InstallDirectory() bool {
	proj := loadProject()
	loadConfig(proj.Directory)
	loadRecords(proj.Directory, false)
	return true
}