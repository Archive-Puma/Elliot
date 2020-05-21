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
		{"Target", gocui.KeyArrowUp, app.keybindingDisabled},
		{"Target", gocui.KeyArrowDown, app.keybindingDisabled},
		{"Plugins", gocui.KeyArrowUp, app.keybindingPreviousPlugin},
		{"Plugins", gocui.KeyArrowDown, app.keybindingNextPlugin},
	}

	for _, keybinding := range app.keybindings {
		if err := app.gui.SetKeybinding(keybinding.view, keybinding.key, gocui.ModNone, keybinding.handler); err != nil {
			return err
		}
	}
	for modal := range app.modalViews {
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

func (app *App) keybindingCloseModal(gui *gocui.Gui, view *gocui.View) error {
	return app.closeModal()
}

func (app *App) keybindingNextView(gui *gocui.Gui, view *gocui.View) error {
	if app.currentView != -1 {
		app.currentView = (app.currentView + 1) % (len(app.mainViews) - 2)
	}
	return app.setFocus()
}

func (app *App) keybindingPreviousPlugin(gui *gocui.Gui, view *gocui.View) error {
	if view != nil {
		ox, oy := view.Origin()
		cx, cy := view.Cursor()
		if err := view.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := view.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	_, app.currentPlugin = view.Cursor()
	return nil
}

func (app *App) keybindingNextPlugin(gui *gocui.Gui, view *gocui.View) error {
	if view != nil {
		ox, oy := view.Origin()
		cx, cy := view.Cursor()
		if cy+1 < plugins.Amount {
			if err := view.SetCursor(cx, cy+1); err != nil {
				if err := view.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}
		}
	}
	_, app.currentPlugin = view.Cursor()
	return nil
}

func (app *App) keybindingExit(gui *gocui.Gui, view *gocui.View) error {
	return gocui.ErrQuit
}
