package project

import (
	"os"
)

// Install Builds the app for use.
func Install() {
	p := setDir()
	setupConfigFile(p)
	setupRecordsFile(p)
}

// Check Returns a bool for triggering the install of the app.
func Check() bool {
	p := setDir()
	if s, err := os.Stat(p); err == nil && s.IsDir() {
		return true // if the directory exists
	}
	return false
}

// Getdir Returns the project dir for access separate of `setDir()`
func Getdir() string {
	return setDir()
}
//func toFloat() []float64 {
//	var c WeightCommand
//	a, err := c.ReportWeightResults()
//	errors.Handle("error", err)
//
//	var b []float64
//	for _, result := range a {
//		b = append(b, result.Total)
//	}
//	return b
//}
//
//func getAverage(n []float64) float64 {
//	total := 0.0
//
//	for _, v := range n {
//		total += v
//	}
//
//	return total / float64(len(n))
//}
//
//func getMinMax(r []schema.Record) (min schema.Record, max schema.Record) {
//	min = r[0]
//	max = r[0]
//
//	for _, result := range r {
//		if result.Total > max.Total {
//			max = result
//		}
//		if result.Total < min.Total {
//			min = result
//		}
//	}
//
//	return min, max
//}
//
//
//func getAverage(n []float64) float64 {
//	total := 0.0
//
//	for _, v := range n {
//		total += v
//	}
//
//	return total / float64(len(n))
//}
//
//func getMinMax(r []schema.Record) (min schema.Record, max schema.Record) {
//	min = r[0]
//	max = r[0]
//
//	for _, result := range r {
//		if result.Total > max.Total {
//			max = result
//		}
//		if result.Total < min.Total {
//			min = result
//		}
//	}
//
//	return min, max
//}
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

