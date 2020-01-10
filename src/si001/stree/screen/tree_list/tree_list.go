package tree_list

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/files"
	"si001/stree/model"
	"si001/stree/screen/botton_box"
	"si001/stree/widgets"
	"strings"
)

type TreeAndList struct {
	Tree         *widgets.Tree
	List         *widgets.List
	FileMode     int
	FileMask     string
	OrderBy      byte
	Divider      int
	ListIsBranch bool
	CurrentPath  string
}

func (self *TreeAndList) ShowDir(s string, node *widgets.TreeNode, actualise bool) {
	var rows []*fmt.Stringer
	if !self.ListIsBranch && node.Value.(*model.Directory).IsReadError() {
		errStyle := tcell.StyleDefault.Foreground(tcell.ColorRed)
		var row fmt.Stringer = model.InfoString{"Directory is not read, Read Error!", errStyle, errStyle}
		rows = append(rows, &row)
	} else if !self.ListIsBranch && node.Value.(*model.Directory).IsNotRead() {
		var row fmt.Stringer = model.InfoString{"Directory is not read", tcell.StyleDefault, tcell.StyleDefault}
		rows = append(rows, &row)
	} else if !self.ListIsBranch && len(node.Value.(*model.Directory).Files) == 0 {
		var row fmt.Stringer = model.InfoString{"No files", tcell.StyleDefault, tcell.StyleDefault}
		rows = append(rows, &row)
	} else {
		var nodes []*widgets.TreeNode
		nodes = append(nodes, node)

		rows = append(rows, self.getFilesRecourse(nodes, self.ListIsBranch)...)
	}

	self.List.StyleNumber = self.FileMode
	self.List.Rows = rows
	self.ReSort()
}

func (self *TreeAndList) getFilesRecourse(nodes []*widgets.TreeNode, recourse bool) (rows []*fmt.Stringer) {
	for _, node := range nodes {
		files := files.NodeGetFiles(node)
		for _, item := range files {
			if len(self.FileMask) == 0 || filterProcessed(item.Name, self.FileMask) {
				var row fmt.Stringer = item
				rows = append(rows, &row)
			}
		}
		if recourse {
			rows = append(rows, self.getFilesRecourse(node.Nodes, true)...)
		}
	}
	return rows
}

func (self *TreeAndList) processNextFileMode() {
	if self.FileMode < 3 {
		self.FileMode++
	} else {
		self.FileMode = 0
	}
}

func filterProcessed(s, f string) bool {
	if f == "*.*" || f == "*" {
		return true
	}
	res := strings.Contains(s, f)
	return res
}

func (self *TreeAndList) Init() {
	self.actionsTree()
	self.actionsList()
}

func ViewModeChange(typ model.VmType) {
	model.ViewModeChng(typ)
	botton_box.NormalBottomBox()
}
