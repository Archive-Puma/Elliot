#!/usr/bin/env sh

if test $# -eq 0
then echo "Usage: ${0} <version>"
else curl "https://proxy.golang.org/github.com/cosasdepuma/elliot/@v/${1}.info"
fi