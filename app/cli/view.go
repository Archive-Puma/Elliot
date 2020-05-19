package cli

// View TODO: Doc
type View struct {
	name       string
	x, y, w, h int
	middle     bool
	editable   bool
	list       bool
	frame      bool
}

var (
	// Current TODO: Doc
	Current = 0
	// MainViews TODO: Doc
	MainViews = []View{
		{name: "Target", x: 0, y: 0, w: -1, h: 2, frame: true, editable: true},
		{name: "Plugins", x: 0, y: 3, w: 18, h: -4, frame: true, list: true},
		{name: "Results", x: 19, y: 3, w: -1, h: -4, frame: true, editable: true},
		{name: "Messages", x: -1, y: -4, w: 0, h: -2},
		{name: "─", x: -1, y: -2, w: 0, h: 0, frame: true},
	}
	// ModalViews TODO: Doc
	ModalViews = map[string]struct {
		name  string
		input bool
	}{
		"Ports": {name: "Ports", input: true},
	}
)
