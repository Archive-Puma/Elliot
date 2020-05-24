package app

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger TODO: Doc
func (app *App) NewLogger(name string) {
	logrus.SetLevel(logrus.DebugLevel)
	logs, err := os.OpenFile(fmt.Sprintf("%s.log", name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		logs, _ = os.Open(os.DevNull)
		logrus.SetLevel(logrus.PanicLevel)
	}
	logrus.SetOutput(logs)
	app.logFile = logs
}
