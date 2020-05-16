package cli

var (
	// Current TODO: Doc
	Current = 0
	// Popup TODO: Doc
	Popup = struct {
		Active bool
		Title  string
		Msg    string
	}{
		Active: false,
		Title:  "Error",
		Msg:    "Cosas guays que suelen quedar bien",
	}
	// Views TODO: Doc
	Views = [...]struct {
		name       string
		x, y, w, h int
		editable   bool
		list       bool
		frame      bool
	}{
		{name: "Target", x: 0, y: 0, w: -1, h: 2, frame: true, editable: true},
		{name: "Plugins", x: 0, y: 3, w: 18, h: -3, frame: true, list: true},
		{name: "Results", x: 19, y: 3, w: -1, h: -3, frame: true},
		{name: "─", x: -1, y: -2, w: 0, h: 0, frame: true},
	}
)
