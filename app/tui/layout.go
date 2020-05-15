package tui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func mainLayout(gui *gocui.Gui) error {
	width, height := gui.Size()

	for index, view := range views {
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

		if panel, err := gui.SetView(view.name, x, y, w, h); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			panel.Wrap = true
			panel.Title = view.name
			panel.Frame = view.frame

			if view.editable {
				panel.Editable = true
				if err := gui.SetKeybinding(view.name, gocui.KeyEnter, gocui.ModNone, disable); err != nil {
					return err
				}
			}
		}

		if index == active {
			if _, err := setCurrentViewOnTop(gui, view.name); err != nil {
				return err
			}
		}
	}

	panel, err := gui.View(views[len(views)-1].name)
	if err != nil {
		return err
	}
	panel.Clear()
	fmt.Fprint(panel, "Shortcuts: [^C] Exit [TAB] Next")

	return nil
}
