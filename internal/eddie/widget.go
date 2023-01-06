package eddie

import ui "github.com/gizak/termui/v3"

type Widget struct {
	Title string
}

type WidgetAction func()

type WidgetEvent func(e ui.Event)

func WidgetActiveStyle() ui.Style {
	return ui.NewStyle(ui.ColorYellow, ui.ColorClear, ui.ModifierBold)
}

func WidgetDefaultStyle() ui.Style {
	return ui.NewStyle(ui.ColorWhite)
}

func WidgetBorderStyle() ui.Style {
	return ui.NewStyle(ui.ColorWhite)
}

func WidgetTitleStyle() ui.Style {
	return ui.NewStyle(ui.ColorBlack, ui.ColorWhite, ui.ModifierBold)
}