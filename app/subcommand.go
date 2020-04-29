package elliot

import "github.com/cosasdepuma/elliot/app/error"

// Subcommand TODO: Doc
type Subcommand interface {
	Help()
	Check() *error.MrRobotError
	Run() ([]string, []*error.MrRobotError)
}
