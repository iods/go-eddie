package cmd
/*
@TODO add symptom tracking (subcommand)
 eddie track symptom [allergies] \
	--tags="one,two,three,four" \
	--severity="3" \
	--important
*/
import (
	"fmt"
	"github.com/iods/go-eddie/internal/cli"
	"time"

	helper "github.com/iods/go-eddie/internal/helpers/time"
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


		t, _ := cmd.Flags().GetString("time")
		c, _ := cmd.Flags().GetInt("duration")
		q, _ := cmd.Flags().GetInt("quality")
		tags, _ := cmd.Flags().GetStringSlice("tags")
		l, _ := cmd.Flags().GetString("location")
		i, _ := cmd.Flags().GetBool("important")
		emojis, _ := cmd.Flags().GetStringSlice("emojis")

		switch args[0] {
		case "sleep":
			sleep(t, c, q, tags, l, false)
		case "weight":
			err := cli.TrackWeight(args[1], i)
			if err != nil {
				panic(err)
			}
		case "mood":
			err := cli.TrackMood(q, tags, emojis, i)
			if err != nil {
				panic(err)
			}
		case "seizure":
			seizure(t, tags, l, false)
		}


		if i != false {
			important()
		}
	},
}

func sleep(t string, d int, q int, tag []string, l string, f bool) {
	n := helper.UpdateTime(t)
	ts := n.Add(time.Duration(-d) * time.Hour)
	fmt.Println("You fell asleep at:", ts)
	fmt.Println("You woke up at:", n)
	fmt.Println("You slept for a total of", d, "hours")
	fmt.Println("You rated your sleep at a", q)
	fmt.Println("You slept on the", l)
	fmt.Println("You included the following tags:")
	for i := 0; i < len(tag); i++ {
		fmt.Println(tag[i])
	}
	if f != false {
		important()
	}
}

func seizure(t string, tag []string, loc string, f bool) {
	n := helper.UpdateTime(t)
	fmt.Println(n)
	for i := 0; i < len(tag); i++ {
		fmt.Println(tag[i])
	}
	fmt.Println("You reported a seizure at the", loc)
	if f != false {
		important()
	}
}

// important Tags records that are of importance in one way or another.
func important() {
	// --important, -i flag (not sure yet)
	fmt.Println("You rated this one as important.")
}

func init() {
	rootCmd.AddCommand(trackCmd)


	trackCmd.Flags().IntP("duration", "d", 0, "The time the event or activity lasted.")
	trackCmd.Flags().BoolP("important", "i", false, "If a record should be tagged important or not.")
	trackCmd.Flags().StringP("location", "l", "none", "Location of activity or event.")
	trackCmd.Flags().IntP("quality", "q", 0, "The quality of the time spent on the activity or event.")
	trackCmd.Flags().Int("total", 0, "The total amount of an input for an activity or event.")
	trackCmd.Flags().StringP("time", "t", "23:00", "The time the event or activity started.")

	trackCmd.Flags().StringSlice("emojis", []string{}, "Emojis to include in your year of pixels.")
	trackCmd.Flags().StringSlice("tags", []string{}, "Tags to include about the activity or event.")
}