package datasource

import (
	"fmt"
	"gorm.io/gorm"
	"log"

	"github.com/iods/go-eddie/internal/model"
)

func Init() {
	/*

	Database Connection

	 */
	InitDatabase()
	db := GetDatabase()

	/*

	Creating Schema & Data

	 */
	// gorm will auto-generate the tables and schema on the fly
	err := db.AutoMigrate(&model.Record{})
	if err != nil {
		log.Fatal(err)
	}

	// fill some default track records
	one := &model.Record{
		Type: "sleep",
		EndTime: "07:44",
		Duration: 7,
		Quality: 9,
		Tags: []model.Tag{
			{Name: "SNL"},
			{Name: "Code"},
			{Name: "Floor"},
			{Name: "Sweats"},
		},
		Important: false,
	}

	two := &model.Record{
		Type: "mood",
		Quality: 7,
		Tags: []model.Tag{
			{Name: "Resume"},
			{Name: "Golang"},
			{Name: "Trash Day"},
			{Name: "Mountain Dew"},
			{Name: "Sativa"},
			{Name: "MPH"},
		},
		Emojis: []model.Emoji{
			{Name: "Excited"},
			{Name: "Bored"},
			{Name: "Anxious"},
			{Name: "Sad"},
		},
		Important: false,
	}

	three := &model.Record{
		Type: "weight",
		Total: "194",
		Important: false,
	}

	four := &model.Record{
		Type: "seizure",
		EndTime: "11:33",
		Tags: []model.Tag{
			{Name: "caffeine"},
			{Name: "video-games"},
			{Name: "flashing-lights"},
			{Name: "politics"},
		},
		Location: "Strip Club",
		Important: true,
	}

	db.Create(&one)
	fmt.Println(one.ID)

	db.Create(&two)
	fmt.Println(two.ID)

	db.Create(&three)
	fmt.Println(three.ID)

	db.Create(&four)
	fmt.Println(four.ID)

	/*

	Reading Data


	var record model.Record

	db.First(&record, 1) // find the record with the integer primary key
	db.First(&record, "type = ?", "sleep") // find record with the type weight
	*/



	/*

		Updating Data


	db.Model(&record).Update("Quality", 8) // update the records name to 'updated'
	db.Model(&record).Updates(model.Record{Important: true}) // non-zero fields
	*/
	newTag := []model.Tag{{Name: "tinlicker"}}
	db.Omit("Tags").Updates(&two) // Update everything but tags.

	fmt.Println(two.Tags)

	err = db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&two).Association("Tags").Append(newTag)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ID: ", two.ID)
	fmt.Println("Created At:", two.CreatedAt)
	fmt.Println("Type:", two.Type)
	fmt.Println("Quality:", two.Quality)
	fmt.Println("Tags:", two.Tags)
	fmt.Println("Emojis:", two.Emojis)
	fmt.Println("Is Important:", two.Important)
	/*

	Deleting Data


	db.Delete(&record, 1) // delete by the id
	*/


}