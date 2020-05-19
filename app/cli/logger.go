package cli

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

// Logger TODO: Doc
var Logger = struct {
	Type string
	Msg  string
}{
	Type: "Info",
	Msg:  "",
}

func displayLogger(gui *gocui.Gui) error {
	loggerView, err := gui.View(MainViews[len(MainViews)-2].name)
	if err != nil {
		return err
	}
	loggerView.Clear()
	switch Logger.Type {
	case "Error":
		loggerView.FgColor = gocui.ColorRed
	case "Info":
		loggerView.FgColor = gocui.ColorDefault
	}
	fmt.Fprintf(loggerView, "%s: %s", Logger.Type, Logger.Msg)
	return nil
}
