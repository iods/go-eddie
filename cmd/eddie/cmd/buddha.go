package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	m "github.com/iods/go-eddie/internal/model"
	"github.com/spf13/cobra"
)

var buddhaCmd = &cobra.Command{
	Use:   "buddha",
	Short: "Returns some words of wisdom from Buddha.",
	Long: `
Eddie returns some advice from Buddha.

Eddie randomizes fortune cookie messages, selects one based 
on some predefined filters, and if you are feeling lucky, eddie
returns some numbers you may want to play the lottery with!
`,
	Run: func(cmd *cobra.Command, args []string) {
		test := fortune()
		fmt.Println(test.Fortunes[0].Id)
		fmt.Println(test.Fortunes[0].Text)
		fmt.Println(test.Fortunes[0].Numbers)
	},
}

// count how many fortunes exist

// have eddie take 







func fortune() m.Fortunes {
	var fortunes m.Fortunes
	filepath := "../../data/fortunes.json"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &fortunes)
	if err != nil {
		log.Fatal(err)
	}
	return fortunes
}

// init Main function for Cobra to build the command.
func init() {
	askCmd.AddCommand(buddhaCmd)
}
