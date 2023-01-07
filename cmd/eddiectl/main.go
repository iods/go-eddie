package main

import (
	"fmt"
	"os"

	"github.com/iods/go-eddie/internal/cmd/root"
)

func main() {
	rootCmd := root.EddiectlRootCmd()
	if _, err := rootCmd.ExecuteC(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "%s\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
