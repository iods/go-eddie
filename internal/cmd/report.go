package cmd

import (
	"github.com/iods/go-eddie/internal/eddie"
	"github.com/spf13/cobra"
)

var demo, mood, seizure, sleep, weight eddie.Eddie

var reportCmd = &cobra.Command{
	Use: "report [activity]",
	Short: "Generate reports for an activity",
	Long: `
Eddie will get a report to you.

Among other things xD. Once Eddie generates all the data,
he creates some dashboards, then builds a webpage with some
cool looking data about what you have been up to.
`,
	Run: func(cmd *cobra.Command, args []string) {

		switch args[0] {
		case "mood":
			// mood.ReportMood()
		case "seizure":
			// seizure.ReportSeizure()
		case "sleep":
			// sleep.ReportSleep()
		case "weight":
			weight.ReportWeight()
		}
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}