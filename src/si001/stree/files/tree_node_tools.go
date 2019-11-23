package files

import (
	"github.com/gizak/termui/v3/widgets"
	"si001/stree/model"
)

func NodeSetParent(node *widgets.TreeNode, parent *widgets.TreeNode) {
	switch d := node.Value.(type) {
	case model.Directory:
		d.Parent = parent
	}
}

func NodeGetParent(node *widgets.TreeNode) *widgets.TreeNode {
	switch d := node.Value.(type) {
	case model.Directory:
		return d.Parent
	}
	return nil
}
