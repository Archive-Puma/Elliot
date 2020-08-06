package main

import (
	"fmt"

	elliot "github.com/cosasdepuma/elliot/app"
	"github.com/cosasdepuma/elliot/app/modules"
)

// Elliot Dev Resources:
// -- Template: https://colorlib.com/polygon/admindek/default/index.html
// -- Template Source Code: https://github.com/baotm/admindek
// -- Icons: https://feathericons.com/
// -- Bootstrap Doc: https://getbootstrap.com/docs/4.5/getting-started/introduction/

func main() {
	// Testing
	// FIXME: Inactive first time in Elliot with some data in DB
	// TODO: SiteCheck in Frontend
	c := make(chan map[string]interface{}, 1)
	defer close(c)
	modules.ModuleSiteCheck("uvigo.es", &c)
	fmt.Println(<-c)

	elliot.Backend.Start()
}
