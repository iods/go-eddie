package models

import (
	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
	"time"
)

type WeightModel struct {}

func (w WeightModel) FindByDate(date time.Time) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()
	var result []schema.Record
	database.Where("created_at = ?", date).Find(&result)
	return result, nil
}

func (w WeightModel) FindByDates(start time.Time, finish time.Time) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()
	var result []schema.Record



	database.Where("created_at >= ? and created_at <= ?", start, finish).Find(&result)
	return result, nil
}

func (w WeightModel) GetRecords() ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()
	var results []schema.Record
	database.Debug().Where("type = ?", "weight").Find(&results)
	return results, nil
}