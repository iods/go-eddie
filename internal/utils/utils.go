package utils

import (
	"fmt"
	"os"
)

// HandleError Handles the error and exits upon a failure.
func HandleError(msg string, err error) {

	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}
