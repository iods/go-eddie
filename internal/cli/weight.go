package cli

import (
	"fmt"
	"strconv"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
)

// TrackWeight Structures a weight record for insertion (create) into the database.
func TrackWeight(t string, i bool) (err error) {

	db.InitDatabase()
	database := db.GetDatabase()

	var amount float64

	amount, err = strconv.ParseFloat(t, 64)
	if err != nil {
		panic(err)
	}

	r := &schema.Record{
		Type: "weight",
		Total: amount,
		Important: i,
	}

	database.Create(&r)

	fmt.Printf("you reported your weight at %f today.\n", r.Total)
	isImportant(i)
	fmt.Println("Record ID:", r.ID)

	return err
}