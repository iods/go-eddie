package eddie

import (
	"fmt"
	"strconv"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
)

var (
	//format = "2006-01-02"
)

// TrackWeight Structures a weight record for insertion (create) into the database.
func TrackWeight(time string, important bool) (err error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var amount float64
	amount, err = strconv.ParseFloat(time, 64)
	if err != nil {
		panic(err)
	}
	r := &schema.Record{
		Type: "weight",
		Total: amount,
		Important: important,
	}
	database.Create(&r)
	return err
}

func (w Weight) RenderDashboard(a string) string {
	return fmt.Sprintf("%s", a)
}



//func GetWeightAverage(monthRange int, reverse bool) (float64, error) {
//	var from, to time.Time
//	var total = 0.0
//
//	if reverse == false {
//		r, _ := model.GetRecordsLast()
//		from = r.CreatedAt
//		to = from.AddDate(0, -monthRange, 0)
//	} else {
//		r, _ := model.GetRecordsFirst()
//		from = r.CreatedAt
//		to = from.AddDate(0, monthRange, 0)
//	}
//
//	f := from.Format(format)
//	t := to.Format(format)
//
//	results, err := weightModel.GetWeightRecordsByRangeCustom(t,f)
//	errors.Handle("handle this one day", err)
//
//	for _, v := range results {
//		total += v.Total
//	}
//
//	return total / float64(len(results)), nil
//}
//
//
//func GetWeightAverageByMonth(month int) (string, float64, float64) {
//	date := time.Now()
//	target := date.AddDate(0, -month, -1)
//
//	results, _ := weightModel.GetWeightRecordsByRange(target.Format(format), 1)
//	var b []float64
//	for _, result := range results {
//		b = append(b, result.Total)
//	}
//
//	var total = 0.0
//	for _, v := range b {
//		total += v
//	}
//
//	min, max := getMinMax(results)
//	avg := total / float64(len(b))
//
//	return fmt.Sprintf("%0.1f", avg), max.Total, min.Total
//}
//
//
//func GetWeightAverageTotal() (float64, string, string, error) {
//	var total = 0.0
//
//	results, err := weightModel.GetWeightRecords()
//	errors.Handle("handle this error one day", err)
//
//	for _, v := range results {
//		total += v.Total
//	}
//
//	min, max := getMinMax(results)
//	m1 := fmt.Sprintf("%0.1f", min.Total)
//	m2 := fmt.Sprintf("%0.1f", max.Total)
//
//	return total / float64(len(results)), m2, m1, nil
//}
//
//func GetWeightAverageTotalTrend() (string, string) {
//
//	results, err := weightModel.GetWeightRecords()
//	errors.Handle("handle it at some point", err)
//
//	min, max := getMinMax(results)
//
//	total := 0.0
//
//	for _, v := range results {
//		total += v.Total
//	}
//
//	tmp := total / float64(len(results))
//
//	l := tmp - min.Total
//	g := max.Total - tmp
//
//	return fmt.Sprintf("%0.1f", l), fmt.Sprintf("%0.1f", g)
//}