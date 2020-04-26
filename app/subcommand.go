package elliot

import "github.com/cosasdepuma/elliot/app/error"

// Subcommand TODO: Doc
type Subcommand interface {
	Check() *error.MrRobotError
	Run() []*error.MrRobotError
}
