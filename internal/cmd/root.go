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
            __     __  __        
 .-----..--|  |.--|  ||__|.-----.
 |  -__||  _  ||  _  ||  ||  -__|
 |_____||_____||_____||__||_____|

 A man's best friend.
`,
}

func Execute() {
	// @TODO confirm this is necessary every execution
	if project.Check() == false {
		project.Install()
		db.Init()
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}