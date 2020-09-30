package modules

import (
	"flag"

	"github.com/cosasdepuma/elliot/pkg/app"
)

// Module TODO
type Module struct {
	Flag *flag.FlagSet
	Run  func(*app.Core)
}
