package model

import (
	"github.com/gizak/termui/v3/widgets"
	"time"
)

type FileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
	Attr    int8
}

type Directory struct {
	FileInfo
	Parent *widgets.TreeNode
	Files  []*FileInfo
}

func (dir Directory) String() string {
	return dir.Name
}

//func (n widgets.TreeNode) Directory() *Directory {
//	return n.
//}
