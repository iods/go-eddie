package cmd
/*
@TODO add symptom tracking (subcommand)
 eddie track symptom [allergies] \
	--tags="one,two,three,four" \
	--severity="3" \
	--important
*/
import (
	"github.com/iods/go-eddie/internal/cli"
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

		t, _ := cmd.Flags().GetString("time") 	// time is translated into To, From, and Time depending on cmd
		i, _ := cmd.Flags().GetBool("important")	// bool for yes/no importance
		d, _ := cmd.Flags().GetInt("duration")	// used with length (length may be used by different scopes in future)
		q, _ := cmd.Flags().GetInt("quality")		// used as a rating system, most often w/ duration
		l, _ := cmd.Flags().GetString("location") // can be used with many, but is specific to seizure tracking (reports)

		tags, _ := cmd.Flags().GetStringSlice("tags")
		emojis, _ := cmd.Flags().GetStringSlice("emojis")

		switch args[0] {
		case "sleep":
			err := cli.TrackSleep(t, d, q, tags, i)
			if err != nil {
				panic(err)
			}
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
			err := cli.TrackSeizure(t, tags, l, i)
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)

	trackCmd.Flags().StringP("time", "t", "23:00", "The time the event or activity started.")
	trackCmd.Flags().BoolP("important", "i", false, "If a record should be tagged important or not.")
	trackCmd.Flags().IntP("duration", "d", 0, "The time the event or activity lasted.")
	trackCmd.Flags().IntP("quality", "q", 0, "The quality of the time spent on the activity or event.")
	trackCmd.Flags().StringP("location", "l", "none", "Location of activity or event.")
	trackCmd.Flags().StringSlice("emojis", []string{}, "Emojis to include in your year of pixels.")
	trackCmd.Flags().StringSlice("tags", []string{}, "Tags to include about the activity or event.")
}