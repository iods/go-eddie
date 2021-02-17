package config

import (
	"fmt"
	"github.com/iods/go-eddie/internal/env"
	"os"
	"path/filepath"
	"runtime"

	"github.com/iods/go-eddie/internal/utils"
)

var dir string

func GetConfigDir() (string, error) {

	dir = env.EnvVar("CONFIG_PATH")

	if len(dir) == 0 {
		if runtime.GOOS == "windows" {
			dir = filepath.Join(os.Getenv("APPDATA"), "eddie")
		} else {
			dir = filepath.Join(os.Getenv("HOME"), ".config/eddie")
		}
	}

	if _, err := os.Stat(filepath.Join(dir, "eddie")); err == nil {
		fmt.Println("nope")
	} else {
		fmt.Println("yep")
	}

	return dir, nil
}

func CreateRecordsFile(dir string) {
	filename := filepath.Join(dir, "/records.json")
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		utils.HandleError("Error opening file", err)
		f.WriteString("{\n")
		f.WriteString("\t\"records\":\n")
		f.WriteString("}")
		f.Close()
	}
}