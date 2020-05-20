package gui

import (
	"errors"
	"strings"

	"github.com/sirupsen/logrus"
)

func (app *App) showModal(name string) error {
	logrus.Debug("Modal shown")
	modal, ok := app.modalViews[name]
	if !ok {
		return errors.New("Modal not found")
	}
	app.Params = nil
	app.currentModal = name
	if app.currentView != -1 {
		app.lastView = app.currentView
	}
	app.currentView = -1
	_, err := app.gui.SetView(name, app.dimensions.width/2+modal.coords.x, app.dimensions.height/2+modal.coords.y, app.dimensions.width/2+(-1*modal.coords.w+2), app.dimensions.height/2+modal.coords.h, 0)
	return err
}

func (app *App) closeModal() error {
	logrus.Debug("Modal closed")
	logrus.Warning(app.lastView, app.currentView)
	modal, ok := app.modalViews[app.currentModal]
	if !ok {
		return errors.New("Modal not found")
	}
	view, err := app.gui.View(app.currentModal)
	if err != nil {
		return err
	}
	app.currentView = app.lastView
	app.Params = strings.TrimSpace(view.Buffer())
	_, err = app.gui.SetView(app.currentModal, modal.coords.x, modal.coords.y, modal.coords.w, modal.coords.h, 0)
	app.currentModal = ""
	return err
}

func (app *App) isModal(name string) bool {
	_, ok := app.modalViews[name]
	return ok
}
