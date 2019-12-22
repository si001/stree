package model

import (
	"si001/stree/widgets"
	"si001/stree/widgets/stuff"
	"strings"
)

type Directory struct {
	FileInfo
	Parent *widgets.TreeNode
	Files  []*FileInfo
}

func (dir Directory) String() string {
	return dir.Name
}

func (dir Directory) ParsePatch(node *widgets.TreeNode) (path, value string) {
	var sb strings.Builder
	if dir.IsNotRead() {
		sb.WriteRune(stuff.HORIZONTAL_DASH)
	} else if len(node.Nodes) == 0 {
		sb.WriteRune(' ')
	} else if node.Expanded {
		sb.WriteRune(stuff.Theme.Tree.Expanded)
	} else {
		sb.WriteRune(stuff.Theme.Tree.Collapsed)
	}
	createPathRecourse(&sb, node, true)
	path = sb.String()
	if dir.Parent == nil {
		value = " " + dir.Name
	} else {
		value = string(stuff.HORIZONTAL_LINE) + dir.Name
	}
	return path, value
}

func createPathRecourse(sb *strings.Builder, node *widgets.TreeNode, lastLevel bool) {
	dir, _ := node.Value.(Directory)
	if dir.Parent != nil {
		nodeP := dir.Parent
		createPathRecourse(sb, nodeP, false)
		last := nodeP.Nodes[len(nodeP.Nodes)-1] == node
		if last {
			if lastLevel {
				sb.WriteString(" " + string(stuff.BOTTOM_LEFT))
			} else {
				sb.WriteString("  ")
			}
		} else if lastLevel {
			sb.WriteString(" " + string(stuff.VERTICAL_RIGHT))
		} else {
			sb.WriteString(" " + string(stuff.VERTICAL_LINE))
		}
	}
}
