#!/usr/bin/env sh

rm -rf **/*_templ.go
./bin/flare fix-workspace
./bin/flare build-plugins
./bin/flare server
