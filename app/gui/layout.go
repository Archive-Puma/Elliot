package gui

import (
	"fmt"
	"sort"

	"github.com/awesome-gocui/gocui"
	"github.com/cosasdepuma/elliot/app/plugins"
	"github.com/sirupsen/logrus"
)

type coordinates struct {
	x, y, w, h int
}

type sView struct {
	name     string
	coords   coordinates
	editable bool
	list     bool
	frame    bool
}

func (app *App) calculatePosition(indexMainView int) coordinates {
	view := app.mainViews[indexMainView]
	coords := coordinates{}
	// X
	coords.x = view.coords.x
	if coords.x < -1 {
		coords.x = app.dimensions.width + coords.x
	}
	// Y
	coords.y = view.coords.y
	if coords.y < -1 {
		coords.y = app.dimensions.height + coords.y
	}
	// W
	coords.w = view.coords.w
	if coords.w <= 0 {
		coords.w = app.dimensions.width + coords.w
	}
	// H
	coords.h = view.coords.h
	if coords.h <= 0 {
		coords.h = app.dimensions.height + coords.h
	}
	return coords
}

func (app *App) drawMainViews() error {
	for index, view := range app.mainViews {
		position := app.calculatePosition(index)
		if panel, err := app.gui.SetView(view.name, position.x, position.y, position.w, position.h, 0); err != nil {
			if !gocui.IsUnknownView(err) {
				return err
			}
			panel.Wrap = true
			panel.Title = view.name
			panel.Frame = view.frame
			panel.Editable = view.editable

			if view.list {
				panel.Highlight = true
				panel.SelBgColor = gocui.ColorWhite
				panel.SelFgColor = gocui.ColorBlack
			}
		}
	}
	return nil
}

func (app *App) drawModalViews() error {
	for _, modal := range app.modalViews {
		if panel, err := app.gui.SetView(modal.name, modal.coords.x, modal.coords.y, modal.coords.w, modal.coords.h, 0); err != nil {
			if !gocui.IsUnknownView(err) {
				return err
			}
			panel.Wrap = true
			panel.Title = modal.name
			panel.Frame = true
			panel.Editable = modal.editable
		}
	}
	return nil
}

func (app *App) displayPlugins() error {
	view, err := app.gui.View("Plugins")
	if err != nil {
		return err
	}
	view.Clear()
	keys := make([]string, 0)
	for key := range plugins.Plugins {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for index, key := range keys {
		selector := " "
		if index == app.currentPlugin {
			selector = ">"
		}
		fmt.Fprintf(view, "%s%s\n", selector, key)
	}
	return nil
}

func (app *App) displayLogger() error {
	view, err := app.gui.View("Logger")
	if err != nil {
		return err
	}
	view.Clear()
	switch app.logLevel {
	case LOGINFO:
		view.FgColor = gocui.ColorDefault
	case LOGERROR:
		view.FgColor = gocui.ColorRed
	}
	fmt.Fprintf(view, "Ɇlliot: %s", app.logMsg)
	return nil
}

func (app *App) layout(gui *gocui.Gui) error {
	app.dimensions.width, app.dimensions.height = gui.Size()
	if err := app.drawMainViews(); err != nil {
		return err
	}
	if err := app.drawModalViews(); err != nil {
		return err
	}

	if err := app.displayPlugins(); err != nil {
		return err
	}
	if err := app.displayLogger(); err != nil {
		return err
	}

	if app.currentModal != "" {
		if err := app.showModal(app.currentModal); err != nil {
			return err
		}
	}
	if err := app.setFocus(); err != nil {
		return err
	}

	logrus.Debug("User Interface refreshed")
	return nil
}
