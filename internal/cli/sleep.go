package cli

import (
	"fmt"
	"time"

	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
	helper "github.com/iods/go-eddie/internal/helpers/time"
	"github.com/iods/go-eddie/internal/util/parse"
)

func TrackSleep(s string, l int, q int, tags []string, b bool) error {
	db.InitDatabase()
	database := db.GetDatabase()


	to := helper.UpdateTime(s)
	from := to.Add(time.Duration(-l) * time.Hour)
	t := parse.Tags(tags)
	r := &schema.Record{
		Type: "sleep",
		From: from,
		To: to,
		Length: l,
		Quality: q,
		Tags: t,
		Important: b,
	}

	database.Create(&r)
	fmt.Println("Record ID:", r.ID)

	return nil
}