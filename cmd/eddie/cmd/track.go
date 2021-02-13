/*

 eddie track sleep
  			 weight
   			 seizure
			 symptom
			 mood

 eddie track sleep \
	--start="0120" \
	--duration="9" \
	--quality="8" \
	--tags="one,two,three,four,five" \
	--location="couch" \
	--important

 eddie track weight 100

 eddie track seizure \
	--start="1020" \
	--tags="one,two,three,four,five" \
	--location="strip club" \
	--important

 eddie track symptom [allergies] \
	--tags="one,two,three,four" \
	--severity="3" \
	--important

 eddie track mood \
	--quality="7" \
	--tags="one,two,three,four,five" \
	--important

*/

package cmd

import (
	"fmt"
	"strings"
	"time"

	tm "github.com/iods/go-eddie/internal/helpers/time"
	"github.com/spf13/cobra"
)

var trackCmd = &cobra.Command{
	Use: "track {event,activity}",
	Short: "Record things like mood or sleep patterns.",
	Long: `
From sleep to mood to symptoms, eddie is like velcro.

Eddie will track your life, sniffing out everything about 
your health. so here are some ways you can use him to the fullest:
  1. eddie sniff [activity] {flags}
  2. eddie sniff [activity] -{emoji, tags, time, duration, stress, quality, location, severity, important}
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Println("eddie is tracking:", strings.Join(args, " "))

		t, _ := cmd.Flags().GetString("time")
		now := tm.Create(t)

		c, _ := cmd.Flags().GetInt("duration")
		then := now.Add(time.Duration(-c) * time.Hour)

		i, _ := cmd.Flags().GetBool("important")
		if i != false {
			fmt.Println("This is important!!")
		}
		fmt.Println(then)
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)

	// time flag (start,end,etc)
	trackCmd.Flags().StringP("time", "t", "10:00", "The time the event or activity started.")
	trackCmd.Flags().IntP("duration", "d", 0, "The time the event or activity lasted.")
	trackCmd.Flags().BoolP("important", "i", false, "If a record should be tagged important or not.")

}