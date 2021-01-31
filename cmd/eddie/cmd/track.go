package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var trackCmd = &cobra.Command{
	Use: "track [behavior]",
	Short: "Track a behavior or routine",
	Long: `
From sleep to mood to symptoms, Eddie is like velcro.

Eddie will track your life. so here are some ways you
can use him to the fullest:
  1. eddie track [behavior] returns a Q&A for those long data-sets
  2. eddie track [behavior] -x sets it as important
  3. Any ideas?
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Track something from eddie. Like: ", strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)
}
