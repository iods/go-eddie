package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use: "report [behavior]",
	Short: "Generate reports for an activity",
	Long: `
Eddie will get a report to you.

Among other things xD. Once Eddie generates all the data,
he creates some dashboards, then builds a webpage with some
cool looking data about what you have been up to.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Eddie will show you what your " + strings.Join(args, "") + " looks like lately.")
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
