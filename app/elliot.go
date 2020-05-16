package elliot

import "github.com/cosasdepuma/elliot/app/cli"

// Entrypoint TODO: Doc
func Entrypoint() {
	if err := cli.ShowUI(); err != nil {
		err.Resolve(true)
	}
}

/*
	Documentation:
		- Subdomains: A lot of APIs (https://github.com/tomnomnom/assetfinder)
*/
