@echo off
SET GOOS=linux
SET GOARCH=386
go build -o bin/linux386/server -ldflags gitlab.com/gevlee/storage-proxy/cmd/server