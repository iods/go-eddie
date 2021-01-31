package cli

import (
	"flag"
	"fmt"
	"os"
)

func New() {

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "This is a description for enabling something.")
	fooName := fooCmd.String("name", "", "This is a description for the name of something.")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "This is the description for the level of something.")


	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands.")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("enable:", *fooEnable)
		fmt.Println("name:", *fooName)
		fmt.Println("tail:", fooCmd.Args())
	case "bar":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("level:", *barLevel)
		fmt.Println("tail:", fooCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands.")
		os.Exit(1)
	}
}