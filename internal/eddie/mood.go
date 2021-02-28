package eddie

import (
	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
	"github.com/iods/go-eddie/internal/util/parse"
)

// TrackMood Creates a record in the database for the fields related to the mood command.
func TrackMood(quality int, stress int, tags []string, emojis []string, isImportant bool) error {
	db.InitDatabase()
	database := db.GetDatabase()

	e := parse.Emojis(emojis)
	t := parse.Tags(tags)
	r := &schema.Record{
		Type: "mood",
		Quality: quality,
		Stress: stress,
		Tags: t,
		Emojis: e,
		Important: isImportant,
	}
	database.Create(&r)
	return nil
}