#!/usr/bin/env bash
env GOOS=windows GOARCH=386 go build -o bin/windows386/server.exe gitlab.com/gevlee/storage-proxy/cmd/server