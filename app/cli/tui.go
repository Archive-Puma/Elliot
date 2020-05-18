package cli

import (
	"github.com/awesome-gocui/gocui"
)

// ShowUI TODO: Doc
func ShowUI() error {
	gui, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		return err
	}
	defer gui.Close()

	gui.SetManagerFunc(mainLayout)

	if err := setKeybindings(gui); err != nil {
		return err
	}

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}
