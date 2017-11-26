#!/usr/bin/env bash
env GOOS=linux GOARCH=386 go build -o bin/linux386/server gitlab.com/gevlee/storage-proxy/cmd/server