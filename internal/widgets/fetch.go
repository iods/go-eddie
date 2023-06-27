package cmd

/*
Standups: A Go application for trying harder.
Copyright (c) 2022 Rye Miller
https://ryemiller.io
*/

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	eventCmd = &cobra.Command{
		Use:   "event",
		Short: "Review and set events on your calendar.",
		Long: `Add events, cancel events, view your calendar and
your teamm embers calendars, make reminders for notes and updates to
other team members. This is the money shot.

At least I am trying to make this happen. xD.
`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("event resource was called :) for standups.")
		},
	}
)

func init() {

}
