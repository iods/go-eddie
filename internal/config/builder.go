package config

//import (
//	"fmt"
//	"github.com/AlecAivazis/survey/v2"
//	"github.com/spf13/viper"
//	"os"
//	"strings"
//
//	c "github.com/iods/go-standups/internal/datasource/fs/model/Config"
//	p "github.com/iods/go-standups/internal/datasource/fs/model/Project"
//)
//
//// Configuration options for standups
//const (
//	Dir      = "Developer" // Dir is a projctl config directory.
//	Filename = ".projects" // Filename is a projctl config file.
//	Filetype = "yaml"      // Filetype is a coctl config file extension.
//)
//
//var (
//	// ErrSkip returns a message when the user skips setup.
//	ErrSkip = fmt.Errorf("skipping the setup configuration")
//
//	// ErrUnexpectedResponseFormat returns a message when the data is in an unexpected format.
//	ErrUnexpectedResponseFormat = fmt.Errorf("unexpected response format")
//)
//
//type Builder struct {
//	conf  *c.Config
//	value struct {
//		directory string
//		project   *p.Project
//		force     bool
//	}
//	projectsMap map[string]*p.Project
//}
//
//// StandupsConfigBuilder will generate a new configuration for the application.
//func StandupsConfigBuilder(config *c.Config) *Builder {
//	b := Builder{
//		conf:        config,
//		projectsMap: make(map[string]*p.Project),
//	}
//	return &b
//}
//
//func (c *Builder) Build() (string, error) {
//	ceb := func() bool {
//		fmt.Println("checking configuration.")
//		return FileExists(viper.ConfigFileUsed())
//	}()
//
//}
//
//
//func (b *Builder) configureDirectoryDetails() error {
//	var qs []*survey.Question
//
//	b.value.directory = b.conf.Directory
//	b.value.project = b.projectsMap[strings.ToUpper()]
//
//
//}
//
//
//
//
//// FileExists returns true if a file exists.
//func FileExists(file string) bool {
//	if file == "" {
//		return false
//	}
//
//	if _, err := os.Stat(file); os.IsNotExist(err) {
//		return false
//	}
//
//	return true
//}
//
//func create(path, name string) error {
//	const perm = 0o700
//
//}
