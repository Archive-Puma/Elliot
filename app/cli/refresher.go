package cli

import "time"

func (app *App) refresher(interval time.Duration, function func() error) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			_ = function()
		case <-app.Channels.Stop:
			return
		}
	}
}
