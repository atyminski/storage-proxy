#!/usr/bin/env bash

echo "Installing github.com/golang/dep/cmd/dep ..."
go get -u github.com/golang/dep/cmd/dep
echo "Resolving dependencies ..."
$GOPATH/bin/dep ensure