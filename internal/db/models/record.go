package models

import (
	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
	"github.com/iods/go-eddie/internal/errors"
	"time"
)

type RecordModel struct {}

var (
	format = "2006-01-02"
)

// GetRecords Returns all entities from the datasource for ...
func (r RecordModel) GetRecords() ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var results []schema.Record

	database.Debug().Find(&results)

	return results, nil
}

// GetRecordsByDate Extends GetRecords by returning records based on a specified date for ...
func (r RecordModel) GetRecordsByDate(d string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var results []schema.Record

	q, _ := time.Parse(format, d)
	database.Debug().Where("created_at = ?", q).Find(&results)

	return results, nil
}

// GetRecordsByDateRange Returns a set of records created between a provided date range for ...
func (r RecordModel) GetRecordsByDateRange(from string, to string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()
	var results []schema.Record

	f, err := time.Parse(format, from)
	t, err := time.Parse(format, to)
	errors.Handle("issue", err)

	database.Debug().Where("created_at BETWEEN ? AND ?", f, t).Find(&results)

	return results, nil
}

