package template

import (
	"fmt"
	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
)

func RenderSeizure() {

	var records []schema.Record

	db.InitDatabase()
	database := db.GetDatabase()

	r := database.Where("type LIKE ?", "%seizure%").Find(&records)
	fmt.Println(r.RowsAffected)
}