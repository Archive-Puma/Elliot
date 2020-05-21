package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cosasdepuma/elliot/app/env"
	"github.com/cosasdepuma/elliot/app/plugins"
	"github.com/sirupsen/logrus"
)

func (app *App) getTarget() error {
	view, err := app.gui.View("Target")
	if err != nil {
		return err
	}
	env.Config.Target = strings.TrimSpace(view.Buffer())
	logrus.Info("Target: ", env.Config.Target)
	return nil
}

func (app *App) getPlugin() error {
	view, err := app.gui.View("Plugins")
	if err != nil {
		return err
	}
	_, cy := view.Cursor()
	plugin, err := view.Line(cy)
	if err != nil {
		return err
	}
	app.pluginName = strings.TrimPrefix(plugin, ">")
	logrus.Info("Selected plugin: ", app.pluginName)
	return nil
}

func (app *App) getParams() error {
	if modal, ok := plugins.Required[app.pluginName]; ok {
		return app.showModal(modal)
	}
	app.sendStartSignal()
	return nil
}

func (app *App) runPlugin() {
	go func(app *App) {
		<-env.Channels.Start
		go app.runner()
		// app.startLoader
		select {
		case err := <-env.Channels.Bad:
			app.logLevel = LOGERROR
			app.logMsg = err.Error()
		case results := <-env.Channels.Ok:
			app.logLevel = LOGINFO
			app.logMsg = "Done."
			if view, err := app.gui.View("Results"); err == nil {
				view.Clear()
				view.SetOrigin(0, 0)
				view.SetCursor(0, 0)
				fmt.Fprintf(view, "%s\n\n\n", results)
			}
		}
		// app.stopLoader
	}(app)
}

func (app *App) runner() {
	plugin, ok := plugins.Plugins[app.pluginName]
	if !ok {
		env.Channels.Bad <- errors.New("Plugin not found")
		return
	}
	logrus.Info("Starting plugin")

	if app.runningInstances == 0 {
		app.lock.Lock()
		app.runningInstances++
		app.lock.Unlock()

		app.logLevel = LOGINFO
		app.logMsg = fmt.Sprintf("Running %s...", app.pluginName)
		plugin.Run()

		app.lock.Lock()
		app.runningInstances--
		app.lock.Unlock()
	}
}