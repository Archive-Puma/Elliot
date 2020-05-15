package tui

import (
	"github.com/jroimartin/gocui"

	mrerr "github.com/cosasdepuma/elliot/app/error"
)

// ShowTUI TODO: Doc
func ShowTUI() *mrerr.MrRobotError {
	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return mrerr.NewCritical("Terminal UI cannot be created")
	}
	defer gui.Close()

	gui.Cursor = true
	gui.Highlight = true
	gui.SelFgColor = gocui.ColorYellow
	gui.SetManagerFunc(mainLayout)

	if err := setKeybindings(gui); err != nil {
		return mrerr.NewCritical("Keybindings cannot be set")
	}

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		return mrerr.NewCritical("Cannot run Terminal UI")
	}

	return nil
}
