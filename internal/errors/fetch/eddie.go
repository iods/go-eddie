package fetch

/*
   Eddie: A CLI service dog.
   Copyright (c) 2023 Rye Miller
   https://ryemiller.io
*/

import (
	"fmt"

	"github.com/spf13/cobra"
)

// EddieRootCmd replaces the Execute() function for running the main cli.
func EddieRootCmd() *cobra.Command {
	eddieCmd := cobra.Command{
		Use:           "eddie",
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

			fmt.Println("Hello Eddie CLI")

		},
	}

	return &eddieCmd
}
