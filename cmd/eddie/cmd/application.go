package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var (
	persistRootFlag bool
	localRootFlag bool
	times int
	oenCmd = &cobra.Command{
		Use: "example",
		Short: "An example cobra program.",
		Long: `
This is an example of a long one.

Here you go.
`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from the root command!")
		},
	}
	echoCmd = &cobra.Command{
		Use: "echo [strings to echo]",
		Short: "prints to stdout",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: ", strings.Join(args, " "))
		},
	}
	timesCmd = &cobra.Command{
		Use: "times [strings to echo]",
		Short: "Prints a string to stdout a number of times.",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < times; i++ {
				fmt.Println("echo:", strings.Join(args, " "))
			}
		},
	}
)

// persistent vs. local - context
func init() {
	oenCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persist", "p", false, "A persistent.")
	oenCmd.Flags().BoolVarP(&localRootFlag, "local", "l", false, "A local flag.")
	timesCmd.Flags().IntVarP(&times, "times", "t", 1, "The times it takes.")
	oenCmd.AddCommand(echoCmd)
	echoCmd.AddCommand(timesCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
