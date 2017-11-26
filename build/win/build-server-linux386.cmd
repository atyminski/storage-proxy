@echo off
SET GOOS=linux
SET GOARCH=386
SET /p VERSION= < VERSION
go build -o bin/linux386/server -ldflags "-X gitlab.com/gevleeog/storage-proxy/cmd/server/app.Version=%VERSION%" gitlab.com/gevleeog/storage-proxy/cmd/server