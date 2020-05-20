package elliot

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/cosasdepuma/elliot/app/gui"
)

// Entrypoint TODO: Doc
func Entrypoint() {
	logs, err := os.OpenFile("elliot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("[!] Error creating logs")
		return
	}
	defer logs.Close()
	logrus.SetOutput(logs)
	logrus.SetLevel(logrus.DebugLevel)

	app, err := gui.NewApp(logs)
	if err != nil {
		fmt.Printf("[!] %s\n", err.Error())
		return
	}

	if err := app.Run(); err != nil {
		fmt.Printf("[!] %s\n", err.Error())
	}
}

/*
	Documentation:
		- Subdomains: A lot of APIs (https://github.com/tomnomnom/assetfinder)
*/
