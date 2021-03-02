package eddie

import (
	"github.com/iods/go-eddie/internal/db/schema"
	"github.com/iods/go-eddie/internal/errors"
	"time"
)

type Weight struct {}

var (
	eddie Eddie
	format = "2006-01-02"
)

// @TODO Re-factor this ASAP xD
func (w Weight) FindAverages(months int, reverse bool) (float64, error) {
	var from, to time.Time
	var total float64

	if months == 0 {
		r, _ := eddie.FetchAllWeight()
		total = 0.0
		for _, v := range r {
			total += v.Total
		}
		return total / float64(len(r)), nil
	}

	if months == 7 && reverse == false {
		r, _ := eddie.FetchLastWeight()
		from = r.CreatedAt
		to = from.AddDate(0, 0, -7)
	} else if reverse == false {
		r, _ := eddie.FetchLastWeight()
		from = r.CreatedAt
		to = from.AddDate(0, -months, 0)
	} else {
		r, _ := eddie.FetchFirstWeight()
		from = r.CreatedAt
		to = from.AddDate(0, months, 0)
	}

	f := from.Format(format)
	t := to.Format(format)

	results, err := eddie.FetchAllWeightBetween(t, f)
	errors.Handle("eddie will handle this error later", err)

	total = 0.0
	for _, v := range results {
		total += v.Total
	}

	return total / float64(len(results)), nil
}


func (w Weight) FindCount() int {
	r, _ := eddie.FetchLastWeight()
	from := r.CreatedAt
	to := from.AddDate(0, -1, 0)

	f := from.Format(format)
	t := to.Format(format)

	one, err := eddie.FetchAllWeightBetween(t, f)

	errors.Handle("handle this.", err)

	return len(one)
}


func (w Weight) FindBmi(wt float64, ht float64) (float64, error) {
	var bmi float64
	bmi = (wt * 703) / (ht * ht)
	return bmi, nil
}


func (w Weight) FindPercentage(weight float64) float64 {
	var diff float64
	v1 := 189.0
	v2 := weight

	if v1 < v2 {
		diff = (v2 - v1) / ((v1 + v2) / 2) * 100
	} else {
		diff = (v1 - v2) / ((v1 + v2) / 2) * 100
	}

	return diff
}

func (w Weight) FindMinMax(months int, reverse bool) (min schema.Record, max schema.Record) {
	var from, to time.Time

	if months == 0 {
		r, _ := eddie.FetchAllWeight()
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

	if months == 7 && reverse == false {
		r, _ := eddie.FetchLastWeight()
		from = r.CreatedAt
		to = from.AddDate(0, 0, -7)
	} else if reverse == false {
		r, _ := eddie.FetchLastWeight()
		from = r.CreatedAt
		to = from.AddDate(0, -months, 0)
	} else {
		r, _ := eddie.FetchFirstWeight()
		from = r.CreatedAt
		to = from.AddDate(0, months, 0)
	}

	f := from.Format(format)
	t := to.Format(format)

	results, err := eddie.FetchAllWeightBetween(t, f)
	errors.Handle("eddie will handle this error later", err)

	min = results[0]
	max = results[0]

	for _, result := range results {
		if result.Total > max.Total {
			max = result
		}
		if result.Total < min.Total {
			min = result
		}
	}
	return min, max
}