package files

import (
	"si001/stree/model"
	"si001/stree/widgets"
)

//func NodeSetParent(node *widgets.TreeNode, parent *widgets.TreeNode) {
//	var value interface{} = node.Value
//	var dir = value.(model.Directory)
//	dir.Parent = parent
//}

func NodeGetFiles(node *widgets.TreeNode) []*model.FileInfo {
	switch d := node.Value.(type) {
	case *model.Directory:
		return d.Files
	}
	return nil
}

func NodeSetFiles(node *widgets.TreeNode, files []*model.FileInfo) {
	switch d := node.Value.(type) {
	case model.Directory:
		d.Files = files
	}
}
