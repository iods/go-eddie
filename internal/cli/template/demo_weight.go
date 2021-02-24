package template

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

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

	avg := widgets.NewBarChart()
	avg.Data = []float64{}
	avg.Labels = []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}
	avg.Border = false
	avg.BarWidth = 5
	avg.BarColors = []ui.Color{ui.ColorCyan, ui.ColorMagenta}
	avg.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}
	avg.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	text := widgets.NewParagraph()
	text.Text = "Eddie currently tracks and reports on your average weight totals over a specified period of time.\n\nIn the future, Eddie will be able to find patterns like worst and best months, median weight, and even connect trends with your sleep and mood."
	text.Border = false


	grid := ui.NewGrid()
	tw, th := ui.TerminalDimensions()
	grid.SetRect(0, 0, tw, th)
	grid.Set(
		ui.NewRow(0.5/5,
			ui.NewCol(2.0/3, title),
			ui.NewCol(1.0/3, clock),
		),
		ui.NewRow(4.5/5,
			ui.NewCol(0.75/2, text),
			ui.NewCol(1.25/2, avg),
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