package template

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"math"
	"time"
)

func RenderDemo() {

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	title := widgets.NewParagraph()
	title.Border = false
	title.Text = "eddie, a man's best friend. v0.1.0"
	title.TextStyle = ui.NewStyle(ui.ColorBlue)
	title.PaddingTop = 3
	title.PaddingLeft = 1

	clock := widgets.NewParagraph()
	clock.Border = false
	clock.Text = time.Now().Format("15:04:05 🕰️ January 2, 2006")
	clock.TextStyle = ui.NewStyle(ui.ColorWhite)
	clock.PaddingLeft = 1

	bc := widgets.NewBarChart()
	bc.Data = []float64{3, 2, 5, 3, 9, 3}
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.Title = "Bar Chart"
	bc.SetRect(5, 5, 100, 25)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	data := []float64{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}

	sl0 := widgets.NewSparkline()
	sl0.Data = data[3:]
	sl0.LineColor = ui.ColorGreen

	// single
	slg0 := widgets.NewSparklineGroup(sl0)
	slg0.Title = "Sparkline 0"
	slg0.SetRect(0, 0, 20, 10)




	sinData := func() [][]float64 {
		n := 220
		data := make([][]float64, 2)
		data[0] = make([]float64, n)
		data[1] = make([]float64, n)
		for i := 0; i < n; i++ {
			data[0][i] = 1 + math.Sin(float64(i)/5)
			data[1][i] = 1 + math.Cos(float64(i)/5)
		}
		return data
	}()

	p0 := widgets.NewPlot()
	p0.Title = "braille-mode Line Chart"
	p0.Data = sinData
	p0.SetRect(0, 0, 50, 15)
	p0.AxesColor = ui.ColorWhite
	p0.LineColors[0] = ui.ColorGreen

	table3 := widgets.NewTable()
	table3.Rows = [][]string{
		[]string{"header1", "header2", "header3"},
		[]string{"AAA", "BBB", "CCC"},
		[]string{"DDD", "EEE", "FFF"},
		[]string{"GGG", "HHH", "III"},
	}
	table3.TextStyle = ui.NewStyle(ui.ColorWhite)
	table3.RowSeparator = true
	table3.BorderStyle = ui.NewStyle(ui.ColorGreen)
	table3.SetRect(0, 30, 70, 20)
	table3.FillRow = true
	table3.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	table3.RowStyles[2] = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
	table3.RowStyles[3] = ui.NewStyle(ui.ColorYellow)

	pc := widgets.NewPieChart()
	pc.Title = "Pie Chart"
	pc.SetRect(5, 5, 70, 36)
	pc.Data = []float64{.25, .25, .25, .25}
	pc.AngleOffset = -.5 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.02f", v)
	}

	c := widgets.NewParagraph()
	c.Border = false

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/3,
			ui.NewCol(1.0/3, pc),
			ui.NewCol(1.0/3, bc),
			ui.NewCol(1.0/3, bc),
		),
		ui.NewRow(1.0/3,
			ui.NewCol(1.0/3, slg0),
			ui.NewCol(2.0/3, p0),
		),
		ui.NewRow(1.0/3,
			ui.NewCol(1.0/3, bc),
			ui.NewCol(1.0/3, table3),
			ui.NewCol(1.0/3,
				ui.NewRow(1.0/2, c),
				ui.NewRow(1.0/2, slg0),
			),
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