package setup

//
//import (
//	"fmt"
//	"github.com/iods/go-eddie/internal/errors"
//	"github.com/iods/go-eddie/internal/util/ioutils"
//	"os"
//	"path/filepath"
//	"runtime"
//)
//
//func setProjectDir() string {
//	var dir string
//	if runtime.GOOS == "windows" {
//		dir = filepath.Join(os.Getenv("APPDATA"), "eddie")
//	} else {
//		dir = filepath.Join(os.Getenv("HOME"), ".config/eddie")
//	}
//	return dir
//}
//
//func checkInstall() bool {
//	home := setProjectDir()
//	if s, err := os.Stat(home); err == nil && s.IsDir() {
//		return true
//	}
//	return false
//}
//
//func setConfigs(home string) {
//	file := filepath.Join(home, "/config")
//	_, err := os.Stat(file)
//	if os.IsNotExist(err) {
//		if _, err := os.Stat(home); os.IsNotExist(err) {
//			err = os.MkdirAll(home, os.ModePerm)
//		}
//		_, err := os.OpenFile(file, os.O_RDONLY|os.O_CREATE, 0644)
//		errors.Handle("unable to open this file", err)
//	}
//}
//
//func setRecords(home string, stub bool) {
//	var filename string
//
//	if stub == true {
//		filename = filepath.Join(home, "/records-stub.db")
//	} else {
//		filename = filepath.Join(home, "/records.db")
//	}
//	_, err := os.Stat(filename)
//	if os.IsNotExist(err) {
//		if _, err := os.Stat(home); os.IsNotExist(err) {
//			err = os.MkdirAll(home, os.ModePerm)
//		}
//		_, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
//		errors.Handle("unable to open the file", err)
//	}
//
//}
//
//func setFortunes(home string) {
//	file := filepath.Join(home, "/fortunes.json")
//	_, err := os.Stat(file)
//	if os.IsNotExist(err) {
//		if _, err := os.Stat(home); os.IsNotExist(err) {
//			err = os.MkdirAll(home, os.ModePerm)
//		}
//		_, err := os.OpenFile(file, os.O_RDONLY|os.O_CREATE, 0644)
//		errors.Handle("unable to open this file", err)
//	}
//	err = ioutils.CopyFile("../../scripts/fortunes.json", file)
//	errors.Handle("handle this error", err)
//}
//
//func Project() error {
//	if checkInstall() == false {
//		home := setProjectDir()
//		setConfigs(home)
//		stub := PromptInstall()
//		setRecords(home, stub)
//		if stub == true {
//			db.Stub()
//		}
//		setFortunes(home)
//	}
//	return nil
//}
//
//func ProjectDir() string {
//	home := setProjectDir()
//	return fmt.Sprintf("%s", home)
//}
