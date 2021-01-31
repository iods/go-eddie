package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "eddie",
	Short: "A cli tool for life.",
	Long:
`
   ____ ___  ___  _  ____ 
  | |_ | | \| | \| || |_  
  |_|__|_|_/|_|_/|_||_|__

  A man's best friend.

  For, and by, the Dark Society.
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}