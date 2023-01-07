package errors

import (
	"fmt"
	"log"
	"os"
)

// Handle Handles the error, prints the message, and exits upon a failure.
func Handle(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

// HandleO handles the error by checking if nil, else panics.
func HandleO(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// HandleV Handles an error, printing a custom message, and then exits.
func HandleV(message string, err error) {
	if err != nil {
		fmt.Println(message, err)
		os.Exit(1)
	}
}

// ExitOnError returns a formatted error message when err is not nil.
func ExitOnError(err error) {
	if err == nil {
		return
	}
	_, err = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	if err != nil {
		return
	}
	os.Exit(1)
}
