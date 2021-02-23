package cmd

import (
	"fmt"
	"github.com/iods/go-eddie/internal/cli/template"
	"github.com/iods/go-eddie/internal/db/models"
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

		switch args[0] {
		case "demo":
			template.RenderDemo()
		case "mood":
			template.RenderMood()
		case "seizure":
			template.RenderSeizure()
		case "sleep":
			template.RenderSleep()
		case "weight":
			// weight.GetAverage()
			//template.RenderWeight()

			reportWeight()


		}
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}

func reportWeight() {
	var weightModel models.WeightModel

	r, _ := weightModel.GetRecords()

	for _, result := range r {
		fmt.Println(result.ToString())
	}
}