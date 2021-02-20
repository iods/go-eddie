package db

import (
	"fmt"
	"gorm.io/gorm"

	"github.com/iods/go-eddie/internal/db/schema"
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

	// fill some default track records
	one := &schema.Record{
		Type: "sleep",
		From: "23:58",
		To: "07:44",
		Length: 7,
		Quality: 9,
		Tags: []schema.Tag{
			{Name: "SNL"},
			{Name: "Code"},
			{Name: "Floor"},
			{Name: "Sweats"},
		},
		Important: false,
	}

	two := &schema.Record{
		Type: "mood",
		Quality: 7,
		Tags: []schema.Tag{
			{Name: "Resume"},
			{Name: "Golang"},
			{Name: "Trash Day"},
			{Name: "Mountain Dew"},
			{Name: "Sativa"},
			{Name: "MPH"},
		},
		Emojis: []schema.Emoji{
			{Name: "Excited"},
			{Name: "Bored"},
			{Name: "Anxious"},
			{Name: "Sad"},
		},
		Important: false,
	}

	three := &schema.Record{
		Type: "weight",
		Total: "194",
		Important: false,
	}

	four := &schema.Record{
		Type: "seizure",
		To: "11:33", // does this need to be a time.Time separate of to/from?
		Tags: []schema.Tag{
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
	newTag := []schema.Tag{{Name: "tinlicker"}}
	db.Omit("Tags").Updates(&two) // Update everything but tags.

	fmt.Println(two.Tags)

	err := db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&two).Association("Tags").Append(newTag)
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