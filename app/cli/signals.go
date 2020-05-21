package cli

import (
	"github.com/cosasdepuma/elliot/app/env"
)

func (app *App) sendStartSignal() {
	go func(start chan struct{}) { start <- struct{}{} }(env.Channels.Start)
}

func (app *App) cleanSignals() {
	for {
		select {
		case <-env.Channels.Start:
		case <-env.Channels.Done:
		case <-env.Channels.Ok:
		case <-env.Channels.Bad:
		default:
			return
		}
	}
}
