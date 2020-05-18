package elliot

import (
	"fmt"

	"github.com/cosasdepuma/elliot/app/cli"
)

// Entrypoint TODO: Doc
func Entrypoint() {
	if err := cli.ShowUI(); err != nil {
		fmt.Printf("[!] %s\n", err.Error())
	}
}

/*
	Documentation:
		- Subdomains: A lot of APIs (https://github.com/tomnomnom/assetfinder)
*/
