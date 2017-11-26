@echo off
SET GOOS=windows
SET GOARCH=386
go build -o bin/win386/server.exe gitlab.com/gevlee/storage-proxy/cmd/server