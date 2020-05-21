package elliot

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/cosasdepuma/elliot/app/cli"
)

// Entrypoint TODO: Doc
func Entrypoint() {
	logrus.SetLevel(logrus.InfoLevel)
	logs, err := os.OpenFile("elliot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		logs, _ = os.Open(os.DevNull)
		logrus.SetLevel(logrus.PanicLevel)
	}
	defer logs.Close()
	logrus.SetOutput(logs)

	app, err := cli.NewApp()
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
