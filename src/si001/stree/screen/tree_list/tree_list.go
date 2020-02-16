package tree_list

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/model"
	"si001/stree/screen/botton_box"
	"si001/stree/tools/files"
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

func (self *TreeAndList) ShowDir(s string, node *widgets.TreeNode, taggedOnly, actualise bool) {
	var rows []*fmt.Stringer
	if !self.ListIsBranch && node.Value.(*model.Directory).IsReadError() {
		errStyle := tcell.StyleDefault.Foreground(tcell.ColorRed)
		var row fmt.Stringer = model.InfoString{" Directory is not read, Read Error!", errStyle, errStyle}
		rows = append(rows, &row)
	} else if !self.ListIsBranch && node.Value.(*model.Directory).IsNotRead() {
		var row fmt.Stringer = model.InfoString{" Directory is not read", tcell.StyleDefault, tcell.StyleDefault}
		rows = append(rows, &row)
	} else if !self.ListIsBranch && len(node.Value.(*model.Directory).Files) == 0 {
		var row fmt.Stringer = model.InfoString{" No files", tcell.StyleDefault, tcell.StyleDefault}
		rows = append(rows, &row)
	} else {
		var nodes []*widgets.TreeNode
		nodes = append(nodes, node)

		rows = append(rows, self.getFilesRecourse(nodes, taggedOnly, self.ListIsBranch)...)
		if len(rows) == 0 {
			var row fmt.Stringer = model.InfoString{" No files matching filespec", tcell.StyleDefault, tcell.StyleDefault}
			rows = append(rows, &row)
		}
	}

	self.List.StyleNumber = self.FileMode
	self.List.Rows = rows
	self.ReSort()
}

func (self *TreeAndList) getFilesRecourse(nodes []*widgets.TreeNode, taggedOnly, recourse bool) (rows []*fmt.Stringer) {
	for _, node := range nodes {
		files := files.NodeGetFiles(node)
		for _, item := range files {
			if (len(self.FileMask) == 0 || filterProcessed(item.Name, self.FileMask)) && (!taggedOnly || item.IsTagged()) {
				var row fmt.Stringer = item
				rows = append(rows, &row)
			}
		}
		if recourse {
			rows = append(rows, self.getFilesRecourse(node.Nodes, taggedOnly, true)...)
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

func filterProcessed(str, mask string) bool {
	if !strings.Contains(str, ".") {
		str += "."
	}
	s, p := []rune(str), []rune(mask)
	rs, rp := []rune{}, []rune{}
	for {
		if len(p) > 0 && p[0] == '*' {
			rs = s
			p = p[1:]
			rp = p
		} else if len(s) == 0 {
			return len(p) == 0
		} else if len(s) > 0 && len(p) > 0 && (s[0] == p[0] || p[0] == '?') {
			s = s[1:]
			p = p[1:]
		} else if len(rs) > 0 {
			rs = rs[1:]
			s = rs
			p = rp
		} else {
			return false
		}
	}
}

func (self *TreeAndList) Init() {
	self.actionsTree()
	self.actionsList()
}

func ViewModeChange(typ model.VmType) {
	model.ViewModeChng(typ)
	botton_box.NormalBottomBox()
}
