#!/bin/sh

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=arm

go build -a -installsuffix cgo
ls -la
echo "built"