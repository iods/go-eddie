package template

import (
	"fmt"
	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
)

func RenderSleep() {

	var records []schema.Record

	db.InitDatabase()
	database := db.GetDatabase()

	r := database.Where("type LIKE ?", "%sleep%").Find(&records)
	fmt.Println(r.RowsAffected)
}