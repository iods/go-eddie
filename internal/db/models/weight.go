package models

import (
	"github.com/iods/go-eddie/internal/errors"
	"time"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
)

type WeightModel struct {}

// GetWeightRecords Returns all records of type `weight` for filtering with and rendering in dashboard views.
func (w WeightModel) GetWeightRecords() ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var results []schema.Record

	database.Where("type = ?", "weight").Find(&results)

	return results, nil
}

// GetWeightRecordsByDate Returns a slice of records matching the created_at date provided for ...
func (w WeightModel) GetWeightRecordsByDate(d string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var results []schema.Record

	q, _ := time.Parse("2006-01-02", d)
	database.Debug().Where("type = ?", "weight").Where("created_at = ?", q).Find(&results)

	return results, nil
}

// GetWeightRecordsByDateRangeCustom Returns a set of records by searching two dates for...
func (w WeightModel) GetWeightRecordsByRange(from string, monthRange int) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var results []schema.Record

	f, err := time.Parse("2006-01-02", from)
	t := f.AddDate(0, monthRange, -1)

	errors.Handle("unable to parse these dates", err)

	database.Debug().Where("created_at BETWEEN ? AND ?", f, t).Where("type = ?", "weight").Find(&results)

	return results, nil
}


func (w WeightModel) GetWeightRecordsByRangeCustom(from string, to string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var results []schema.Record

	f, err := time.Parse("2006-01-02", from)
	t, err := time.Parse("2006-01-02", to)

	errors.Handle("unable to parse these.", err)

	database.Debug().Where("created_at BETWEEN ? AND ?", f, t).Where("type = ?", "weight").Find(&results)

	return results, nil
}






// GetWeightRecordsByImportance Returns a slice of records for filtering or comparing with others.
func (w WeightModel) GetWeightRecordsByImportance() ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()
	var results []schema.Record

	database.Debug().Find(&results, "type = ? AND important = ?", "weight", true)

	return results, nil
}