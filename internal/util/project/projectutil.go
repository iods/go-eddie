package project

import "os"

// Install Builds the app for use.
func Install() {
	p := setDir()
	setupConfigFile(p)
	setupRecordsFile(p)
}

// Check Returns a bool for triggering the install of the app.
func Check() bool {
	p := setDir()
	if s, err := os.Stat(p); err == nil && s.IsDir() {
		return true // if the directory exists
	}
	return false
}

// Getdir Returns the project dir for access separate of `setDir()`
func Getdir() string {
	return setDir()
}