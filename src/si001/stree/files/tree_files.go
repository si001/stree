package files

import (
	"fmt"
	"si001/stree/model"
	"si001/stree/screen/botton_box/actions"
	"si001/stree/widgets"
	"strings"
)

func BuildTree(path string) *widgets.Tree {
	ph := strings.Split(path, model.PathDivider)
	var nodePath []*widgets.TreeNode
	var root widgets.TreeNode
	var dir *widgets.TreeNode
	for i, dirNm := range ph {
		if i == 0 {
			if dirNm == "" {
				dirNm = model.PathDivider
			}
			root = *GetRoot()
			dir = &root
			for _, rootElt := range dir.Nodes {
				if dirNm == rootElt.Value.String() {
					dir = rootElt
					break
				}
			}
		} else {
			dir = NewDir(dirNm, dir)
		}
		nodePath = append(nodePath, dir)
		dir.Value = ReadDir(dir)
	}
	var l = widgets.NewTree()
	l.WrapText = false
	l.SetNodes(root.Nodes)
	l.ScrollTop()
	counter := 77 // anti-loop
	for _, node := range nodePath {
		if node.Value.String() == "" {
			continue
		}
		for node != l.SelectedNode() && counter > 0 {
			l.ScrollDown()
			counter--
		}
		l.Expand()
	}
	return l
}

func LogTree(path string, tree *widgets.Tree) *widgets.Tree {
	ph := strings.Split(path, model.PathDivider)
	var nodePath []*widgets.TreeNode
	var root *widgets.TreeNode
	var dir *widgets.TreeNode
	for i, dirNm := range ph {
		if len([]rune(strings.TrimSpace(dirNm))) == 0 {
			continue
		}
		dirNm = UpcaseIfWindows(dirNm)
		if i == 0 {
			if dirNm == "" {
				dirNm = model.PathDivider
			}
			read := GetRoot()
			tRoot := &widgets.TreeNode{
				Nodes: tree.GetRoot(),
			}
			root, dir = checkNodes(tRoot, read, nil, dirNm)
		} else {
			//dir = newDir(dirNm, dir)
			old := dir
			for _, d := range dir.Nodes {
				if UpcaseIfWindows(d.Value.String()) == dirNm {
					dir = d
					break
				}
			}
			if dir == old || dir == nil {
				actions.RequestMessageBoxCenter(fmt.Sprintf("Path '%s' is incorrect!", path), nil)
				return tree
			}
		}
		if dir == nil {
			actions.RequestMessageBoxCenter(fmt.Sprintf("Path '%s' is incorrect!", path), nil)
			return tree
		}
		nodePath = append(nodePath, dir)
		var old widgets.TreeNode
		(old.Value) = &model.Directory{
			FileInfo: (dir.Value).(*model.Directory).FileInfo,
		}
		(old.Value).(*model.Directory).FileInfo.Owner = nil
		path := TreeNodeToPath(dir)
		old.Value = ReadDirPath(&old, path)
		(old.Value).(*model.Directory).FileInfo.Owner = (dir.Value).(*model.Directory).FileInfo.Owner
		_, _ = checkNodes(dir, &old, dir, "")
	}
	tree.WrapText = false
	tree.SetNodes(root.Nodes)
	tree.ScrollTop()
	for _, node := range nodePath {
		if node.Value.String() == "" {
			continue
		}
		var old *widgets.TreeNode = nil
		for node != tree.SelectedNode() && old != tree.SelectedNode() {
			old = tree.SelectedNode()
			tree.ScrollDown()
		}
		tree.Expand()
	}
	return tree
}

func checkNodes(exist, read, parent *widgets.TreeNode, dirNm string) (*widgets.TreeNode, *widgets.TreeNode) {
	var dir *widgets.TreeNode
	j := 0
	i := 0
	var rd, ex *widgets.TreeNode
	var rds, exs string
	for i < len(read.Nodes) || j < len(exist.Nodes) {
		rd, ex = nil, nil
		rds, exs = "", ""
		if i < len(read.Nodes) {
			rd = read.Nodes[i]
			rds = UpcaseIfWindows(rd.Value.String())
		}
		if j < len(exist.Nodes) {
			ex = exist.Nodes[j]
			exs = UpcaseIfWindows(ex.Value.String())
		}
		if rd != nil && ex != nil && exs == rds {
			i++
			j++
		} else if (rd == nil) || (ex != nil && exs < rds) {
			exist.Nodes = append(exist.Nodes, read.Nodes[j])
		} else {
			n := exist.Nodes
			exist.Nodes = append(exist.Nodes[:j], rd)
			rd.Value.(*model.Directory).FileInfo.Owner = parent
			j++
			if j < len(exist.Nodes) {
				exist.Nodes = append(exist.Nodes, n[j:]...)
			}
			i++
		}
	}
	for _, n := range exist.Nodes {
		if dirNm == UpcaseIfWindows(n.Value.String()) {
			dir = n
			break
		}
	}
	return exist, dir
}
