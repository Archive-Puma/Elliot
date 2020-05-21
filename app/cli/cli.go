package cli

import (
	"sync"
	"time"

	"github.com/awesome-gocui/gocui"
	"github.com/sirupsen/logrus"

	"github.com/cosasdepuma/elliot/app/env"
)

const (
	// LOGINFO displays no colored messages
	LOGINFO = iota
	// LOGERROR displays highlighted error messages
	LOGERROR
)

// App contains all the information related to the application displayed on the screen
type App struct {
	gui              *gocui.Gui
	lock             *sync.Mutex
	logMsg           string
	logLevel         int
	mainViews        []sView
	modalViews       map[string]sView
	keybindings      []sKeybinding
	currentPlugin    int
	pluginName       string
	currentModal     string
	currentView      int
	lastView         int
	runningInstances int
	dimensions       struct {
		width, height int
	}
}

// NewApp generates a new user interface
func NewApp() (*App, error) {
	gui, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		return nil, err
	}
	return &App{
		gui:              gui,
		lock:             &sync.Mutex{},
		logMsg:           "",
		logLevel:         LOGINFO,
		lastView:         0,
		currentView:      0,
		currentPlugin:    0,
		runningInstances: 0,
		mainViews: []sView{
			{name: "Target", coords: coordinates{0, 0, -1, 2}, frame: true, editable: true},
			{name: "Plugins", coords: coordinates{0, 3, 18, -4}, frame: true, list: true},
			{name: "Results", coords: coordinates{19, 3, -1, -4}, frame: true, editable: true},
			{name: "Logger", coords: coordinates{-1, -4, 0, -2}},
			{name: "─", coords: coordinates{-1, -2, 0, 0}, frame: true},
		},
		modalViews: map[string]sView{
			"Ports": {name: "Ports", coords: coordinates{-20, -1, -18, 1}, editable: true},
		},
		currentModal: "",
	}, nil
}

// Run allows to start the application
func (app *App) Run() error {
	defer logrus.Debug("Application closed")
	// Configuration
	defer app.gui.Close()
	env.Channels.Ok = make(chan string)
	defer close(env.Channels.Ok)
	env.Channels.Bad = make(chan error)
	defer close(env.Channels.Bad)
	env.Channels.Done = make(chan struct{})
	defer close(env.Channels.Done)
	env.Channels.Start = make(chan struct{})
	defer close(env.Channels.Start)
	env.Channels.Stop = make(chan struct{})
	defer close(env.Channels.Stop)
	logrus.Debug("Application configured")

	// Init the refresher
	go app.refresher(time.Second)
	logrus.Debug("Refresher running")

	// Theme
	app.gui.Highlight = true
	app.gui.FgColor = gocui.ColorDefault
	app.gui.SelFgColor = gocui.ColorCyan
	logrus.Debug("Style established")

	// Manager
	app.gui.SetManager(gocui.ManagerFunc(app.layout), gocui.ManagerFunc(app.getFocusLayout()))
	logrus.Debug("Managers configured")

	// Keybindings
	if err := app.setKeybindings(); err != nil {
		return err
	}

	// Main Loop
	logrus.Debug("Running user interface")
	if err := app.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}
