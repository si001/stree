#!/bin/sh

env GOPATH=/home/si001/dev/stree GOOS=linux GOARCH=386 go build si001/stree
mv stree stree386

env GOPATH=/home/si001/dev/stree GOOS=linux GOARCH=amd64 go build si001/stree
mv stree stree64

env GOPATH=/home/si001/dev/stree GOOS=windows GOARCH=386 go build si001/stree
mv stree.exe stree386.exe

env GOPATH=/home/si001/dev/stree GOOS=windows GOARCH=amd64 go build si001/stree
mv stree.exe stree64.exe
