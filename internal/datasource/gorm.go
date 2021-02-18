package datasource

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"

	"github.com/iods/go-eddie/internal/model"
	"github.com/iods/go-eddie/internal/utils"
)

func Init() {
	db, err := gorm.Open(sqlite.Open(filepath.Join(utils.GetProjectDir(), "records.db")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Record{})
	db.Create(&model.Record{
		Name:  "Example Record",
		Type:  "Sleep",
	})
	var records []model.Record
	db.Find(&records)
	data, _ := json.Marshal(&records)
	fmt.Println(data)
}