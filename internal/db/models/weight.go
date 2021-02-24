package models

import (
	"github.com/iods/go-eddie/internal/errors"
	"time"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
)

type WeightModel struct {}

// GetWeightRecords Returns all records of type `weight` for ...
func (w WeightModel) GetWeightRecords() ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var results []schema.Record

	database.Where("type = ?", "weight").Find(&results)

	return results, nil
}

// GetWeightRecordsByDate Returns a set of records within the date provided for ...
func (w WeightModel) GetWeightRecordsByDate(d string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var results []schema.Record

	q, _ := time.Parse("2006-01-02", d)
	database.Debug().Where("type = ?", "weight").Where("created_at = ?", q).Find(&results)

	return results, nil
}

func (w WeightModel) GetWeightRecordsByDateRange(from string, to string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var results []schema.Record

	f, err := time.Parse("2006-01-02", from)
	t, err := time.Parse("2006-01-02", to)
	errors.Handle("unable to parse these dates", err)

	database.Debug().Where("created_at BETWEEN ? AND ?", f, t).Where("type = ?", "weight").Find(&results)

	return results, nil
}







func (w WeightModel) FindByDates(start time.Time, finish time.Time) ([]schema.Record, error) {
	var result []schema.Record

	db.InitDatabase()
	database := db.GetDatabase()
	database.Where("created_at BETWEEN ? AND ?", start, finish)
	database.Where("type = ?", "weight").Find(&result)

	return result, nil
}


func (w WeightModel) FindByMonth(month string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var results []schema.Record

	start, _ := time.Parse("2006-01-02 3:04PM", "2020-" + month + "-01 0:00AM")
	end, _ := time.Parse("2006-01-02 3:04PM",  "2020-" + month + "-31 0:00AM")

	database.Where("created_at BETWEEN ? AND ?", start, end)
	database.Where("type = ?", "weight")

	database.Find(&results)

	return results, nil
}