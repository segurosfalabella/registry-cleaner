#!/bin/sh

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
export GOARM=7

go build -a -installsuffix cgo
ls -la
echo "built"