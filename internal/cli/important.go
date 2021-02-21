package cli

import "fmt"

func isImportant(b bool) {
	if b != false {
		fmt.Println("This record was important")
	}
}