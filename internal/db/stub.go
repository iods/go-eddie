package db

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/iods/go-eddie/internal/db/schema"
)


func stubMood()  {
	var stub schema.Record
	gofakeit.Struct(&stub)

	tags := stub.Tags
	emojis := stub.Emojis

	mood := &schema.Record{

		Quality: stub.Quality,
		Tags: tags,
		Emojis: emojis,
		Important: stub.Important,
	}

	fmt.Println("You recorded your mood a", mood.Quality)
	fmt.Println("You recorded the following tags with your mood:", mood.Tags)
	fmt.Println("You recorded the following emojis with your mood:", mood.Emojis)
	if mood.Important != false {
		fmt.Println("You recorded this record as important.")
	}
}

func StubDatabase() {

	for i := 0; i < 5; i++ {
		stubMood()
	}
}
