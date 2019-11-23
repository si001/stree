package files

import (
	"github.com/gizak/termui/v3/widgets"
	"os"
	"runtime"
	"si001/stree/model"
)

func newDir(nm string, parent *widgets.TreeNode) (dir *widgets.TreeNode) {
	dir = newDirFI(model.FileInfo{
		Name: nm,
		Attr: 1,
	}, parent)
	return dir
}

func newDirFI(fInfo model.FileInfo, parent *widgets.TreeNode) (dir *widgets.TreeNode) {
	dir = &widgets.TreeNode{
		Value: model.Directory{
			FileInfo: fInfo,
			Parent:   parent,
		},
	}
	if parent != nil {
		parent.Nodes = append(parent.Nodes, dir)
	}
	return dir
}

func GetRoot() (r *widgets.TreeNode) {
	if runtime.GOOS == "windows" {
		r = newDir("", nil)
		for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
			_, err := os.Open(string(drive) + ":\\")
			if err == nil {
				newDir(string(drive)+":", r)
			}
		}
	} else {
		r = newDir("", nil)
		newDir("/", r)
	}
	return r
}
