package template

//import (
//	"fmt"
//	"github.com/iods/go-eddie/internal/cli"
//	"log"
//	"time"
//
//	ui "github.com/gizak/termui/v3"
//	"github.com/gizak/termui/v3/widgets"
//)
//
//func ReportDashboard() {
//	if err := ui.Init(); err != nil {
//		log.Fatalf("failed to initialize termui: %v", err)
//	}
//	defer ui.Close()
//
//	title := widgets.NewParagraph()
//	title.Border = false
//	title.Text = `eddie 🐕 the cli service dog. v0.1.0`
//	title.TextStyle = ui.NewStyle(ui.ColorBlue)
//	title.PaddingLeft = 3
//
//
//	clock := widgets.NewParagraph()
//	clock.Border = false
//	clock.Text = time.Now().Format("January 2, 2006 🕰️  15:04:05")
//	clock.TextStyle = ui.NewStyle(ui.ColorWhite)
//	clock.PaddingLeft = 3
//
//	var yearData []float64
//	for i := 1; i < 13; i++ {
//		avg, _, _ := cli.GetWeightAverageByMonth(i)
//		yearData = append(yearData, avg)
//	}
//
//	year := widgets.NewBarChart()
//	year.Border = false
//	year.NumFormatter = func(a float64) string { return fmt.Sprintf("%.0f", a) }
//	year.BarWidth = 5
//	year.BarGap = 1
//	year.BarColors = []ui.Color{ui.ColorBlue, ui.ColorCyan}
//	year.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorCyan)}
//	year.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
//	year.Data = yearData
//	year.Labels = []string{"1m", "2m", "3m", "4m", "5m", "6m", "7m", "8m", "9m", "10m", "11m", "12m"}
//
//	sbc := widgets.NewStackedBarChart()
//	sbc.Border = false
//	sbc.Data = make([][]float64, 5)
//	sbc.Data[0] = []float64{6, 3}
//	sbc.Data[1] = []float64{7, 2}
//	sbc.Data[2] = []float64{1, 6}
//	sbc.Data[3] = []float64{6, 2}
//	sbc.Data[4] = []float64{1, 8}
//	sbc.Labels = []string{"1m", "3m", "6m", "9m", "T"}
//	sbc.BarWidth = 5
//	sbc.BarColors = []ui.Color{ui.ColorMagenta, ui.ColorWhite}
//	sbc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorCyan)}
//	sbc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
//	sbc.BarGap = 1
//	sbc.PaddingLeft = 1
//
//
//	text := widgets.NewParagraph()
//	text.Text = "Eddie currently tracks and reports on your average weight totals over a specified period of time.\n"
//	text.Border = false
//	text.PaddingLeft = 1
//
//
//	oneMonth, oneH, oneL := cli.GetWeightAverageByMonth(1)
//
//	// starting weight is first record (first "created_at")
//
//
//	track := widgets.NewTable()
//	track.Rows = [][]string{
//		{"Timeframe", "", "Weight (+)", "Average", "Maximum", "Minimum" },
//		{
//			"30 Days",
//			"1.0",
//			"1.0",
//			fmt.Sprintf("%0.1f", oneMonth),
//			fmt.Sprintf("%0.1f", oneH),
//			fmt.Sprintf("%0.1f", oneL),
//		},
//		{"90 Days", "0.0", "4.0", fmt.Sprintf("%0.1f", threeMonth), fmt.Sprintf("%0.1f", threeH), fmt.Sprintf("%0.1f", threeL)},
//		{"6 Months", "3.0", "0.0", fmt.Sprintf("%0.1f", sixMonths), fmt.Sprintf("%0.1f", sixH), fmt.Sprintf("%0.1f", sixL)},
//		{"9 Months", "3.0", "0.0", fmt.Sprintf("%0.1f", nineMonths), fmt.Sprintf("%0.1f", nineF), fmt.Sprintf("%0.1f", nineL)},
//		{"Total", l, g, fmt.Sprintf("%0.1f", total), m1, m2},
//	}
//	track.TextStyle = ui.NewStyle(ui.ColorWhite)
//	track.RowSeparator = true
//	track.Border = false
//
//	track.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
//	track.RowStyles[2] = ui.NewStyle(ui.ColorClear, ui.ColorGreen, ui.ModifierBold)
//	track.RowStyles[1] = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
//	track.RowStyles[5] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
//
//	grid := ui.NewGrid()
//	tw, th := ui.TerminalDimensions()
//	grid.SetRect(0, 0, tw, th)
//	grid.Set(
//		ui.NewRow(0.5/5,
//			ui.NewCol(2.0/3, title),
//			ui.NewCol(1.0/3, clock),
//		),
//		ui.NewRow(2.25/5,
//			ui.NewCol(0.75/2, text),
//			ui.NewCol(1.25/2, year),
//		),
//		ui.NewRow(2.25/5,
//			ui.NewCol(0.88/3, sbc),
//			ui.NewCol(2.12/3, track),
//		),
//	)
//	ui.Render(grid)
//
//	uiEvents := ui.PollEvents()
//	for {
//		e := <-uiEvents
//		switch e.ID {
//		case "q", "<C-c>":
//			return
//		case "<Resize>":
//			payload := e.Payload.(ui.Resize)
//			grid.SetRect(0, 0, payload.Width, payload.Height)
//			ui.Clear()
//			ui.Render(grid)
//		}
//	}
//}