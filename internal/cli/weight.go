package cli

import (
	"fmt"
	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
)

func TrackWeight(w string, i bool) (err error) {

	db.InitDatabase()
	data := db.GetDatabase()

	r := &schema.Record{
		Type: "weight",
		Total: w,
		Important: i,
	}

	data.Create(&r)

	fmt.Printf("you reported your weight at %s today.\n", w)
	fmt.Println("Record ID:", r.ID)
	return err
}