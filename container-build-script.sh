#!/bin/sh

export GO111MODULE=on
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

go version
go env
go build -v -o build/k8s-sdkms-plugin
