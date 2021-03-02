package eddie

import (
	"fmt"
	"strconv"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
)

var (
	//format = "2006-01-02"
)

// TrackWeight Structures a weight record for insertion (create) into the database.
func TrackWeight(time string, important bool) (err error) {
	db.InitDatabase()
	database := db.GetDatabase()

	var amount float64
	amount, err = strconv.ParseFloat(time, 64)
	if err != nil {
		panic(err)
	}
	r := &schema.Record{
		Type: "weight",
		Total: amount,
		Important: important,
	}
	database.Create(&r)
	fmt.Println("id: ", r.ID, "was entered.")
	return err
}