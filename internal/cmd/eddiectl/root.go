package eddiectl

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
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
)

func ExecuteA() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("ooooooops")
		os.Exit(1)
	}
}
