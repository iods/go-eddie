package eddie

//
//import (
//	"strconv"
//	"time"
//
//	"github.com/iods/go-eddie/internal/db"
//	"github.com/iods/go-eddie/internal/db/schema"
//	"github.com/iods/go-eddie/internal/util"
//)
//
//// TrackMood Creates a record in the database for the fields related to the mood command.
//func TrackMood(quality int, stress int, tags []string, emojis []string, important bool) error {
//	db.InitDatabase()
//	database := db.GetDatabase()
//
//	e := util.ParseEmojis(emojis)
//	t := util.ParseTags(tags)
//	r := &schema.Record{
//		Type:      "mood",
//		Quality:   quality,
//		Stress:    stress,
//		Tags:      t,
//		Emojis:    e,
//		Important: important,
//	}
//	database.Create(&r)
//	return nil
//}
//
//// TrackSeizure Structures and creates a record for tracking a seizure.
//func TrackSeizure(seizureTime string, tags []string, location string, isImportant bool) error {
//	db.InitDatabase()
//	database := db.GetDatabase()
//
//	t := util.ParseTags(tags)
//	r := &schema.Record{
//		Type:      "seizure",
//		Time:      util.UpdateTime(seizureTime),
//		Tags:      t,
//		Location:  location,
//		Important: isImportant,
//	}
//	database.Create(&r)
//	return nil
//}
//
//// TrackSleep Structures and creates a record in the database for the `track sleep` command.
//func TrackSleep(sleepTime string, length int, quality int, tags []string, isImportant bool) error {
//	db.InitDatabase()
//	database := db.GetDatabase()
//
//	to := util.UpdateTime(sleepTime)
//	from := to.Add(time.Duration(-length) * time.Hour)
//	t := util.ParseTags(tags)
//	r := &schema.Record{
//		Type:      "sleep",
//		From:      from,
//		To:        to,
//		Length:    length,
//		Quality:   quality,
//		Tags:      t,
//		Important: isImportant,
//	}
//	database.Create(&r)
//	return nil
//}
//
//// TrackWeight Structures a weight record for insertion (create) into the database.
//func TrackWeight(time string, important bool) error {
//	db.InitDatabase()
//	database := db.GetDatabase()
//
//	var amount float64
//	amount, err := strconv.ParseFloat(time, 64)
//	if err != nil {
//		panic(err)
//	}
//	r := &schema.Record{
//		Type:      "weight",
//		Total:     amount,
//		Important: important,
//	}
//	database.Create(&r)
//	return nil
//}
