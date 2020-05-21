package cli

import "github.com/cosasdepuma/elliot/app/env"

func (app *App) sendStartSignal() {
	go func(start chan struct{}) { start <- struct{}{} }(env.Channels.Start)
}
