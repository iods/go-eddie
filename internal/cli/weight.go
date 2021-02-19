package cli

import (
	"fmt"
	"github.com/iods/go-eddie/internal/datasource"
	"github.com/iods/go-eddie/internal/model"
)

func TrackWeight(w string, i bool) (err error) {

	datasource.InitDatabase()
	db := datasource.GetDatabase()

	r := &model.Record{
		Type: "weight",
		Total: w,
		Important: i,
	}

	db.Create(&r)

	fmt.Printf("you reported your weight at %s today.\n", w)
	fmt.Println("Record ID:", r.ID)
	return err
}

