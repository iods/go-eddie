package cmd

import (
	"github.com/gen2brain/beeep"
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

		err := beeep.Notify("Title", "Message body", "assets/warning.png")
		if err != nil {
			panic(err)
		}


	},
}

func init() {
	rootCmd.AddCommand(askCmd)
	askCmd.Flags().BoolP(
		"daily", "d", false,
		"Ask daily for a reminder about tracking your routines.")
}
