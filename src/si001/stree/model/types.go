package model

import (
	"github.com/gizak/termui/v3/widgets"
	"time"
)

type FileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
	Attr    byte
}

const (
	ATTR_NOTREAD     = 0
	ATTR_FILE        = 1
	ATTR_DIR         = 2
	ATTR_ARCH        = 4
	ATTR_ERR_MESSAGE = 255
)

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
