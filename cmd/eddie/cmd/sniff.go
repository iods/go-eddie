/*

 eddie sniff sleep
  			 weight
   			 seizure
			 symptom
			 mood

 eddie sniff sleep \
	--start="0120" \
	--duration="9" \
	--quality="8" \
	--tags="one,two,three,four,five" \
	--location="couch" \
	--important

 eddie sniff weight 100

 eddie sniff seizure \
	--start="1020" \
	--tags="one,two,three,four,five" \
	--location="strip club" \
	--important

 eddie sniff symptom [allergies] \
	--tags="one,two,three,four" \
	--severity="3" \
	--important

 eddie sniff mood \
	--quality="7" \
	--tags="one,two,three,four,five" \
	--important

*/


package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var sniffCmd = &cobra.Command{
	Use: "sniff [event or behavior]",
	Short: "Have eddie track an event or behavior.",
	Long: `
From sleep to mood to symptoms, eddie is like velcro.

Eddie will track your life, sniffing out everything about 
your health. so here are some ways you can use him to the fullest:
  1. eddie sniff [activity] {flags}
  2. eddie sniff [activity] -{emoji, tags, time, duration, stress, quality, location, severity, important}
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Eddie is gonna track something!: ", strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(sniffCmd)
}
