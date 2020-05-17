package cli

import (
	"fmt"

	"github.com/cosasdepuma/elliot/app/env"
	"github.com/cosasdepuma/elliot/app/plugins"
	"github.com/jroimartin/gocui"
)

func runner(gui *gocui.Gui) {
	go plugins.Plugins[env.Params.Plugin].Run()

	gui.Update(func(gui *gocui.Gui) error {
		resultsView, err := gui.View("Results")
		if err != nil {
			return err
		}
		resultsView.Clear()

		// Ok
		select {
		case result := <-env.Channels.Ok:
			resultsView, err := gui.View("Results")
			if err != nil {
				return err
			}
			resultsView.Clear()
			fmt.Fprint(resultsView, result)
			Logger.Type = "Info"
			Logger.Msg = fmt.Sprintf("Plugin '%s' done.", env.Params.Plugin)
		case err := <-env.Channels.Bad:
			Logger.Type = "Error"
			Logger.Msg = err.Error()
		}
		return nil
	})
}
