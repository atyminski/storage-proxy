@echo off
SET GOOS=windows
SET GOARCH=386
SET /p VERSION= < VERSION
go build -o bin/win386/server.exe -ldflags "-X gitlab.com/gevleeog/storage-proxy/cmd/server/app.Version=%VERSION%" gitlab.com/gevleeog/storage-proxy/cmd/server