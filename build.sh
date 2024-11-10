#!/bin/bash
set -e
selfDir=$(cd $(dirname $0); pwd)
projectDir=$selfDir
cd $projectDir

GOOS=linux GOARCH=amd64 go build -o ./debug/tuissh-linux-x64 .
GOOS=linux GOARCH=arm64 go build -o ./debug/tuissh-linux-arm .

GOOS=windows GOARCH=amd64 go build -o ./debug/tuissh-win-x64.exe .

GOOS=darwin GOARCH=arm64 go build -o ./debug/tuissh-mac-arm .
