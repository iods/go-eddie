package cmd

import (
	"fmt"
	"github.com/iods/go-eddie/internal/db/models"
	"github.com/iods/go-eddie/internal/errors"
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
			var model models.RecordModel

			a, err := model.GetRecords()
			errors.Handle("something went wrong retrieving all records.", err)

			for _, r1 := range a {
				fmt.Printf("id: %d\ncreated at: %v\n", r1.ID, r1.CreatedAt)
			}

			b, err := model.GetRecordsByDate("2020-04-26")
			errors.Handle("something bad happened", err)

			for _, r2 := range b {
				fmt.Println(r2.ID)
			}

			c, err := model.GetRecordsByDateRange("2020-04-01", "2020-04-30")
			errors.Handle("something bad happened when getting the range records.", err)

			for _, r3 := range c {
				fmt.Printf("id: %d\ncreated at: %v\n", r3.ID, r3.CreatedAt)
			}
		case "mood":




		case "seizure":






		case "sleep":











		case "weight":
			var model models.WeightModel

			a, err := model.GetWeightRecords()
			errors.Handle("something happened while getting all of the weight records", err)

			for _, r1 := range a {
				fmt.Printf("id: %d\ncreated at: %v\n", r1.ID, r1.CreatedAt)
			}

			b, err := model.GetWeightRecordsByDate("2020-04-26")
			errors.Handle("something happened while getting the records by date", err)

			for _, r2 := range b {
				fmt.Printf("id: %d\ncreated at: %v\n", r2.ID, r2.CreatedAt)
			}

			c, err := model.GetWeightRecordsByDateRange("2020-12-22", "2021-02-22")
			errors.Handle("something bad happened when getting the range records.", err)

			for _, r3 := range c {
				fmt.Printf("id: %d\ncreated at: %v\n", r3.ID, r3.CreatedAt)
			}
			// get avg of 3, 6, 9 months


			// get monthly average

			// get highest weight, and when it was
			// get lowest weight, and when it was

			// get and list all the important ones

			// pie chart avg for season (weighed most in winter)

		}
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}