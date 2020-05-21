package cli

import (
	"github.com/awesome-gocui/gocui"
	"github.com/sirupsen/logrus"
)

func (app *App) getFocusLayout() func(gui *gocui.Gui) error {
	return func(gui *gocui.Gui) error {
		view := gui.CurrentView()
		if err := app.onFocusChange(); err != nil {
			return err
		}
		if !app.isModal(view.Name()) {
			if err := app.onFocusLost(); err != nil {
				return err
			}
			if err := app.onFocus(view); err != nil {
				return err
			}
		}
		return nil
	}
}

func (app *App) setFocus() error {
	name := app.currentModal
	if app.currentView != -1 {
		name = app.mainViews[app.currentView].name
	}
	if _, err := app.gui.SetCurrentView(name); err != nil {
		return err
	}
	_, err := app.gui.SetViewOnTop(name)

	return err
}

func (app *App) onFocus(view *gocui.View) error {
	logrus.Debug("Focus gained: ", view.Name())
	return nil
}

func (app *App) onFocusLost() error {
	logrus.Debug("Focus lost")
	return nil
}

func (app *App) onFocusChange() error {
	current := app.gui.CurrentView()
	for _, view := range app.gui.Views() {
		view.Highlight = view == current
	}
	return nil
}
