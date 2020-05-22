// Copyright 2020 Kike Fontán (@CosasDePuma). All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package elliot defines the core of the program with the same name.

If you do not know Elliot, you are not aware of the number of possibilities that you are wasting when it comes to perform your pentestings.
A new all-in-one hacking framework is going to be unleashed... or is it just a product of your imagination?

	┌─Target────────────────────────────────────────────┐
	│pkg.go.dev                                         │
	└───────────────────────────────────────────────────┘
	┌─Plugins─────────┐┌─Results────────────────────────┐
	│>portscanner     ││     PORT    STATE      SERVICE │
	│ robots.txt      ││   80/tcp    open       http    │
	│ subdomain       ││  443/tcp    open       https   │
	└─────────────────┘└────────────────────────────────┘
	Ɇlliot: Done.
	─────────────────────────────────────────────────────
	Shortcuts: [^C] Exit [TAB] Next Frame [Enter] Run

Elliot not only has a constantly growing variety of plugins that will help you to perform basic pentesting tests, but it is also a tool with a very good performance, due to its purely Golang-based implementation.

Currently the available plugins are:

	portscanner := scans for open ports
	robots.txt  := returns the robots.txt of a web page
	subdomain   := collects from different sources all subdomains associated with a domain

You can also execute the application in containerized environments like Docker. To download the image, just run:

	docker pull cosasdepuma/elliot:latest

The recommended way to run the image is:

	docker run --rm -it cosasdepuma/elliot

For more information about Elliot, check out his repository on GitHub: https://github.com/cosasdepuma/elliot
*/
package main
