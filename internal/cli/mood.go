package cli

import (
	"fmt"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
)

func TrackMood(q int, tags []string, emojis []string, i bool) (err error) {

	db.InitDatabase()
	database := db.GetDatabase()

	// create tool for reading records from DB in go-pherit, using GORM and sqlite in home dir

	t := parseTags(tags)
	e := parseEmoji(emojis)

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
	return err
}

func parseEmoji(e []string) []schema.Emoji {
	var emojis []schema.Emoji

	l := len(e)
	for i := 0; i < l; i++ {
		emoji := []schema.Emoji{{Name: e[i]}}
		emojis = append(emojis, emoji...)
	}

	return emojis
}

func parseTags(t []string) []schema.Tag {
	var tags []schema.Tag

	l := len(t)
	for i := 0; i < l; i++ {
		tag := []schema.Tag{{Name: t[i]}}
		tags = append(tags, tag...)
	}

	return tags
}