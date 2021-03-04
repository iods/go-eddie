package cmd

import (
	"fmt"
	"os"

	util "github.com/iods/go-eddie/internal/util/project"
	"github.com/spf13/cobra"
)

var (

	rootCmd = &cobra.Command{
		Use: "eddie",
		Short: "the cli service dog.",
		Long:
`

  _______ ______  ______  _____ _______
  |______ |     \ |     \   |   |______
  |______ |_____/ |_____/ __|__ |______
                                      
     the cli service dog 🐕 v0.1.0


`,
	}
	setup = &util.Project{}
)

func Execute() {

	if setup.CheckDirectory() == false {
		setup.InstallDirectory()
		return
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}