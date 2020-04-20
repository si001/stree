#!/bin/sh

env GOPATH=/home/si001/dev/stree GOOS=linux GOARCH=386 go build -o stree386 si001/stree

env GOPATH=/home/si001/dev/stree GOOS=linux GOARCH=amd64 go build -o stree64 si001/stree

env GOPATH=/home/si001/dev/stree GOOS=windows GOARCH=386 go build -o stree386.exe si001/stree

env GOPATH=/home/si001/dev/stree GOOS=windows GOARCH=amd64 go build -o stree64.exe si001/stree

env GOPATH=/home/si001/dev/stree GOOS=linux GOARCH=arm64 go build -o streeArm64 si001/stree
