package _old

import (
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"

	"github.com/cosasdepuma/elliot/app/env"
)

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
