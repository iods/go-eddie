package main

/*

eddie [global options...] {command} [--option] [argument]

eddie ask		// an automated notifier, then template to record some notes
eddie report	// a data manager, compiling data in the right place for API's to access it
eddie track		// the "order entry" of the application if you will, takes in data about a single activity (routine/behavior)


eddie ask buddha (returns a fortune cookie saying)
eddie ask buddha -n (returns the #'s also, randomizes them if none)

eddie ask [activity] -{d,w,m}
eddie ask [symptom]

eddie report [activity]
eddie report [symptom]

eddie track sleep \
	-tags="this,that,then,those,that" \
	-start=0100 \
	-end=0100 \
	-quality=3 (measuring tags, quality, duration)

eddie track weight 100 (average weight)
eddie track mood 7 -tags="this,that,then,bag,of,chips" -stress=6 -quality=3 -emoji="joyful,happy,grateful,sad,loving,anxious,bored,frustrated,stressed)

eddie track symptom \
	-allergies=3 \
	-back-pain=2 \
	-chest-pain=2 \

eddie track seizure -tags="place,why,weird" -time=10202 -location="place important?"


*/
import "github.com/iods/go-eddie/cmd/eddie/cmd"

func main() {
	cmd.Execute()
}
