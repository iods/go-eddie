package eddie

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)


func (e Eddie) ReportWeight() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to init termui: %v", err)
	}
	defer ui.Close()

	tw, th := ui.TerminalDimensions()
	grid := ui.NewGrid()
	grid.SetRect(0, 0, tw, th)
	grid.Set(
		ui.NewRow(0.5/4,
			ui.NewCol(2.0/3, titleParagraph()),
			ui.NewCol(1.0/3, clockParagraph()),
		),
		ui.NewRow(1.0/4,
			ui.NewCol(1.0/4, statOne()),
			ui.NewCol(1.0/4, statTwo()),
			ui.NewCol(1.0/4, statThree()),
			ui.NewCol(1.0/4, statFour()),
		),
		ui.NewRow(2.5/4,
			ui.NewCol(1/1, trendTable()),
		),
	)

	ui.Render(grid)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "e", "q", "<C-c>":
			return
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			grid.SetRect(0, 0, payload.Width, payload.Height)
			ui.Clear()
			ui.Render(grid)
		}
	}
}

func titleParagraph() *widgets.Paragraph {
	title := widgets.NewParagraph()
	title.Border = false
	title.Text = "eddie 🐕 the cli service dog. v0.1.0"
	title.TextStyle = ui.NewStyle(ui.ColorMagenta)
	title.PaddingLeft = 3

	return title
}

func clockParagraph() *widgets.Paragraph {
	clock := widgets.NewParagraph()
	clock.Border = false
	clock.Text = time.Now().Format("January 2, 2006 🕰️  15:04:05")
	clock.TextStyle = ui.NewStyle(ui.ColorWhite)
	clock.PaddingLeft = 3

	return clock
}

func trendTable() *widgets.Table {
	trends := widgets.NewTable()
	trends.Border = false
	trends.TextStyle = ui.NewStyle(ui.ColorWhite)

	trends.Rows = [][]string{
		{
			"", " 1 Week", " 30 Days", " 60 Days", " 90 Days", " 6 Months", " Total",
		},
		{"[Average](fg:magenta,mod:bold)", "188", "187", "189", "188", "190", "188"},
		{"[Maximum](fg:magenta,mod:bold)", "188", "[187](fg:clear,bg:red,mod:bold)", "189", "188", "190", "188"},
		{"[Minimum](fg:magenta,mod:bold)", "188", "187", "189", "188", "190", "188"},
		{"[Loss %](fg:cyan,mod:bold)", "0.0", "1.2", "0.9", "1.1", "1.4", "[1.3](fg:clear,bg:green,mod:bold)"},
		{"[Gain %](fg:cyan,mod:bold)", "0.0", "1.2", "0.9", "1.1", "1.4", "1.3"},
	}
	trends.RowSeparator = true
	trends.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	trends.PaddingTop = 1
	trends.PaddingRight = 1
	trends.PaddingLeft = 1

	return trends
}

func statOne() *widgets.Paragraph {
	text := widgets.NewParagraph()
	text.Text = "eddie calculated your average weight for this month is around [188](fg:green).\n"
	text.Border = false
	text.PaddingLeft = 3
	text.PaddingTop = 1

	return text
}

func statTwo() *widgets.Paragraph {
	text := widgets.NewParagraph()
	text.Border = false
	text.Text = "eddie tracked your weight [8x](fg:green) times this month [(+2)](fg:green).\n"
	text.PaddingLeft = 3
	text.PaddingTop = 1

	return text
}

func statThree() *widgets.Paragraph {
	text := widgets.NewParagraph()
	text.Text = "eddie calculated your body/mass index for this month at [123](fg:red).\n"
	text.Border = false
	text.TextStyle = ui.NewStyle(ui.ColorCyan)
	text.PaddingLeft = 3
	text.PaddingTop = 1

	return text
}

func statFour() *widgets.Paragraph {
	text := widgets.NewParagraph()
	text.Text = "eddie sees a [2%](fg:red)\nincrease in your\nweight in 60 days.\n"
	text.Border = false
	text.TextStyle = ui.NewStyle(ui.ColorCyan)
	text.PaddingLeft = 3
	text.PaddingTop = 1

	return text
}