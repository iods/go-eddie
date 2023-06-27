package eddie

/*
  Eddie: A CLI toolset and service dog.
  Copyright (c) 2017-Present, Rye Miller
  https://ryemiller.io
*/

import (
	"fmt"

	"github.com/spf13/cobra"
)

var test bool

var (
	fetchCmd = &cobra.Command{
		Use:   "fetch",
		Short: "Have Eddie retrieve various bits of information.",
		Long: `Eddie is great at retrieving things. Use this command to fetch
information and posts like wikipedia articles, sports scores, and even the weather.

Maybe one of these days he will be able to fetch a stick. xD.
`,
		Run: func(cmd *cobra.Command, args []string) {
			t := test
			if t == false {
				fmt.Println("Eddie is going to fetch.")
			} else {
				fmt.Println("You decided to test the fetch command.")
			}
		},
	}
)

func init() {
	fetchCmd.Flags().BoolVarP(&test, "test", "t", false, "Test the flag.")

	rootsCmd.AddCommand(fetchCmd)
}
