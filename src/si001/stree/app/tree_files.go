package app

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"os"
	"runtime"
	"si001/stree/model"
	"strings"
)

func buildTree(path string) *widgets.Tree {
	divider := "/"
	if runtime.GOOS == "windows" {
		divider = "\\"
	}
	ph := strings.Split(path, divider)
	var nodePath []*widgets.TreeNode
	var root widgets.TreeNode
	var dir *widgets.TreeNode
	//fmt.Println("path: ", path)
	for i, dirNm := range ph {
		if i == 0 {
			if dirNm == "" {
				dirNm = divider
			}
			root = *getRoot()
			dir = &root
			for _, rootElt := range dir.Nodes {
				if dirNm == rootElt.Value.String() {
					dir = rootElt
					break
				}
			}
			nodePath = append(nodePath, dir)
		} else {
			dir = newDir(dirNm, dir)
			nodePath = append(nodePath, dir)
		}
	}

	var l = widgets.NewTree()
	l.TextStyle = ui.NewStyle(ui.ColorYellow, ui.ColorBlack, ui.ModifierClear)
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

			//fmt.Println("counter: ", counter, l.SelectedRow, "|", l.SelectedNode().Value.String(), " - ", node.Value.String())
		}
		l.Expand()
		//fmt.Println(">: ", node.Value.String(), " - ", l.SelectedNode().Value.String())
	}
	//time.Sleep(time.Second * 9)

	return l
}

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

func getRoot() (r *widgets.TreeNode) {
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
