package template

import (
	"fmt"
	"github.com/iods/go-eddie/internal/cli"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func ReportDashboard() {
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
	clock.Text = time.Now().Format("15:04:05 🕰️   January 2, 2006")
	clock.TextStyle = ui.NewStyle(ui.ColorWhite)
	clock.PaddingLeft = 3

	var yearData []float64
	var yearLabel []string

	for i := 1; i < 13; i++ {
		label, avg := cli.GetWeightAverageByMonth(i)
		yearData = append(yearData, avg)
		yearLabel = append(yearLabel, label)
	}

	year := widgets.NewBarChart()
	year.Border = false
	year.BarWidth = 5
	year.BarGap = 1
	year.BarColors = []ui.Color{ui.ColorBlue, ui.ColorCyan}
	year.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorCyan)}
	year.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
	year.Data = yearData
	year.Labels = yearLabel
	year.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}


	sbc := widgets.NewStackedBarChart()
	sbc.Border = false
	sbc.Data = make([][]float64, 5)
	sbc.Data[0] = []float64{6, 3}
	sbc.Data[1] = []float64{7, 2}
	sbc.Data[2] = []float64{1, 6}
	sbc.Data[3] = []float64{6, 2}
	sbc.Data[4] = []float64{1, 8}
	sbc.Labels = []string{"1m", "3m", "6m", "9m", "T"}
	sbc.BarWidth = 5
	sbc.BarColors = []ui.Color{ui.ColorMagenta, ui.ColorWhite}
	sbc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorCyan)}
	sbc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
	sbc.BarGap = 1
	sbc.PaddingLeft = 1


	text := widgets.NewParagraph()
	text.Text = "Eddie currently tracks and reports on your average weight totals over a specified period of time.\n"
	text.Border = false
	text.PaddingLeft = 1


	oneMonth, _ := cli.GetWeightAverage(1, false)
	threeMonths, _ := cli.GetWeightAverage(3, false)
	sixMonths, _ := cli.GetWeightAverage(6, false)
	nineMonths, _ := cli.GetWeightAverage(9, false)
	total, _ := cli.GetWeightAverageTotal()

	track := widgets.NewTable()
	track.Rows = [][]string{
		{"Timeframe", "Weight (-)", "Weight (+)", "Weight (avg)", "Weight (h)", "Weight (l)" },
		{"30 Days", "1.0", "1.0", fmt.Sprintf("%.1f", oneMonth),"184", "184"},
		{"90 Days", "0.0", "4.0", fmt.Sprintf("%.1f", threeMonths),"192", "185"},
		{"6 Months", "3.0", "0.0", fmt.Sprintf("%.1f", sixMonths),"188", "186"},
		{"9 Months", "3.0", "0.0", fmt.Sprintf("%.1f", nineMonths),"189", "181"},
		{"Total", "3.0", "2.0", fmt.Sprintf("%.1f", total),"187", "179"},
	}
	track.TextStyle = ui.NewStyle(ui.ColorWhite)
	track.RowSeparator = true
	track.Border = false

	track.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	track.RowStyles[2] = ui.NewStyle(ui.ColorClear, ui.ColorGreen, ui.ModifierBold)
	track.RowStyles[1] = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
	track.RowStyles[5] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)

	grid := ui.NewGrid()
	tw, th := ui.TerminalDimensions()
	grid.SetRect(0, 0, tw, th)
	grid.Set(
		ui.NewRow(0.5/5,
			ui.NewCol(2.0/3, title),
			ui.NewCol(1.0/3, clock),
		),
		ui.NewRow(2.25/5,
			ui.NewCol(0.75/2, text),
			ui.NewCol(1.25/2, year),
		),
		ui.NewRow(2.25/5,
			ui.NewCol(0.88/3, sbc),
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