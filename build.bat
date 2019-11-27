
set GOPATH=x:\si\golang\stree
set GOOS=linux
set GOARCH=386 
go build si001/stree
rename stree stree386

set GOOS=linux
set GOARCH=amd64
go build si001/stree
rename stree stree64

set GOOS=windows
set GOARCH=386
go build si001/stree
rename stree.exe stree386.exe

set GOOS=windows
set GOARCH=amd64
go build si001/stree
rename stree.exe stree64.exe
