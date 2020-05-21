package gui

import (
	"os"
	"time"

	"github.com/awesome-gocui/gocui"
	"github.com/sirupsen/logrus"
)

const (
	// LOGINFO TODO: Doc
	LOGINFO = iota
	// LOGERROR TODO Doc
	LOGERROR
)

// App TODO: Doc
type App struct {
	gui           *gocui.Gui
	logMsg        string
	logLevel      int
	mainViews     []sView
	modalViews    map[string]sView
	keybindings   []sKeybinding
	currentPlugin int
	currentModal  string
	currentView   int
	lastView      int
	dimensions    struct {
		width, height int
	}
	Params   interface{}
	Channels struct {
		Ok   chan string
		Bad  chan error
		Stop chan struct{}
	}
}

// NewApp TODO: Doc
func NewApp(logs *os.File) (*App, error) {
	gui, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		return nil, err
	}
	return &App{
		gui:           gui,
		logMsg:        "",
		logLevel:      LOGINFO,
		lastView:      0,
		currentView:   0,
		currentPlugin: 0,
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
		Params:       nil,
		currentModal: "",
	}, nil
}

// Run TODO: Doc
func (app *App) Run() error {
	defer logrus.Debug("Application closed")
	// Configuration
	defer app.gui.Close()
	app.Channels.Ok = make(chan string)
	defer close(app.Channels.Ok)
	app.Channels.Bad = make(chan error)
	defer close(app.Channels.Bad)
	app.Channels.Stop = make(chan struct{})
	defer close(app.Channels.Stop)
	logrus.Debug("Application configured")

	// Init the refresher
	go app.refresher(time.Second*10, func() error { return nil })
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
