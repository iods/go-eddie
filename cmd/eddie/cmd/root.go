package cmd

import (
	"fmt"
	"log"

	"github.com/iods/go-eddie/internal/config"
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

	c, err := config.GetConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)

	config.CreateRecordsFile(c)

	//if err := rootCmd.Execute(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
}