package db

//import (
//	"log"
//	"os"
//	"path/filepath"
//
//	"github.com/iods/go-eddie/internal/db/schema"
//	"github.com/iods/go-eddie/internal/util"
//	"gorm.io/driver/sqlite"
//	"gorm.io/gorm"
//)
//
//var DB *gorm.DB
//
//// InitDatabase Structures a connection between a source for accessing it's data.
//func InitDatabase() {
//	var filename string
//
//	if _, err := os.Stat(filepath.Join(util.GetProjectHome(), "/records-stub.db")); err == nil {
//		filename = "/records-stub.db"
//	} else {
//		filename = "/records.db"
//	}
//
//	db, err := gorm.Open(sqlite.Open(
//		filepath.Join(util.GetProjectHome(), filename,
//	)), &gorm.Config{
//
//	})
//
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	err = db.AutoMigrate(&schema.Record{})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	DB = db
//}
//
//// GetDatabase Returns a database object for accessing the records data.
//func GetDatabase() *gorm.DB {
//	return DB
//}
