package cli

import (
	"fmt"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
	"github.com/iods/go-eddie/internal/helpers/time"
	"github.com/iods/go-eddie/internal/util/parse"
)

// TrackSeizure
func TrackSeizure(s string, tags []string, l string, b bool) (err error) {
	db.InitDatabase()
	database := db.GetDatabase()

	t := parse.Tags(tags)

	r := &schema.Record{
		Type: "seizure",
		Time: time.UpdateTime(s),
		Tags: t,
		Location: l,
		Important: b,
	}

	database.Create(&r)

	fmt.Printf("You recorded a seizure at %v\n", r.Time)
	fmt.Println("Record ID:", r.ID)
	fmt.Println("Tags:", r.Tags)
	fmt.Println("Location:", r.Location)
	isImportant(b)
	return err
}