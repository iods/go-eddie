package template

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/iods/go-eddie/internal/db/schema"
)

type Render struct {
	Records []schema.Record
}

func RenderWeight() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	title := widgets.NewParagraph()
	title.Border = false
	title.Text = "eddie, a man's best friend. v0.1.0"
	title.TextStyle = ui.NewStyle(ui.ColorBlue)
	title.PaddingLeft = 3

	clock := widgets.NewParagraph()
	clock.Border = false
	clock.Text = time.Now().Format("15:04:05 🕰️ January 2, 2006")
	clock.TextStyle = ui.NewStyle(ui.ColorWhite)
	clock.PaddingLeft = 3

	year := widgets.NewBarChart()
	year.Border = false
	year.Title = "Yearly Average"
	year.Data = []float64{ 1, 2, 3, 4, 5, 1, 1, 3, 4, 5, 1, 1}
	year.Labels = []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}
	year.BarColors = []ui.Color{ui.ColorRed}
	year.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorCyan)}
	year.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	trends := widgets.NewBarChart()
	trends.Border = false
	trends.Title = "current trends"
	trends.BarWidth = 5
	trends.Data = []float64{ 182, 191, 183, 189, 187}
	trends.Labels = []string{"1m", "3m", "6m", "9m", "total"}
	trends.BarColors = []ui.Color{ui.ColorCyan, ui.ColorMagenta}
	trends.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorCyan)}
	trends.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	trends.PaddingLeft = 1

	text := widgets.NewParagraph()
	text.Text = "Eddie currently tracks and reports on your average weight totals over a specified period of time.\n\nIn the future, Eddie will be able to find patterns like worst and best months, median weight, and even connect trends with your sleep and mood."
	text.Border = false

	track := widgets.NewTable()
	track.Rows = [][]string{
		{"Timeframe", "Weight (-)", "Weight (+)", "Weight (avg)", "Weight (h)", "Weight (l)" },
		{"30 Days", "1.0", "1.0", "182","184", "184"},
		{"90 Days", "0.0", "4.0", "191","192", "185"},
		{"6 Months", "3.0", "0.0", "183","188", "186"},
		{"9 Months", "3.0", "0.0", "185","189", "181"},
		{"Total", "3.0", "2.0", "186","187", "179"},
	}
	track.TextStyle = ui.NewStyle(ui.ColorWhite)
	track.RowSeparator = true
	track.Border = false

	track.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	track.RowStyles[2] = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
	track.RowStyles[3] = ui.NewStyle(ui.ColorYellow)
	track.RowStyles[4] = ui.NewStyle(ui.ColorMagenta)

	grid := ui.NewGrid()
	tw, th := ui.TerminalDimensions()
	grid.SetRect(0, 0, tw, th)
	grid.Set(
		ui.NewRow(0.5/5,
			ui.NewCol(2.0/3, title),
			ui.NewCol(1.0/3, clock),
		),
		ui.NewRow(2.0/5,
			ui.NewCol(0.75/2, text),
			ui.NewCol(1.25/2, year),
		),
		ui.NewRow(2.5/5,
			ui.NewCol(0.88/3, trends),
			ui.NewCol(2.12/3, track),
		),
	)
	ui.Render(grid)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			grid.SetRect(0, 0, payload.Width, payload.Height)
			ui.Clear()
			ui.Render(grid)
		}
	}

}

// finish template for weight
// finish writing functions
// reformat current
// identify functions necessary
// write them






func getAverage(n []float64) float64 {
	total := 0.0

	for _, v := range n {
		total += v
	}

	return total / float64(len(n))
}