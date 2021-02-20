package cli

import (
	"fmt"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
)

func TrackMood(q int, t []string, i bool) (err error) {

	db.InitDatabase()
	database := db.GetDatabase()

	testStrings(t)

	r := &schema.Record{
		Type: "mood",
		Quality: q,
		Important: i,
		Tags: testStrings(t),
	}

	database.Create(&r)

	fmt.Printf("You rated your mood a %d.\n", q)
	fmt.Println("Record ID:", r.ID)
	fmt.Println("Tags:", r.Tags)
	return err
}

func testStrings(strings []string) []schema.Tag {

	// strings = []string{"hello", "world"}


	var tag []schema.Tag

	for _, s := range strings {
		tag = []schema.Tag{{Name: s}}
	}

	return tag

}

func mood(q int, tag []string, f bool) {
	fmt.Println("You rated your mood at a", q)
	for i := 0; i < len(tag); i++ {
		fmt.Println(tag[i])
	}
}