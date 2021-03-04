package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/iods/go-eddie/internal/errors"
	"github.com/manifoldco/promptui"
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
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
		}
		_, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
		errors.Handle("unable to open file", err)
	}
}

func loadPrompt() bool {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		return err
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{  . }}\n\n",
		Valid:   "{{  . | bold | green }} ",
		Invalid: "{{  . | bold}} ",
		Success: "{{  . | bold | green }} ",

	}

	fmt.Printf(`
 so you've decided to take home
   _______ ______  ______  _____ _______
   |______ |     \ |     \   |   |______
   |______ |_____/ |_____/ __|__ |______

 there are a few things eddie needs to know
 before you can get started:

`)
	sp := promptui.Prompt{
		Label: " [1]  Would you like to stub the database? ",
		Templates: templates,
		IsConfirm: true,
	}
	one, err := sp.Run()

	wp := promptui.Prompt{
		Label: " [2]  What is your starting weight (lbs) ? ",
		Templates: templates,
		Validate: validate,
	}
	two, err := wp.Run()

	hp := promptui.Prompt{
		Label: " [3]  What is your starting height (in) ? ",
		Templates: templates,
		Validate: validate,
	}
	three, err := hp.Run()

	if err != nil {
		fmt.Printf("prompt failed %v\n", err)
	}

	writeConfig("weight", two)
	writeConfig("height", three)

	if one == "y" {
		fmt.Printf("\nThat's it. Now run `eddie --help` to get started and view a demo setup.\n")
		return true
	}

	fmt.Printf("\nThat's it. Now run `eddie --help` to get started.\n")
	return false
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

// writeConfig Writes the prompt answers to a configuration file for setting application defaults.
func writeConfig(config string, record string) {
	path := loadProject()
	filename := filepath.Join(path.Directory, "/config")
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	errors.Handle("handle this some time.", err)
	if _, err := f.Write([]byte(config + ":" + record + "\n")); err != nil {
		f.Close()
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
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
	stub := loadPrompt()
	loadRecords(proj.Directory, stub)
	if stub == false {
		return false
	}
	return true
}