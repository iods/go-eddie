// command Creates a "controller" like boilerplate for various functionality and access to models across the application
package cli

import (
	"time"

	"github.com/iods/go-eddie/internal/db/models"
	"github.com/iods/go-eddie/internal/errors"
)

var model models.RecordModel


// GetFirstCreatedAt Returns the first records created_at value in time format for range values in other functions.
func GetFirstCreatedAt() time.Time {
	c, err := model.GetRecordsFirst()
	errors.Handle("handle this error at some point", err)

	return c.CreatedAt
}

// GetLastCreatedAt Returns the last created_at (most recent) record in time format for range values in other functions.
func GetLastCreatedAt() time.Time {
	c, err := model.GetRecordsLast()
	errors.Handle("handle this error at some point", err)

	return c.CreatedAt
}