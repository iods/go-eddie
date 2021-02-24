package db

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/iods/go-eddie/internal/db/schema"
	"gorm.io/gorm"
)

func stubMood() *schema.Record {

	InitDatabase()
	database := GetDatabase()

	var r schema.Record
	gofakeit.Struct(&r)

	now := time.Now()
	then := now.AddDate(0, -13, -31)
	tmp := gofakeit.DateRange(then, now)
	d := tmp.Format("2006-01-02")
	date, _ := time.Parse("2006-01-02", d)

	sleep := &schema.Record{
		Model:     gorm.Model{
			CreatedAt: date,
		},
		Length:    0,
		Emojis:    r.Emojis,
		Frequency: 0,
		From:      time.Time{},
		Important: r.Important,
		Location:  "",
		Quality:   r.Quality,
		Tags:      r.Tags,
		To:        time.Time{},
		Time:      date,
		Total:     0,
		Type:      "mood",
	}

	database.Create(&sleep)
	return sleep
}

func stubSleep() *schema.Record {
	InitDatabase()
	database := GetDatabase()

	var r schema.Record
	gofakeit.Struct(&r)

	now := time.Now()
	then := now.AddDate(0, -13, -31)
	tmp := gofakeit.DateRange(then, now)
	d := tmp.Format("2006-01-02")
	date, _ := time.Parse("2006-01-02", d)
	datefrom := date.Add(time.Duration(-r.Length) * time.Hour)

	sleep := &schema.Record{
		Model:     gorm.Model{
			CreatedAt: date,
		},
		Length:    r.Length,
		Emojis:    nil,
		Frequency: 0,
		From:      datefrom,
		Important: r.Important,
		Location:  r.Location,
		Quality:   r.Quality,
		Tags:      r.Tags,
		To:        date,
		Time:      date,
		Total:     0,
		Type:      "sleep",
	}

	database.Create(&sleep)
	return sleep
}

func stubSeizure() *schema.Record {
	InitDatabase()
	database := GetDatabase()

	var r schema.Record
	gofakeit.Struct(&r)

	now := time.Now()
	then := now.AddDate(0, -13, -31)
	tmp := gofakeit.DateRange(then, now)
	d := tmp.Format("2006-01-02")
	date, _ := time.Parse("2006-01-02", d)

	seizure := &schema.Record{
		Model:     gorm.Model{
			CreatedAt: date,
		},
		Length:    0,
		Emojis:    nil,
		Frequency: 0,
		From:      time.Time{},
		Important: r.Important,
		Location:  r.Location,
		Quality:   0,
		Tags:      r.Tags,
		To:        date,
		Time:      date,
		Total:     0,
		Type:      "seizure",
	}

	database.Create(&seizure)
	return seizure
}

func stubWeight() *schema.Record {
	InitDatabase()
	database := GetDatabase()

	var r schema.Record
	gofakeit.Struct(&r)

	now := time.Now()
	then := now.AddDate(0, -13, -31)
	tmp := gofakeit.DateRange(then, now)
	d := tmp.Format("2006-01-02")
	date, _ := time.Parse("2006-01-02", d)

	weight := &schema.Record{
		Model:     gorm.Model{
			CreatedAt: date,
		},
		Length:    0,
		Emojis:    nil,
		Frequency: 0,
		From:      time.Time{},
		Important: r.Important,
		Location:  "",
		Quality:   0,
		Tags:      nil,
		To:        time.Time{},
		Time:      date,
		Total:     r.Total,
		Type:      "weight",
	}

	database.Create(&weight)
	return weight
}

func stub(f func() *schema.Record, c int) {
	for i := 0; i < c; i++ {
		f()
	}
}

func StubDatabase() {
	stub(stubMood, 300)
	stub(stubSleep, 365)
	stub(stubSeizure, 45)
	stub(stubWeight, 130)
}