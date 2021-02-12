package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

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
		m := getMessage()
		flag, _ := cmd.Flags().GetBool("lucky")
		if flag {
			fmt.Printf("%s %v\n", m.Text, m.Numbers)
		} else {
			fmt.Println(m.Text)
		}
	},
}

// getCount Returns how many objects (fortunes) exist in the struct.
func getCount() int {
	f := fortune()
	return len(f.Fortunes)
}

// getMessage Returns the string of a random fortune object.
func getMessage() m.Fortune {
	v := getCount()
	r := getRandom(v)
	f := fortune().Fortunes[r]
	return f
}

// returnRandom Returns a random number between one and the function's argument.
func getRandom(value int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	random := r1.Intn(value)
	return random
}

// fortune Reads the json file and returns a dataset of fortunes.
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

// init Main function for Cobra to build the buddha command.
func init() {
	askCmd.AddCommand(buddhaCmd)
	buddhaCmd.Flags().BoolP("lucky", "l", false, "Have eddie include some lucky numbers with your message.")
}
