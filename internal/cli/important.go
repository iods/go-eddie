package cli

import "fmt"

func isImportant(b bool) {
	// will do something to these
	// SELECT stmt pulling all important ones maybe, then filter then...
	if b != false {
		fmt.Println("This record was important")
	}
}