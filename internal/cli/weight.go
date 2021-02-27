package cli

import (
	"fmt"
	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/models"
	"github.com/iods/go-eddie/internal/db/schema"
	"github.com/iods/go-eddie/internal/errors"
	"strconv"
	"time"
)

type WeightCommand struct {}

var (
	format = "2006-01-02"
	weightModel = models.WeightModel{}
)

// TrackWeight Structures a weight record for insertion (create) into the database.
func TrackWeight(t string, i bool) (err error) {

	db.InitDatabase()
	database := db.GetDatabase()

	var amount float64

	amount, err = strconv.ParseFloat(t, 64)
	if err != nil {
		panic(err)
	}

	r := &schema.Record{
		Type: "weight",
		Total: amount,
		Important: i,
	}

	database.Create(&r)

	fmt.Printf("you reported your weight at %f today.\n", r.Total)
	isImportant(i)
	fmt.Println("Record ID:", r.ID)

	return err
}


func GetWeightAverage(monthRange int, reverse bool) (float64, error) {
	var from, to time.Time
	var total = 0.0

	if reverse == false {
		r, _ := model.GetRecordsLast()
		from = r.CreatedAt
		to = from.AddDate(0, -monthRange, 0)
	} else {
		r, _ := model.GetRecordsFirst()
		from = r.CreatedAt
		to = from.AddDate(0, monthRange, 0)
	}

	f := from.Format(format)
	t := to.Format(format)

	results, err := weightModel.GetWeightRecordsByRangeCustom(t,f)
	errors.Handle("handle this one day", err)

	for _, v := range results {
		total += v.Total
	}

	return total / float64(len(results)), nil
}


func GetWeightAverageByMonth(month int) (string, float64, float64) {
	date := time.Now()
	target := date.AddDate(0, -month, -1)

	results, _ := weightModel.GetWeightRecordsByRange(target.Format(format), 1)
	var a string
	var b []float64
	for _, result := range results {
		a = result.CreatedAt.UTC().Format("Jan")
		b = append(b, result.Total)
	}

	var total = 0.0
	for _, v := range b {
		total += v
	}

	c := fmt.Sprintf("%.1f", total / float64(len(b)))

	return a, c, total / float64(len(b))
}


func GetWeightAverageTotal() (float64, error) {
	var total = 0.0

	results, err := weightModel.GetWeightRecords()
	errors.Handle("handle this error one day", err)

	for _, v := range results {
		total += v.Total
	}
	return total / float64(len(results)), nil
}












func (w WeightCommand) ReportWeightResults() ([]schema.Record, error) {
	var model models.WeightModel

	r, err := model.GetWeightRecords()
	errors.Handle("issue with all records", err)

	return r, nil
}

func toFloat() []float64 {
	var c WeightCommand
	a, err := c.ReportWeightResults()
	errors.Handle("error", err)

	var b []float64
	for _, result := range a {
		b = append(b, result.Total)
	}
	return b
}


func FindThirtyDays() {
	var model models.WeightModel

	// Average (1 month)
	one, err := model.GetWeightRecordsByRange("2021-01-24", 1)
	errors.Handle("something went wrong while getting the records", err)

	var o []float64
	for _, v := range one {
		o = append(o, v.Total)
	}

	a1 := getAverage(o)
	min, max := getMinMax(one)

	fmt.Printf("Your average for the past month is %.0f (%0.2f) and max was %v and min was %v\n", a1, a1, max.Total, min.Total)
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