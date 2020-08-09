package main

import (
	"fmt"

	elliot "github.com/cosasdepuma/elliot/pkg/app"
)

// Elliot Dev Resources:
// -- Template: https://colorlib.com/polygon/admindek/default/index.html
// -- Template Source Code: https://github.com/baotm/admindek
// -- Icons: https://feathericons.com/
// -- Bootstrap Doc: https://getbootstrap.com/docs/4.5/getting-started/introduction/

func main() {
	// TODO: SiteCheck in Frontend

	fmt.Println("🤖  Starting\t\033[1;33mElliot\033[0m")
	elliot.Start()
}
