package cli

import "fmt"

// isImportant Tags records that are of importance in one way or another.
func isImportant(b bool) {
	// will do something to these
	// SELECT stmt pulling all important ones maybe, then filter then...
	// --important, -i flag
	if b != false {
		fmt.Println("This record was important")
	}
}