package root

/*
   Eddie: A CLI service dog.
   Copyright (c) 2022 Rye Miller
   https://ryemiller.io
*/

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

const (
	gcloudAPITokenLink = ""
	slackAPITokenLink  = ""
)

// EddiectlRootCmd replaces the Execute() function for running the main cli.
func EddiectlRootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:           "eddiectl",
		Short:         "A CLI tool and service dog.",
		SilenceErrors: true,
		SilenceUsage:  true,
		Long: `
 _______ ______  ______  _____ _______
 |______ |     \ |     \   |   |______
 |______ |_____/ |_____/ __|__ |______

 the cli tool and service dog 🐕 v0.1.0 🖕🤬

`,
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("Hello Eddie CLI")

			home, err := homedir.Dir()
			if err != nil {
				fmt.Println(home)
			}
			fmt.Println(home)
		},
	}

	return &cmd
}
