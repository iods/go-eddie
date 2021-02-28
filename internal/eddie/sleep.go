package eddie

import (
	"time"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
	time2 "github.com/iods/go-eddie/internal/helpers/time"
	"github.com/iods/go-eddie/internal/util/parse"
)

// TrackSleep Structures and creates a record in the database for the `track sleep` command.
func TrackSleep(s string, length int, quality int, tags []string, isImportant bool) error {
	db.InitDatabase()
	database := db.GetDatabase()

	to := time2.UpdateTime(s)
	from := to.Add(time.Duration(-length) * time.Hour)
	t := parse.Tags(tags)
	r := &schema.Record{
		Type: "sleep",
		From: from,
		To: to,
		Length: length,
		Quality: quality,
		Tags: t,
		Important: isImportant,
	}
	database.Create(&r)
	return nil
}