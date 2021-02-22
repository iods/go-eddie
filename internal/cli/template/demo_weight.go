package template

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/iods/go-eddie/internal/db"
	"github.com/iods/go-eddie/internal/db/schema"
	"log"
	"strconv"
)

func RenderWeight() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	var records []schema.Record

	db.InitDatabase()
	database := db.GetDatabase()

	r := database.Where("type LIKE ?", "%weight%").Find(&records)

	for _, value := range records {
		fmt.Println("Total weight was:", value.Total)
	}


	p := widgets.NewParagraph()
	p.Text = "You had " + strconv.FormatInt(r.RowsAffected, 10) + " entries."


	grid := ui.NewGrid()
	tw, th := ui.TerminalDimensions()
	grid.SetRect(0, 0, tw, th)
	grid.Set(
		ui.NewRow(1.0/3,
			ui.NewCol(1.0/3, p),
			ui.NewCol(2.0/3, p),
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