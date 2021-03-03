package cmd

import (
	"fmt"
	"os"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/util/project"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "eddie",
	Short: "A cli tool for life.",
	Long:
`

  _______ ______  ______  _____ _______
  |______ |     \ |     \   |   |______
  |______ |_____/ |_____/ __|__ |______
                                      
     the cli service dog 🐕 v0.1.0


`,
}

func Execute() {
	// @TODO confirm this is necessary every execution
	if project.Check() == false {
		project.Install()
		db.StubDatabase()
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}