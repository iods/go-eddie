package common

//
//import (
//	"fmt"
//	"github.com/mitchellh/go-homedir"
//	"os"
//	"path/filepath"
//	"runtime"
//
//	model "github.com/iods/go-standups/internal/datasource/fs/model/Project"
//)
//
//// loadBaseProjectPath returns a OS specific path for setting the home directory (base path) for the application.
//func loadBaseProjectPath() model.Project {
//	p := model.Project{}
//	userOS := runtime.GOOS
//
//	switch userOS {
//	case "windows":
//		fmt.Println("Windows")
//		// p.Directory = filepath.Join(os.Getenv("APPDATA"), "standups")
//	case "darwin":
//		fmt.Println("OSX")
//	case "linux":
//		fmt.Println("Linux")
//	default:
//		fmt.Printf("%s", userOS)
//	}
//	return p
//}
//
//// Success returns the success message and prints it in stdout.
//func Success(message string, args ...interface{}) {
//	_, err := fmt.Fprintf(os.Stdout, fmt.Sprintf("\n\u001B[0;32m✓\u001B[0m %s\n", message), args...)
//	if err != nil {
//		return
//	}
//}
//
//// Warn returns the warning message in stderr.
//func Warn(message string, args ...interface{}) {
//	_, err := fmt.Fprintf(os.Stderr, fmt.Sprintf("\u001B[0;33m%s\u001B[0m\n", message), args...)
//	if err != nil {
//		return
//	}
//}
//
//// Failure prints failure messages in stderr.
//func Failure(message string, args ...interface{}) {
//	_, err := fmt.Fprintf(os.Stderr, fmt.Sprintf("\u001B[0;31m✗\u001B[0m %s\n", message), args...)
//	if err != nil {
//		return
//	}
//}
//
//// Failed prints the failure message and then exits.
//func Failed(message string, args ...interface{}) {
//	Failure(message, args...)
//	os.Exit(1)
//}
//
//// ConfigDir returns the home directory with the .config dir appended to it.
//func ConfigDir() (string, error) {
//	home := os.Getenv("XDG_CONFIG_HOME")
//	if home != "" {
//		return home, nil
//	}
//
//	home, err := homedir.Dir()
//	if err != nil {
//		return "", err
//	}
//
//	dir := filepath.Join(home, ".config")
//
//	return dir, nil
//}
