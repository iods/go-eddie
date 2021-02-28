// command Creates a "controller" like boilerplate for various functionality and access to models across the application
package cli

import (
	"github.com/iods/go-eddie/internal/db/models"
	"github.com/iods/go-eddie/internal/db/schema"
	"github.com/iods/go-eddie/internal/errors"
)

type Eddie struct {}

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
