package utils

import (
	"os"
)

// ProjectInstall Quick installs the eddie CLI for SoC and modularity.
func ProjectInstall() {
	dir := setProjectDir()
	setupConfigFile(dir)
	setupRecordsFile(dir)
}

// ProjectCheck Returns bool for triggering the eddie installation and configuration.
func ProjectCheck() bool {
	dir := setProjectDir()
	if s, err := os.Stat(dir); err == nil && s.IsDir() {
		return true // if the directory exists
	}
	return false
}

// GetProjectDir Returns a public facing, always available string of eddie's installation path.
func GetProjectDir() string {
	return setProjectDir()
}