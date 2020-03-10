
set GOPATH=x:\si\golang\stree
set GOOS=linux
set GOARCH=386 
go build -o stree386. si001/stree
rem rename stree stree386

set GOOS=linux
set GOARCH=amd64
go build -o stree64. si001/stree
rem rename stree stree64

set GOOS=windows
set GOARCH=386
go build -o stree386.exe si001/stree
rem rename stree.exe stree386.exe

set GOOS=windows
set GOARCH=amd64
go build -o stree64.exe si001/stree
rem rename stree.exe stree64.exe
