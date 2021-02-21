package cli

import (
	"fmt"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
	"github.com/iods/go-eddie/internal/util/parse"
)

func TrackMood(q int, tags []string, emojis []string, i bool) (err error) {

	db.InitDatabase()
	database := db.GetDatabase()

	e := parse.Emojis(emojis)
	t := parse.Tags(tags)

	r := &schema.Record{
		Type: "mood",
		Quality: q,
		Tags: t,
		Emojis: e,
		Important: i,
	}

	database.Create(&r)

	fmt.Printf("You rated your mood a %d.\n", q)
	fmt.Println("Record ID:", r.ID)
	fmt.Println("Tags:", r.Tags)
	fmt.Println("Emojis", r.Emojis)
	isImportant(i)

	return err
}