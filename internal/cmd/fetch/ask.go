package fetch

//
//import (
//	"github.com/iods/go-eddie/internal/api"
//	"github.com/spf13/cobra"
//)
//
//var askCmd = &cobra.Command{
//	Use:   "ask [reminder]",
//	Short: "Reminders for tracking activity",
//	Long: `
//Eddie will set you a reminder for compiling this information.
//
//Among other things xD.
//`,
//	Run: func(cmd *cobra.Command, args []string) {
//		switch args[0] {
//		case "buddha":
//			api.AskBuddha()
//		}
//	},
//}
//
//func init() {
//	rootCmd.AddCommand(askCmd)
//	askCmd.Flags().BoolP(
//		"daily", "d", false,
//		"Ask daily for a reminder about tracking your routines.")
//}
//
//// create a file in the home folder with each behavior
//// add a frequency (already in struct) to that file
//// every time eddie is run it checks a var that confirms the behavior was done or not
//// basic writing to file, reading from file when specific commands are run
