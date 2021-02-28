package models

import (
	"time"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
	"github.com/iods/go-eddie/internal/errors"
)

type RecordModel struct {}

var (
	format = "2006-01-02"
	results []schema.Record
)

// GetRecords Returns all entities from the datasource for ...
func (r RecordModel) GetRecords() ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	database.Debug().Find(&results)
	return results, nil
}

// GetRecordsFirst Returns the first created_at record in the database for use with filtering and calculations.
func (r RecordModel) GetRecordsFirst() (schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var result schema.Record
	database.Debug().Order("created_at asc").First(&result)
	return result, nil
}

// GetRecordsLast Returns the most recent created_at record in the database for use with filtering and calculations.
func (r RecordModel) GetRecordsLast() (schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var result schema.Record
	database.Debug().Order("created_at desc").First(&result)
	return result, nil
}

// GetRecordsByDate Extends GetRecords by returning records based on a specified date for ...
func (r RecordModel) GetRecordsByDate(d string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	q, _ := time.Parse(format, d)
	database.Debug().Where("created_at = ?", q).Find(&results)
	return results, nil
}

// GetRecordsByDateRange Returns a set of records created between a provided date range for ...
func (r RecordModel) GetRecordsByDateRange(from string, to string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	f, err := time.Parse(format, from)
	t, err := time.Parse(format, to)
	errors.Handle("issue", err)
	database.Debug().Where("created_at BETWEEN ? AND ?", f, t).Find(&results)
	return results, nil
}

// GetRecordsByImportance Returns a slice of records based on their importance, true or false.
func (r RecordModel) GetRecordsByImportance() ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	database.Debug().Find(&results, "important = ?", true)
	return results, nil
}