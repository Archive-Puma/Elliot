package tui

type sView struct {
	name       string
	x, y, w, h int
	editable   bool
	frame      bool
}

var (
	active = 0
	views  = []sView{
		{name: "Target", x: 0, y: 0, w: -1, h: 2, frame: true, editable: true},
		{name: "Subcommands", x: 0, y: 3, w: 23, h: -3, frame: true},
		{name: "Results", x: 24, y: 3, w: -1, h: -3, frame: true},
		{name: "─", x: -1, y: -2, w: 0, h: 0, frame: true},
	}
)
