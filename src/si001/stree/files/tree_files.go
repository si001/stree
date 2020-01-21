package files

import (
	"runtime"
	"si001/stree/model"
	"si001/stree/widgets"
	"strings"
)

func BuildTree(path string) *widgets.Tree {
	model.PathDivider = "/"
	if runtime.GOOS == "windows" {
		model.PathDivider = "\\"
	}

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
			dir = newDir(dirNm, dir)
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
