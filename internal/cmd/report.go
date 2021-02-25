package cmd

import (
	"fmt"
	"github.com/iods/go-eddie/internal/cli/template"
	"github.com/iods/go-eddie/internal/db/models"
	"github.com/iods/go-eddie/internal/db/schema"
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

			d, err := model.GetRecordsByImportance()
			errors.Handle("could not get important records", err)

			for _, r4 := range d {
				fmt.Printf("id: %d\ncreated at: %v\nimportance: %v\n", r4.ID, r4.CreatedAt, r4.Important)
			}
		case "mood":

		case "seizure":

		case "sleep":
			template.RenderWeight()
		case "weight":
			var model models.WeightModel

			// Average (1 month)
			one, err := model.GetWeightRecordsByDateRange("2021-01-24", "2021-02-24")
			errors.Handle("something went wrong while getting the records", err)

			var o []float64
			for _, v := range one {
				o = append(o, v.Total)
			}

			a1 := getAverage(o)
			min, max := getMinMax(one)

			fmt.Printf("Your average for the past month is %.0f (%0.2f) and max was %v and min was %v\n", a1, a1, max.Total, min.Total)

			// Average (3 months)
			three, err := model.GetWeightRecordsByDateRange("2020-11-22", "2021-02-22")
			errors.Handle("something bad happened when getting the range records.", err)

			var t []float64
			for _, v := range three {
				t = append(t, v.Total)
			}

			a3 := getAverage(t)
			fmt.Printf("Your average for three months is %.0f (%.02f)\n", a3, a3)

			// Average (6 months)
			six, err := model.GetWeightRecordsByDateRange("2020-08-22", "2021-02-22")
			errors.Handle("something bad happened when getting the range records.", err)

			var s []float64
			for _, v := range six {
				s = append(s, v.Total)
			}

			a6 := getAverage(s)
			fmt.Printf("Your average for six months is %.0f (%.02f)\n", a6, a6)

			// Average (9 months)
			nine, err := model.GetWeightRecordsByDateRange("2020-05-22", "2021-02-22")
			errors.Handle("something bad happened when getting the range records.", err)

			var n []float64
			for _, v := range nine {
				n = append(n, v.Total)
			}

			a9 := getAverage(n)
			fmt.Printf("Your average for nine months is %.0f (%.02f)\n", a9, a9)

			// Average (total)
			total, err := model.GetWeightRecordsByDateRange("2020-01-01", "2021-02-22")
			errors.Handle("something happened while getting your records", err)

			var a []float64
			for _, v := range total {
				a = append(a, v.Total)
			}

			all := getAverage(a)
			fmt.Printf("Your total weight average is %.0f (%.02f)\n", all, all)

			m1, m2 := getMinMax(total)
			fmt.Printf("Your min/max is \nmin:%v (created :%v)\nmax:%v (%v)\n", m1.Total, m1.CreatedAt, m2.Total, m2.CreatedAt)


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