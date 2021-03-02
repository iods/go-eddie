package models

import (
	"fmt"
	"time"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
	"github.com/iods/go-eddie/internal/errors"
)

type RecordModel struct {}

var (
	format  = "2006-01-02"
	result  schema.Record
	results []schema.Record
)


// GetRecords Returns all entities from the datasource for ...
func (r RecordModel) GetRecords() ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	database.Find(&results)
	return results, nil
}


// GetRecordsBetween Returns a slice of records based on a record type, a field, and a range.
func (r RecordModel) GetRecordsBetween(recordType string, flag string, min float64, max float64) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	query := fmt.Sprintf("%s BETWEEN %f AND %f", flag, min, max)
	database.Where("type = ?", recordType).Where(query).Find(&results)
	return results, nil
}


// GetRecordsOfType Returns records of a specified type for filtering and extending GetRecords.
func (r RecordModel) GetRecordsOfType(recordType string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	database.Where("type = ?", recordType).Find(&results)
	return results, nil
}


// GetRecordOfType Returns a single record (first or last created_at) of the type specified.
func (r RecordModel) GetRecordOfType(recordType string, asc bool) (schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var value string
	if asc == false {
		value = "created_at desc"
	} else {
		value = "created_at"
	}
	database.Where("type = ?", recordType).Order(value).First(&result)
	return result, nil
}


// GetRecordsByDate Extends GetRecords by returning records based on a specified date for ...
func (r RecordModel) GetRecordsByDate(d string, recordType string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	query, _ := time.Parse(format, d)
	database.Where("created_at = ? AND type = ?", query, recordType).Find(&results)
	return results, nil
}


// GetRecordsByDateRange Returns a set of records created between a provided date range for ...
func (r RecordModel) GetRecordsByDateRange(from string, to string, recordType string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	f, err := time.Parse(format, from)
	t, err := time.Parse(format, to)
	errors.Handle("issue", err)
	database.Where("created_at BETWEEN ? AND ?", f, t).Where("type = ?", recordType).Find(&results)
	return results, nil
}


// GetRecordsStartingWith Returns a slice of related records.
func (r RecordModel) GetRecordsStartingWith(keyword string, flag string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	database.Where(flag + " LIKE ?", keyword + "%").Find(&results)
	return results, nil
}


// GetRecordsStartingWith Returns a slice of related records.
func (r RecordModel) GetRecordsContaining(keyword string, flag string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	database.Where(flag + " LIKE ?", "%" + keyword + "%").Find(&results)
	return results, nil
}


// GetRecordsStartingWith Returns a slice of related records.
func (r RecordModel) GetRecordsEndingWith(keyword string, flag string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	database.Where(flag + " LIKE ?", "%" + keyword).Find(&results)
	return results, nil
}


// GetRecordsByImportance Returns a slice of records based on their importance, true or false.
func (r RecordModel) GetRecordsByImportance(recordType string) ([]schema.Record, error) {
	db.InitDatabase()
	database := db.GetDatabase()

	database.Find(&results, "type = ? AND important = true", recordType)
	return results, nil
}