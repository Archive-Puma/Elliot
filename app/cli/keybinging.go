package cli

import (
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"

	"github.com/cosasdepuma/elliot/app/env"
	"github.com/cosasdepuma/elliot/app/plugins"
)

func exitApplication(_ *gocui.Gui, _ *gocui.View) error {
	return gocui.ErrQuit
}

func runModule(gui *gocui.Gui, _ *gocui.View) error {
	targetView, err := gui.View("Target")
	if err != nil {
		return err
	}
	env.Params.Target = strings.TrimSpace(targetView.Buffer())

	pluginView, err := gui.View("Plugins")
	if err != nil {
		return err
	}
	_, cy := pluginView.Cursor()
	env.Params.Plugin, err = pluginView.Line(cy)
	if err != nil {
		return err
	}

	if len(env.Params.Arguments) > 0 && env.Params.Arguments[0] == "" {
		env.Params.Arguments = nil
	}
	if env.Params.Arguments == nil {
		switch env.Params.Plugin {
		case "portscanner":
			Modal.Active = true
			Modal.Current = "Ports"
		default:
			env.Params.Arguments = []interface{}{true}
		}
	}
	if env.Params.Arguments != nil {
		go runner(gui)
		Logger.Type = "Info"
		Logger.Msg = fmt.Sprintf("Running %s...", env.Params.Plugin)
	}

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
		if cy+1 < plugins.Amount {
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
	Views := MainViews
	Current = (Current + move) % (len(Views) - 2)
	if _, err := setCurrentViewOnTop(gui, Views[Current].name); err != nil {
		return err
	}

	gui.Cursor = (Current == 0)

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

func setKeybindings(gui *gocui.Gui) error {
	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, exitApplication); err != nil {
		return err
	}
	if err := gui.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, runModule); err != nil {
		return err
	}

	return nil
}
