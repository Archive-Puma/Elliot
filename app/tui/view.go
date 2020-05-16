package tui

type sView struct {
	name       string
	x, y, w, h int
	editable   bool
	list       bool
	frame      bool
}

var (
	// ActiveView TODO: Doc
	ActiveView = 0
	// NumberPlugins TODO: Doc
	NumberPlugins = 0
	// Views TODO: Doc
	Views = [...]sView{
		{name: "Target", x: 0, y: 0, w: -1, h: 2, frame: true, editable: true},
		{name: "Plugins", x: 0, y: 3, w: 18, h: -3, frame: true, list: true},
		{name: "Results", x: 19, y: 3, w: -1, h: -3, frame: true},
		{name: "─", x: -1, y: -2, w: 0, h: 0, frame: true},
	}
)
