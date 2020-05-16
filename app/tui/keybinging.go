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

func cursorUp(gui *gocui.Gui, view *gocui.View) error {
	if view != nil {
		ox, oy := view.Origin()
		cx, cy := view.Cursor()
		if err := view.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := view.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorDown(gui *gocui.Gui, view *gocui.View) error {
	if view != nil {
		ox, oy := view.Origin()
		cx, cy := view.Cursor()
		if cy+1 < NumberPlugins {
			if err := view.SetCursor(cx, cy+1); err != nil {
				if err := view.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func changeView(move int, gui *gocui.Gui, view *gocui.View) error {
	ActiveView = (ActiveView + move) % (len(Views) - 1)
	if _, err := setCurrentViewOnTop(gui, Views[ActiveView].name); err != nil {
		return err
	}

	gui.Cursor = (ActiveView == 0)

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
