package cli

import (
	"fmt"
	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/model"
)

func TrackMood() (err error) {

	db.InitDatabase()
	database := db.GetDatabase()

	r := &model.Record{
		Type: "weight",

	}

	database.Create(&r)

	fmt.Println("Record ID:", r.ID)
	return err
}

func UpdateMood() (err error) {

	return err
}