package db

import (
	"fmt"
	"gorm.io/gorm"

	"github.com/iods/go-eddie/internal/db/schema"
)

func Test() {
	/*

	Database Connection

	 */
	InitDatabase()
	db := GetDatabase()

	/*

	Creating Schema & Data

	 */
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
		Total: 194,
		Important: false,
	}

	db.Create(&two)
	fmt.Println(two.ID)

	db.Create(&three)
	fmt.Println(three.ID)


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