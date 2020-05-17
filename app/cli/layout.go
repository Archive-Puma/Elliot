package cli

import (
	"fmt"
	"sort"

	"github.com/jroimartin/gocui"

	plgns "github.com/cosasdepuma/elliot/app/plugins"
)

func mainLayout(gui *gocui.Gui) error {
	width, height := gui.Size()

	gui.Cursor = true
	gui.Highlight = true
	gui.SelFgColor = gocui.ColorCyan

	for index, view := range Views {
		// Calculate position
		x := view.x
		if x < -1 {
			x = width + x
		}
		y := view.y
		if y < -1 {
			y = height + y
		}
		w := view.w
		if w <= 0 {
			w = width + w
		}
		h := view.h
		if h <= 0 {
			h = height + h
		}
		// Create the view
		if panel, err := gui.SetView(view.name, x, y, w, h); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			panel.Wrap = true
			panel.Title = view.name
			panel.Frame = view.frame
			panel.Editable = view.editable

			if err := gui.SetKeybinding(view.name, gocui.KeyTab, gocui.ModNone, nextView); err != nil {
				return err
			}

			if view.list {
				panel.Highlight = true
				panel.SelBgColor = gocui.ColorWhite
				panel.SelFgColor = gocui.ColorBlack

				if err := gui.SetKeybinding(view.name, gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
					return err
				}
				if err := gui.SetKeybinding(view.name, gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
					return err
				}
			}
		}

		if index == Current {
			if _, err := setCurrentViewOnTop(gui, view.name); err != nil {
				return err
			}
		}
	}
	// Display plugins
	plugins, err := gui.View("Plugins")
	if err != nil {
		return err
	}
	plugins.Clear()
	keys := make([]string, 0)
	for key := range plgns.Plugins {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Fprintln(plugins, key)
	}
	// Display logs
	logger, err := gui.View(Views[len(Views)-2].name)
	if err != nil {
		return err
	}
	logger.Clear()
	switch Logger.Type {
	case "Error":
		logger.FgColor = gocui.ColorRed
	case "Info":
		logger.FgColor = gocui.ColorDefault
	}
	fmt.Fprintf(logger, "%s: %s", Logger.Type, Logger.Msg)

	// Display shortcuts
	shortcuts, err := gui.View(Views[len(Views)-1].name)
	if err != nil {
		return err
	}
	shortcuts.Clear()
	fmt.Fprint(shortcuts, "Shortcuts: [^C] Exit [TAB] Next Frame [Enter] Run ")

	// Custom shortcuts
	if Current == 1 {
		fmt.Fprint(shortcuts, "[Up|Down] Navigate")
	}

	return nil
}
