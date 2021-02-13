package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var sleepCmd = &cobra.Command{
	Use:   "sleep",
	Short: "Sleep is important. Track it.",
	Long: `
Eddie tracks your sleep data.

Using some pretty nifty args and flags, eddie will take
this data and store it for reporting on it later. The main
things eddie records are quality and duration.
`,
	Run: func(cmd *cobra.Command, args []string) {

		a := args

		q, _ := cmd.Flags().GetInt("quality")
		d, _ := cmd.Flags().GetInt("duration")
		t, _ := cmd.Flags().GetStringArray("tags")
		s, _ := cmd.Flags().GetString("start")
		e, _ := cmd.Flags().GetString("end")

		fmt.Printf("Arguments: %v\n", a)
		fmt.Printf("Quality: %d\n", q)
		fmt.Printf("Duration: %d\n", d )
		fmt.Printf("Tags: %v\n", t)
		fmt.Printf("Start: %s\n", s)
		fmt.Printf("End: %s\n", e)
	},
}

// init Main function for Cobra to build the buddha command.
func init() {
	var test []string
	trackCmd.AddCommand(sleepCmd)
	sleepCmd.Flags().IntP("quality", "q", 0, "Quality of your sleep from one being lowest to 10 being highest.")
	sleepCmd.Flags().IntP("duration", "d", 8, "Duration of sleep in hours (default is 8).")
	sleepCmd.Flags().StringArrayP("tags", "t", test, "Tags eddie can link to a sleep record.")
	sleepCmd.Flags().StringP("start", "s", "23:59", "Time eddie starts your sleep record.")
	sleepCmd.Flags().StringP("end", "e", "07:00", "Time eddie stops your sleep record.")
}
