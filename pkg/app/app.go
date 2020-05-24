package app

import (
	"os"

	"github.com/awesome-gocui/gocui"
	"github.com/sirupsen/logrus"

	"github.com/cosasdepuma/elliot/pkg/gui"
)

// App TODO: Doc
type App struct {
	// -- Gui related
	gui           *gocui.Gui
	width, height int
	keybindings   []gui.Keybinding
	// -- Log related
	logFile *os.File
}

// NewApp TODO: Doc
func NewApp(name string) (*App, error) {
	defer logrus.Debug("Application created")
	// Create the Application
	app := new(App)
	// Create the Logger
	app.NewLogger(name)
	// Create the Gui
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		return nil, err
	}
	app.gui = g
	// Set the Keybindings
	app.SetKeybinding()
	// Create the app
	return app, nil
}

// Destroy TODO: Doc
func (app *App) Destroy() {
	defer app.logFile.Close()
	app.gui.Close()
	logrus.Debug("Application destroyed")
}
