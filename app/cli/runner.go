package cli

import (
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"

	"github.com/cosasdepuma/elliot/app/env"
	"github.com/cosasdepuma/elliot/app/plugins"
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
			resultsView.SetOrigin(0, 0)
			resultsView.SetCursor(0, 0)
			fmt.Fprint(resultsView, strings.TrimSpace(result))

			// FIXME: Bug on cursor
			// -- Despues de mostrar un resultado grande, si se muestra uno pequeño se buguea

			Logger.Type = "Info"
			Logger.Msg = fmt.Sprintf("Plugin '%s' done.", env.Params.Plugin)
		case err := <-env.Channels.Bad:
			Logger.Type = "Error"
			Logger.Msg = err.Error()
		}
		return nil
	})
}
