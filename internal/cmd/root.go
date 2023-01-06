package cmd

import (
	"fmt"
	"github.com/iods/go-eddie/internal/setup"
	"os"

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
)

func Execute() {





	if err := setup.Project(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	//if util.CheckDirectory() == false {
	//	if util.Install() != false {
	//		db.Stub()
	//	}
	//	return
	//}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}