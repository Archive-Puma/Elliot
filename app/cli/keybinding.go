package cli

import (
	"github.com/awesome-gocui/gocui"
	"github.com/sirupsen/logrus"

	"github.com/cosasdepuma/elliot/app/plugins"
)

type sKeybinding struct {
	view    string
	key     gocui.Key
	handler func(*gocui.Gui, *gocui.View) error
}

func (app *App) setKeybindings() error {
	app.keybindings = []sKeybinding{
		{"", gocui.KeyCtrlC, app.keybindingExit},
		{"", gocui.KeyTab, app.keybindingNextView},
		{"Target", gocui.KeyArrowUp, app.keybindingPreviousPlugin},
		{"Target", gocui.KeyArrowDown, app.keybindingNextPlugin},
		{"Target", gocui.KeyEnter, app.keybindingRun},
		{"Results", gocui.KeyEnter, app.keybindingRun},
	}

	for _, keybinding := range app.keybindings {
		if err := app.gui.SetKeybinding(keybinding.view, keybinding.key, gocui.ModNone, keybinding.handler); err != nil {
			return err
		}
	}
	for modal := range app.modalViews {
		if err := app.gui.SetKeybinding(modal, gocui.KeyEsc, gocui.ModNone, app.keybindingCancelModal); err != nil {
			return err
		}
		if err := app.gui.SetKeybinding(modal, gocui.KeyEnter, gocui.ModNone, app.keybindingCloseModal); err != nil {
			return err
		}
	}
	logrus.Debug("Keybindings configured")
	return nil
}

func (app *App) keybindingDisabled(gui *gocui.Gui, view *gocui.View) error {
	return nil
}

func (app *App) keybindingCancelModal(_ *gocui.Gui, _ *gocui.View) error {
	return app.closeModal()
}

func (app *App) keybindingCloseModal(_ *gocui.Gui, _ *gocui.View) error {
	app.sendStartSignal()
	return app.closeModal()
}

func (app *App) keybindingNextView(_ *gocui.Gui, _ *gocui.View) error {
	switch app.currentView {
	case 0:
		app.currentView = 2
	case 2:
		app.currentView = 0
	}
	return app.setFocus()
}

func (app *App) keybindingPreviousPlugin(_ *gocui.Gui, _ *gocui.View) error {
	view, err := app.gui.View("Plugins")
	if err != nil {
		return err
	}
	ox, oy := view.Origin()
	cx, cy := view.Cursor()
	if err := view.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := view.SetOrigin(ox, oy-1); err != nil {
			return err
		}
	}

	_, app.currentPlugin = view.Cursor()
	return nil
}

func (app *App) keybindingNextPlugin(_ *gocui.Gui, _ *gocui.View) error {
	view, err := app.gui.View("Plugins")
	if err != nil {
		return err
	}
	ox, oy := view.Origin()
	cx, cy := view.Cursor()
	if cy+1 < plugins.Amount {
		if err := view.SetCursor(cx, cy+1); err != nil {
			if err := view.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}

	_, app.currentPlugin = view.Cursor()
	return nil
}

func (app *App) keybindingRun(_ *gocui.Gui, _ *gocui.View) error {
	if err := app.getTarget(); err != nil {
		return err
	}
	if err := app.getPlugin(); err != nil {
		return err
	}
	if err := app.getParams(); err != nil {
		return err
	}
	app.runPlugin()
	return nil
}

func (app *App) keybindingExit(_ *gocui.Gui, _ *gocui.View) error {
	return gocui.ErrQuit
}
