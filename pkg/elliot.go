package elliot

import (
	"fmt"
	"os"

	"github.com/cosasdepuma/elliot/pkg/app"
)

// Entrypoint defines the starting point of the program
func Entrypoint() {
	// Create the app
	elliot, err := app.NewApp("elliot")
	handle(err)
	// Destroy the app
	defer elliot.Destroy()
}

func handle(err error) {
	if err != nil {
		fmt.Sprintln("[!] %s", err.Error())
		os.Exit(1)
	}
}
