package template

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
	"gorm.io/gorm"
	"log"
	"time"
)

func RenderMood() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	var records schema.Record

	db.InitDatabase()
	database := db.GetDatabase()

	database.Where("type LIKE ?", "%mood%").Find(&records)

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


	grid := ui.NewGrid()
	tw, th := ui.TerminalDimensions()
	grid.SetRect(0, 0, tw, th)
	grid.Set(
		ui.NewRow(0.5/5,
			ui.NewCol(2.0/3, title),
			ui.NewCol(1.0/3, clock),
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

func getTags() *gorm.DB {
	var tags []schema.Record

	db.InitDatabase()
	database := db.GetDatabase()

	return database.Select("record_tags").Find(&tags)
}