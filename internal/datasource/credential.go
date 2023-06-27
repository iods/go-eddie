package fetch

/*
Standups: A Go application for trying harder.
Copyright (c) 2022 Rye Miller
https://ryemiller.io
*/

import (
	"fmt"

	"github.com/spf13/cobra"
)

var add bool

var (
	credentialCmd = &cobra.Command{
		Use:   "fetch",
		Short: "Store and encrypt passwords for client projects.",
		Long: `Add a password, and the project it is for, you can also have
it encrypted and stored so that you can paste it right into the field online.

At least I am trying to make this happen. xD.
`,
		Run: func(cmd *cobra.Command, args []string) {
			a := add
			if a == false {
				fmt.Println("credential resource was called :) for standups.")
			} else {
				fmt.Println("you would like to add a credential, huh?")
			}
		},
	}
)

func init() {
	credentialCmd.Flags().BoolVarP(&add, "add", "a", false, "Add a credential.")

	rootCmd.AddCommand(credentialCmd)
}
