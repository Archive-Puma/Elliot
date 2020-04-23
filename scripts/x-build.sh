#!/usr/bin/env sh

#### Description: Cross compiling script
#### Author: Kike Fontán (@CosasDePuma)
#### See also: https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63

### Supported Architectures: x86 (386), x64 (amd64)
### Supported Operative Systems: Darwin, FreeBSD, Linux, Windows

APPNAME="elliot"
ARCH="386 amd64"
OS="darwin freebsd linux windows"

# Directories

PROJECTDIR=$(cd "$(dirname "${0}")"/.. && pwd)
DISTDIR="${PROJECTDIR}/dist"

test -d "${DISTDIR}" && rm -rf "${DISTDIR}"
mkdir "${DISTDIR}"

# Cross compiling

for GOOS in $OS; do
  for GOARCH in $ARCH; do
    echo "[+] Compiling ${APPNAME} for ${GOOS}/${GOARCH}"

    BINFILE="${APPNAME}-${GOOS}-${GOARCH}"
    test "${GOOS}" = "windows" && BINFILE="${BINFILE}.exe"

    GOOS="${GOOS}" GOARCH="${GOARCH}" go build -o "${DISTDIR}/${BINFILE}" ../main.go
  done
done