package eddie

import (
	"fmt"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var weight Weight

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

	sevenAvg, _ := weight.FindAverages(7, false)
	thirtyAvg, _ := weight.FindAverages(1, false)
	sixtyAvg, _ := weight.FindAverages(2, false)
	ninetyAvg, _ := weight.FindAverages(3, false)
	sixAvg, _ := weight.FindAverages(6, false)
	total, _ := weight.FindAverages(0, false)

	average := []string{
		"[Average](fg:magenta,mod:bold)",
		fmt.Sprintf(" %0.2f", sevenAvg),
		fmt.Sprintf(" %0.2f", thirtyAvg),
		fmt.Sprintf(" %0.2f", sixtyAvg),
		fmt.Sprintf(" %0.2f", ninetyAvg),
		fmt.Sprintf(" %0.2f", sixAvg),
		fmt.Sprintf(" %0.2f", total),
	}

	sevenMin, sevenMax := weight.FindMinMax(7, false)
	thirtyMin, thirtyMax := weight.FindMinMax(1, false)
	sixtyMin, sixtyMax := weight.FindMinMax(2, false)
	ninetyMin, ninetyMax := weight.FindMinMax(3, false)
	sixMin, sixMax := weight.FindMinMax(6, false)
	totalMin, totalMax := weight.FindMinMax(0, false)

	maximum := []string{
		"[Maximum](fg:magenta,mod:bold)",
		fmt.Sprintf(" %0.2f", sevenMax.Total),
		fmt.Sprintf(" %0.2f", thirtyMax.Total),
		fmt.Sprintf(" %0.2f", sixtyMax.Total),
		fmt.Sprintf(" %0.2f", ninetyMax.Total),
		fmt.Sprintf(" %0.2f", sixMax.Total),
		fmt.Sprintf(" %0.2f", totalMax.Total),
	}

	minimum := []string{
		"[Minimum](fg:magenta,mod:bold)",
		fmt.Sprintf(" %0.2f", sevenMin.Total),
		fmt.Sprintf(" %0.2f", thirtyMin.Total),
		fmt.Sprintf(" %0.2f", sixtyMin.Total),
		fmt.Sprintf(" %0.2f", ninetyMin.Total),
		fmt.Sprintf(" %0.2f", sixMin.Total),
		fmt.Sprintf(" %0.2f", totalMin.Total),
	}

	trends := widgets.NewTable()
	trends.Border = false
	trends.TextStyle = ui.NewStyle(ui.ColorWhite)

	p7 := weight.FindPercentage(sevenMin.Total)
	pTotal := weight.FindPercentage(totalMin.Total)

	loss := []string{
		"[Loss %](fg:cyan,mod:bold)",
		fmt.Sprintf(" %0.1f", p7) + "%",
		fmt.Sprintf(" %0.1f", weight.FindPercentage(thirtyMin.Total)) + "%",
		fmt.Sprintf(" %0.1f", weight.FindPercentage(sixtyMin.Total)) + "%",
		fmt.Sprintf(" %0.1f", weight.FindPercentage(ninetyMin.Total)) + "%",
		fmt.Sprintf(" %0.1f", weight.FindPercentage(sixMin.Total)) + "%",
		fmt.Sprintf(" %0.1f", pTotal) + "%",
	}

	gain := []string{
		"[Gain %](fg:cyan,mod:bold)",
		fmt.Sprintf(" %0.1f", weight.FindPercentage(sevenMax.Total)) + "%",
		fmt.Sprintf(" %0.1f", weight.FindPercentage(thirtyMax.Total)) + "%",
		fmt.Sprintf(" %0.1f", weight.FindPercentage(sixtyMax.Total)) + "%",
		fmt.Sprintf(" %0.1f", weight.FindPercentage(ninetyMax.Total)) + "%",
		fmt.Sprintf(" %0.1f", weight.FindPercentage(sixMax.Total)) + "%",
		fmt.Sprintf(" %0.1f", weight.FindPercentage(totalMax.Total)) + "%",
	}

	trends.Rows = [][]string{
		{
			"", " 1 Week", " 30 Days", " 60 Days", " 90 Days", " 6 Months", " Total",
		},
		average, maximum, minimum, loss, gain,
	}
	trends.RowSeparator = true
	trends.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	trends.PaddingTop = 1
	trends.PaddingRight = 1
	trends.PaddingLeft = 1

	return trends
}

func statOne() *widgets.Paragraph {
	average, _ := weight.FindAverages(1, false)
	avg := fmt.Sprintf("%0.0f", average)

	a := "[" + avg + "](fg:red)"
	b := "[" + avg + "](fg:green)"

	text := widgets.NewParagraph()

	if average >= 188 {
		text.Text =  "eddie calculated your average weight over the past month around " + a
	} else {
		text.Text =  "eddie calculated your average weight over the past month around " + b
	}
	text.Border = false
	text.TextStyle = ui.NewStyle(ui.ColorCyan)
	text.PaddingLeft = 3
	text.PaddingTop = 1

	return text
}

func statTwo() *widgets.Paragraph {
	present := weight.FindCount()
	count := fmt.Sprintf("%d", present)

	text := widgets.NewParagraph()
	text.Border = false
	text.Text = "eddie tracked your weight a total of [" + count + "](fg:green) times this month"
	text.PaddingLeft = 3
	text.PaddingTop = 1

	return text
}

func statThree() *widgets.Paragraph {

	total, _ := weight.FindAverages(0, false)
	bmiTotal, _ := weight.FindBmi(total, 80)
	bmi := fmt.Sprintf("%0.0f", bmiTotal)

	text := widgets.NewParagraph()

	if bmiTotal > 25 {
		text.Text = "eddie calculated your BMI at [" + bmi + "](fg:red), which is considered [obese](fg:red)"
	} else if bmiTotal < 25 && bmiTotal > 18.5 {
		text.Text = "eddie calculated your BMI at [" + bmi + "](fg:green), which is considered [optimal](fg:green)"
	} else {
		text.Text = "eddie calculated your BMI at [" + bmi + "](fg:green), which is considered [underweight](fg:green)"
	}
	text.Border = false
	text.PaddingLeft = 3
	text.PaddingTop = 1

	return text
}

func statFour() *widgets.Paragraph {

	total, _ := weight.FindAverages(3, false)
	if total == 0 {

	}

	percentage := weight.FindPercentage(total)
	percent := fmt.Sprintf("%0.1f", percentage)


	text := widgets.NewParagraph()

	if total > 188 {
		text.Text = "eddie noticed a [" + percent + "%](fg:green)\ndecrease in your\nweight over 90 days"
	} else {
		text.Text = "eddie noticed a [" + percent + "%](fg:red)\nincrease in your\nweight over 90 days"
	}
	text.Border = false
	text.PaddingLeft = 3
	text.PaddingTop = 1

	return text
}