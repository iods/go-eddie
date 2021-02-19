package datasource

import (
	"log"
	"path/filepath"

	"github.com/iods/go-eddie/internal/model"
	"github.com/iods/go-eddie/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open(
		filepath.Join(utils.GetProjectDir(), "records.db", // users home path
	)), &gorm.Config{
		// include any configs?
	})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&model.Record{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func GetDatabase() *gorm.DB {
	return DB
}