package eddie

import (
	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
	"github.com/iods/go-eddie/internal/helpers/time"
	"github.com/iods/go-eddie/internal/util/parse"
)

// TrackSeizure Structures and creates a record for tracking a seizure.
func TrackSeizure(s string, tags []string, l string, isImportant bool) (err error) {
	db.InitDatabase()
	database := db.GetDatabase()

	t := parse.Tags(tags)
	r := &schema.Record{
		Type: "seizure",
		Time: time.UpdateTime(s),
		Tags: t,
		Location: l,
		Important: isImportant,
	}
	database.Create(&r)
	return err
}