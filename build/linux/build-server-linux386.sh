#!/usr/bin/env bash
env GOOS=linux GOARCH=386 go build -o bin/linux386/server -ldflags "-X gitlab.com/gevleeog/storage-proxy/cmd/server/app.Version=$(cat VERSION)" gitlab.com/gevleeog/storage-proxy/cmd/server