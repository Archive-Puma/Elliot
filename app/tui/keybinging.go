package tui

import (
	"github.com/jroimartin/gocui"

	mrerr "github.com/cosasdepuma/elliot/app/error"
)

func exitApplication(_ *gocui.Gui, _ *gocui.View) error {
	return gocui.ErrQuit
}

func disable(_ *gocui.Gui, view *gocui.View) error {
	return nil
}

func changeView(move int, gui *gocui.Gui, view *gocui.View) error {
	active = (active + move) % (len(views) - 1)
	if _, err := setCurrentViewOnTop(gui, views[active].name); err != nil {
		return err
	}

	gui.Cursor = (active == 0)

	return nil
}

func nextView(gui *gocui.Gui, view *gocui.View) error {
	return changeView(1, gui, view)
}

func setCurrentViewOnTop(gui *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := gui.SetCurrentView(name); err != nil {
		return nil, err
	}
	return gui.SetViewOnTop(name)
}

func setKeybindings(gui *gocui.Gui) *mrerr.MrRobotError {
	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, exitApplication); err != nil {
		return mrerr.NewCritical("Cannot set Ctrl+C keybinding")
	}
	if err := gui.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return mrerr.NewCritical("Cannot set TAB keybinding")
	}

	return nil
}
