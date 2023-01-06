package eddie

type Dashboard struct {
	Type  string
	Title string
}

type DashboardRender func()