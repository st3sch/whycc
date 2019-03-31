#!/bin/bash
VERSION=`git describe --abbrev=0 --tags`
COMMIT=`git log --pretty=format:'%h' -n 1`
DATE=`date +"%Y%m%d%H%M%S"`
GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$VERSION -X main.commit=$COMMIT -X main.date=$DATE" -o whycc.exe whycc.go