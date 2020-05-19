package cli

import (
	"fmt"
	"sort"

	"github.com/awesome-gocui/gocui"

	"github.com/cosasdepuma/elliot/app/plugins"
)

func displayPlugins(gui *gocui.Gui) error {
	pluginView, err := gui.View("Plugins")
	if err != nil {
		return err
	}
	pluginView.Clear()
	keys := make([]string, 0)
	for key := range plugins.Plugins {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Fprintln(pluginView, key)
	}
	return nil
}

func displayShortcuts(gui *gocui.Gui) error {
	shortcutsView, err := gui.View(MainViews[len(MainViews)-1].name)
	if err != nil {
		return err
	}
	shortcutsView.Clear()
	fmt.Fprint(shortcutsView, "Shortcuts: [^C] Exit [TAB] Next Frame [Enter] Run ")

	// Custom shortcuts
	if Current == 1 {
		fmt.Fprint(shortcutsView, "[Up|Down] Navigate")
	}
	return nil
}

func configureView(gui *gocui.Gui, view *gocui.View, name string, hasFrame bool, isEditable bool, isList bool) error {
	view.Wrap = true
	view.Title = name
	view.Frame = hasFrame
	view.Editable = isEditable

	if err := gui.SetKeybinding(name, gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return err
	}

	if isList {
		view.Highlight = true
		view.SelBgColor = gocui.ColorWhite
		view.SelFgColor = gocui.ColorBlack

		if err := gui.SetKeybinding(name, gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
			return err
		}
		if err := gui.SetKeybinding(name, gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
			return err
		}
	}
	return nil
}

func mainLayout(gui *gocui.Gui) error {
	width, height := gui.Size()

	gui.Cursor = true
	gui.Highlight = true
	gui.SelFgColor = gocui.ColorCyan

	Views := MainViews
	for index, view := range Views {
		// Calculate position
		x, y, w, h := calculatePosition(width, height, view)
		// Create the view
		if panel, err := gui.SetView(view.name, x, y, w, h, 1); err != nil {
			if !gocui.IsUnknownView(err) {
				return err
			}
			if err := configureView(gui, panel, view.name, view.frame, view.editable, view.list); err != nil {
				return err
			}
		}
		// Set the current view on top
		if index == Current {
			if _, err := setCurrentViewOnTop(gui, view.name); err != nil {
				return err
			}
		}
	}

	// Display the plugins
	if err := displayPlugins(gui); err != nil {
		return err
	}
	// Display the logger
	if err := displayLogger(gui); err != nil {
		return err
	}

	// Display the shortcuts
	if err := displayShortcuts(gui); err != nil {
		return err
	}

	if err := displayModal(gui); err != nil {
		return err
	}

	return nil
}
