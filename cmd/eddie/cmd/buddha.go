package cmd

import (
	"encoding/json"
	"fmt"
	m "github.com/iods/go-eddie/internal/model"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
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

// getCount Returns how many objects (fortunes) exist in the struct.
func getCount() int {
	f := fortune()
	return len(f.Fortunes)
}

// have eddie take the returned integer and use in range

// use range function to calculate number between 1 and what eddie returns as count of fortunes

// use rand to generate random number from that returned integer

// set variable for mapped dataset to pull in returned integer and print message





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
