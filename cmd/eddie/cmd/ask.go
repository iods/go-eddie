package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use: "ask [reminder]",
	Short: "Reminders for tracking activity",
	Long: `
Eddie will set you a reminder for compiling this information.

Among other things xD.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Set an example from eddie.")
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
