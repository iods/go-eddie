package file

import "os"

// IfNotExists
func IfNotExists(path string) (err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// path does not exist
	}
	return err
}

// IfExists
func IfExists(path string) (err error) {
	if _, err := os.Stat(path); err == nil {
		// path does exist
	}
	return err
}

