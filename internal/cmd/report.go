package cmd

import (
	"fmt"
	"github.com/iods/go-eddie/internal/cli/template"
	"github.com/iods/go-eddie/internal/errors"

	"github.com/iods/go-eddie/internal/cli"
	"github.com/iods/go-eddie/internal/db/schema"
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

			one, err := cli.GetWeightAverage(1, false)
			errors.Handle("handle this", err)
			fmt.Println(one)
			// return all records
			// return all records by a date
			// return all records by a date range
			// return all records by a date range custom
			// return all records
		case "mood":
			var yearData []float64
			var yearLabel []string

			for i := 1; i < 12; i++ {
				label, avg, _ := cli.GetWeightAverageByMonth(i)
				yearData = append(yearData, avg)
				yearLabel = append(yearLabel, label)
				fmt.Println(label, avg)
			}

		case "seizure":

		case "sleep":
			template.ReportDashboard()
		case "weight":

			// Average (1 month)
			//FindAverageMonthly(-9)
		}
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}

func getAverage(n []float64) float64 {
	total := 0.0

	for _, v := range n {
		total += v
	}

	return total / float64(len(n))
}

func getMinMax(r []schema.Record) (min schema.Record, max schema.Record) {
	min = r[0]
	max = r[0]

	for _, result := range r {
		if result.Total > max.Total {
			max = result
		}
		if result.Total < min.Total {
			min = result
		}
	}

	return min, max
}
//
//func FindAverageMonthly(monthRange int) {
//	var model models.WeightModel
//
//	now := time.Now()
//	date := now.String()
//	date = now.Format("2006-01-02")
//
//	records, err := model.GetWeightRecordsByRangeCustom(date, monthRange)
//	errors.Handle("handle this error message at some point.", err)
//
//	fmt.Println(records)
//
//	var total = 0.0
//
//	for _, record := range records {
//		total += record.Total
//	}
//
//	t := total / float64(len(records))
//	fmt.Println(t)
//}

// get first record
//

