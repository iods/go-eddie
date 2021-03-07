package eddie

import (
	"github.com/iods/go-eddie/internal/db/models"
	"github.com/iods/go-eddie/internal/db/schema"
	"github.com/iods/go-eddie/internal/errors"
)

type Eddie struct {
	errors []string
}

var (
	model 	models.RecordModel
)

// FetchAll Returns all records for calculating frequency and participation of the application.
func (e Eddie) FetchAll() ([]schema.Record, error) {
	r, err := model.GetRecords()
	errors.Handle("eddie couldn't fetch those records", err)
	return r, nil
}

// FetchAllWeight Returns all weight records for dashboard views, calculating percentages and progress, etc.
func (e Eddie) FetchAllWeight() ([]schema.Record, error) {
	r, err := model.GetRecordsOfType("weight")
	errors.Handle("eddie couldn't fetch those records for weight", err)
	return r, nil
}

// FetchFirstWeight Returns first weight record by `created_at` for setting baselines; like starting weight.
func (e Eddie) FetchFirstWeight() (schema.Record, error) {
	r, err := model.GetRecordOfType("weight", true)
	errors.Handle("eddie couldn't fetch that record you were looking for", err)
	return r, nil
}

// FetchLastWeight Returns the most recent (or last) `created_at` record for calculating totals across weight numbers...
func (e Eddie) FetchLastWeight() (schema.Record, error) {
	r, err := model.GetRecordOfType("weight", false)
	errors.Handle("eddie couldn't fetch that record you were looking for", err)
	return r, nil
}

func (e Eddie) FetchAllWeightBetween(from string, to string) ([]schema.Record, error) {
	r, err := model.GetRecordsByDateRange(from, to, "weight")
	errors.Handle("eddie could not fetch those records between that range of dates", err)
	return r, nil
}

func (e Eddie) FetchAllBetween() ([]schema.Record, error) {
	r, err := model.GetRecordsBetween("sleep", "quality", 3, 5)
	errors.Handle("eddie couldn't get that record you are looking for.", err)
	return r, nil
}

func (e Eddie) FetchWeightCount() (int, error) {
	r, err := model.GetRecordsOfType("weight")
	errors.Handle("handle this error at some point", err)
	return len(r), nil
}