package tree_list

import (
	"si001/stree/model"
	"si001/stree/screen/botton_box"
	"si001/stree/screen/botton_box/actions"
	"si001/stree/tools/files"
)

func (self *TreeAndList) actionsTree() {
	var actions = []botton_box.Action{
		botton_box.Action{
			ActName: "`Tag",
			ActKey:  "rune[t]",
			Callback: func() {
				setTagDir(self, true, true)
			},
		},
		botton_box.Action{
			ActName: "`Untag",
			ActKey:  "rune[u]",
			Callback: func() {
				setTagDir(self, false, true)
			},
		},
		botton_box.Action{
			ActName: "`Branch",
			ActKey:  "rune[b]",
			Callback: func() {
				ViewModeChange(model.VM_FILELIST_1)
				self.ListIsBranch = true
				self.ShowDir(model.CurrentPath, self.Tree.SelectedNode(), false, false)
				self.List.ScrollTop()
			},
		},
		botton_box.Action{
			ActName: "`Ctl+`Branch",
			ActKey:  "ctrl+b",
			Callback: func() {
				ViewModeChange(model.VM_FILELIST_1)
				self.ListIsBranch = true
				self.ShowDir(model.CurrentPath, self.Tree.SelectedNode(), true, false)
				self.List.ScrollTop()
			},
		},
		botton_box.Action{
			ActName: "`Filespec",
			ActKey:  "rune[f]",
			Callback: func() {
				actions.RequestFileMask(self.FileMask, self.setFileMask)
			},
		},
		botton_box.Action{
			ActName: "`Alt+`File display",
			ActKey:  "alt+rune[f]",
			Callback: func() {
				self.processNextFileMode()
			},
		},
		botton_box.Action{
			ActName: "`Log path",
			ActKey:  "rune[l]",
			Callback: func() {
				self.actionLog()
			},
		},
		botton_box.Action{
			ActName: "`Make dir",
			ActKey:  "rune[m]",
			Callback: func() {
				self.actionMkDir()
			},
		},
		botton_box.Action{
			ActName: "`Delete dir",
			ActKey:  "rune[d]",
			Callback: func() {
				self.actionRmDir()
			},
		},
		botton_box.Action{
			ActName: "`Rename dir",
			ActKey:  "rune[r]",
			Callback: func() {
				self.actionRename(true)
			},
		},
		botton_box.Action{ // new line
			ActName: "",
			ActKey:  "",
		},
		botton_box.Action{
			ActName: "`F`3 Reload dir",
			ActKey:  "f3",
			Callback: func() {
				node := self.Tree.SelectedNode()
				files.CloseDir(node)
				files.ReadDir(node)
				self.ShowDir(model.CurrentPath, node, false, false)
				self.Tree.Expand()
			},
		},
		botton_box.Action{
			ActName:      "`* read branch",
			ActKey:       "rune[*]",
			PrefixPermit: "shift+",
			Callback: func() {
				files.RefreshTreeNode(self.Tree.SelectedNode())
				self.Tree.ExpandRecursive()
			},
		},
		botton_box.Action{
			ActName:      "",
			ActKey:       "rune[+]",
			PrefixPermit: "shift+",
			Callback: func() {
				l := self.Tree
				node := l.SelectedNode()
				if node.Value.(*model.Directory).IsNotRead() {
					files.ReadDir(node)
					self.ShowDir(model.CurrentPath, node, false, false)
					l.Expand()
				} else if !l.SelectedNode().Expanded && len(l.SelectedNode().Nodes) > 0 {
					l.Expand()
				}
			},
		},
		botton_box.Action{
			ActName:      "",
			ActKey:       "rune[-]",
			PrefixPermit: "shift+",
			Callback: func() {
				l := self.Tree
				node := l.SelectedNode()
				files.CloseDir(node)
				if l.SelectedNode().Expanded {
					l.Collapse()
				}
				self.ShowDir(model.CurrentPath, l.SelectedNode(), false, false)
			},
		},
		botton_box.Action{
			ActName: "`F`5 one level",
			ActKey:  "f5",
			Callback: func() {
				self.pressedF5()
			},
		},
		botton_box.Action{
			ActName: "`F`6 expand/collapse",
			ActKey:  "f6",
			Callback: func() {
				self.pressedF6()
			},
		},
		botton_box.Action{
			ActName: "",
			ActKey:  "rune[ ]",
			Callback: func() {
				self.Tree.ScrollDown()
			},
		},
		botton_box.Action{
			ActName: "`Quit",
			ActKey:  "rune[q]",
			Callback: func() {
				actionQuit()
			},
		},
	}
	botton_box.SetTreeActions(actions)
}
