package cli

import (
	"time"

	"github.com/awesome-gocui/gocui"

	"github.com/cosasdepuma/elliot/app/env"
)

func (app *App) refresher(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			app.gui.Update(func(_ *gocui.Gui) error { return nil })
		case <-env.Channels.Stop:
			return
		}
	}
}
