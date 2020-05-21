<h1 align="center">
  <img src="https://raw.githubusercontent.com/CosasDePuma/Elliot/master/.github/readme/elliot.gif" alt="Elliot" width="500">
  <br><br>
  <img src="https://raw.githubusercontent.com/CosasDePuma/Elliot/master/.github/readme/logo.png" alt="Logo" width="600">
</h1>

[![Golang](https://img.shields.io/github/go-mod/go-version/cosasdepuma/elliot?style=for-the-badge)](https://pkg.go.dev/mod/github.com/cosasdepuma/elliot)
[![Go Report Card](https://goreportcard.com/badge/github.com/cosasdepuma/elliot?style=for-the-badge)](https://goreportcard.com/report/github.com/cosasdepuma/elliot)
[![Latest Version](https://img.shields.io/badge/latest-v0.0.2-green?style=for-the-badge)](https://github.com/CosasDePuma/Elliot/releases/)
[![License](https://img.shields.io/github/license/cosasdepuma/elliot?style=for-the-badge&color=important)](./LICENSE)

If you do not know **Elliot**, you are not aware of the number of possibilities that you are wasting when it comes to perform your pentestings. A new all-in-one hacking framework is going to be unleashed... or is it just a product of your imagination?

🖥️ Installation
---
To install the tool, the easiest way is to use the `go get` command:

```go
go get github.com/cosasdepuma/elliot
```

You can also download the precompiles binary for your system in the [**Release**](https://github.com/CosasDePuma/Elliot/releases) tab.

🐋 Dockerize
---
You can execute the application in containerized environments like Docker. To download the image, just run:

```sh
docker pull cosasdepuma/elliot:latest
```

The recommended way to run the image is:

```
docker run --rm -it cosasdepuma/elliot
```

🔩 Develop
---
You can develop your own modules or contribute to the development and improvements of the project freely.

The first thing you need to do is clone the project using the `git` tool:

```sh
git clone https://github.com/cosasdepuma/Elliot
```

Once downloaded, do not hesitate to modify everything you think necessary.

You should take a look at the [TODO](./TODO.md) and [Roadmap](./ROADMAP.md) files to get an idea of what needs to be done.

🔧 Compile
---

Compiling Elliot is extremely easy. Just run the command:

```go
go build -o elliot main.go
```

You can also specify both the operating system and the target architecture:

```go
GOOS=windows GOARCH=amd64 go build -o elliot.exe main.go
```

You can automate the compilation process using the script found inside the **scripts** folder called `x-build`. It supports the following os/arch:

| Supported x32 | Supported x64 |
| --- | --- |
| windows/386 | windows/am64 |
| linux/386 | linux/amd64 |
| darwin/386 | darwin/amd64 |
| freebsd/386 | freebsd/amd64 |

🐙 Support the developer!
----
Everything I do and publish can be used for free whenever I receive my corresponding merit.

Anyway, if you want to help me in a more direct way, you can leave me a tip by clicking on this badge:

<p align="center">
    </br>
    <a href="https://www.paypal.me/cosasdepuma/"><img src="https://img.shields.io/badge/Donate-PayPal-blue.svg?style=for-the-badge" alt="PayPal Donation"></a>
</p>