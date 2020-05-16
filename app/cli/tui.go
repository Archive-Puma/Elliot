package cli

import (
	"github.com/jroimartin/gocui"

	mrerr "github.com/cosasdepuma/elliot/app/error"
)

// ShowUI TODO: Doc
func ShowUI() *mrerr.MrRobotError {
	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return mrerr.NewCritical("Terminal UI cannot be created")
	}
	defer gui.Close()

	gui.SetManagerFunc(layoutManager)

	if err := setKeybindings(gui); err != nil {
		return mrerr.NewCritical("Keybindings cannot be set")
	}

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		return mrerr.NewCritical("Cannot run Terminal UI")
	}

	return nil
}
