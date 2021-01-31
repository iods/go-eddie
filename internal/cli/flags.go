package cli

import (
	"flag"
	"fmt"
	"os"
)

var (
	boolFlag bool
	count	 []string
	Flagname int
	help	 bool
	Ip		 int
	Strflag	 string
	Wordptr	 string
)


func Flags() {

	flag.StringVar(&Strflag, "t", "", "This is a description of the flag. (Required.)")
	flag.BoolVar(&boolFlag, "b", false, "This is another description of the flag.")
	flag.IntVar(&Flagname, "flagname", 12345, "Another help message for this.")
	flag.IntVar(&Ip, "ip", 12700001, "The help message for this god damn flag.")

	flag.StringVar(&Wordptr, "words", "fooooooo", "Another word in the mix.")

	flag.BoolVar(&help, "help", false, "Show this help message.")
	flag.BoolVar(&help, "h", false, "Show this help message.")

	var svar string

	wordPtr := flag.String("word", "foo", "A string.")
	numberPtr := flag.Int("number", 123, "A number, integer.")
	boolPtr := flag.Bool("abool", false, "A boolean.")
	flag.StringVar(&svar, "svar", "bar", "A string variable.")

	flag.Parse()

	fmt.Println("-t: ", Strflag)
	fmt.Println("-b: ", boolFlag)
	fmt.Println("-flagname: ", Flagname)
	fmt.Println("word: ", *wordPtr)
	fmt.Println("number: ", *numberPtr)
	fmt.Println("Bool: ", *boolPtr)
	fmt.Println("String Var: ", svar)


	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	count = os.Args
	if len(count) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
}