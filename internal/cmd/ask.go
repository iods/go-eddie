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



// create a file in the home folder with each behavior
// add a frequency (already in struct) to that file
// every time eddie is run it checks a var that confirms the behavior was done or not
// basic writing to file, reading from file when specific commands are run

// get the monthly fixed
// get min and max
// amount gained, lost, conditions, maybe look up some interesting ones

// write some readme

// clean uo some code

// eddie roll over to clear everything and reset the database

// function to get sum of type, you entered your sleep (X) times in the past month, good job, bad job