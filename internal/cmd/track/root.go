package track

/*
Standups: A Go application for trying harder.
Copyright (c) 2022 Rye Miller
https://ryemiller.io
*/

import (
	"fmt"
	"github.com/iods/go-eddie/internal/util/common"

	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// set this by default unless updated through viper
const (
	bitbucketAPITokenLink = ""
	jiraAPITokenLink      = ""
	slackAPITokenLink     = ""
)

var (
	conf    string
	verbose bool
)

// rootCmd holds the `projctl` command for running the application.
var (
	rootCmd = &cobra.Command{
		Use:           "projctl",
		Short:         "A Go application for TL's wanting to try harder.",
		SilenceErrors: true, // do i need this?
		SilenceUsage:  true, // do i need this?
		Long: `
	 _____ _____ _____ _____ _____ _____ _____ _____
	||s  |||t  |||a  |||n  |||d  |||u  |||p  |||s  ||
	||___|||___|||___|||___|||___|||___|||___|||___||
	|/___\|/___\|/___\|/___\|/___\|/___\|/___\|/___\|

	       the cli tool for TLs 🖕🤬 v0.1.0

`,
		// Run: func(cmd *cobra.Command, args []string) {fmt.Println("Hello Standups CLI")},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// run some configuration files
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&conf, "config", "", "Config file (default is $HOME/.projects.yaml)")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose error output")
}

func initConfig() {
	if conf != "" {
		viper.SetConfigFile(conf)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			common.Failed("Error: %s", err)
			return
		}

		viper.AddConfigPath(home)
		viper.SetConfigName("projects")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	fmt.Println("Using configuration file:", viper.ConfigFileUsed())

	viper.SetDefault("projects.client_a", "DefaultProject")

	test := viper.GetString("projects.client_a.name")
	fmt.Println("name", test)
}
