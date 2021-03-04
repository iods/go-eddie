package db

import (
	"log"
	"path/filepath"

	"github.com/iods/go-eddie/internal/db/schema"
	util "github.com/iods/go-eddie/internal/util/project"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open(
		filepath.Join(util.GetProjectHome(), "records.db", // users home path
	)), &gorm.Config{
		// include any configs?
	})

	if err != nil {
		panic("failed to connect database")
	}

	// gorm will auto-generate the tables and schema on the fly
	err = db.AutoMigrate(&schema.Record{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func GetDatabase() *gorm.DB {
	return DB
}