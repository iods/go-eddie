package error

import (
	"fmt"
	"os"
)

// Handle Handles the error, prints the message, and exits upon a failure.
func Handle(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}