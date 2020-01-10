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

var mouseLastEvent *tcell.EventMouse
var mouseClickTimerForDbl int64

func (self *TreeAndList) PutEventTreeList(event tcell.Event) bool {
	l := self.Tree
	node := l.SelectedNode()
	switch ev := event.(type) {
	case *tcell.EventResize:
	case *tcell.EventMouse:
		toLastEvent := ev
		switch ev.Buttons() {
		case tcell.Button1:
			if ev.Buttons()&mouseLastEvent.Buttons() == 0 {
				var ms int64 = ev.When().UnixNano() / 1000000
				if ms-mouseClickTimerForDbl < 400 {
					if node.Value.(*model.Directory).IsNotRead() {
						files.ReadDir(node)
						self.ShowDir(model.CurrentPath, node, false)
						l.Expand()
					} else if len(node.Value.(*model.Directory).Files) > 0 {
						ViewModeChange(model.VM_FILELIST_1)
						self.List.ScrollTop()
					}
				}
				mouseClickTimerForDbl = ms
				if l.CheckIn(ev.Position()) {
					l.ScrollToMouse(ev.Position())
				} else if self.List.CheckIn(ev.Position()) {
					if node.Value.(*model.Directory).IsNotRead() {
						files.ReadDir(node)
						self.ShowDir(model.CurrentPath, node, false)
						l.Expand()
					}
					ViewModeChange(model.VM_FILELIST_1)
					self.List.ScrollTop()
				}
			} else {
				if l.CheckIn(ev.Position()) || l.CheckIn(mouseLastEvent.Position()) {
					l.ScrollToMouse(ev.Position())
					toLastEvent = mouseLastEvent
				}
			}
		case tcell.WheelUp:
			l.ScrollUp()
		case tcell.WheelDown:
			l.ScrollDown()
		}
		x, y := ev.Position()
		model.LastEvent = fmt.Sprintf("%d:%d / %s : %s", x, y, string(ev.Buttons()), ev.Modifiers())
		mouseLastEvent = toLastEvent
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyDown: //"<Down>", "<MouseWheelDown>", "<Space>":
			if ev.Modifiers() == tcell.ModAlt {
				if self.Divider > 1 {
					self.Divider--
				}
			} else {
				if ev.Modifiers() == tcell.ModCtrl {
					l.ScrollScreenDown()
				} else {
					l.ScrollDown()
				}
			}
		case tcell.KeyUp: //, "<MouseWheelUp>":
			if ev.Modifiers() == tcell.ModAlt {
				if self.Divider < model.ScreenHeight {
					self.Divider++
				}
			} else {
				if ev.Modifiers() == tcell.ModCtrl {
					l.ScrollScreenUp()
				} else {
					l.ScrollUp()
				}
			}
		case tcell.KeyPgDn:
			l.ScrollPageDown()
		case tcell.KeyPgUp:
			l.ScrollPageUp()
		case tcell.KeyF3:
			node := l.SelectedNode()
			files.ReadDir(node)
			self.ShowDir(model.CurrentPath, node, false)
			l.Expand()
		case tcell.KeyF5:
			pressedF5(l)
		case tcell.KeyF6:
			pressedF6(l)
		case tcell.KeyEnter:
			if node.Value.(*model.Directory).IsNotRead() {
				files.ReadDir(node)
				self.ShowDir(model.CurrentPath, node, false)
				l.Expand()
			} else if !node.Expanded && len(node.Nodes) > 0 {
				l.Expand()
			} else {
				ViewModeChange(model.VM_FILELIST_1)
				self.List.ScrollTop()
			}
		case tcell.KeyHome:
			l.ScrollTop()
		case tcell.KeyEnd:
			l.ScrollBottom()
		case tcell.KeyRight: //, "+":
			if node.Value.(*model.Directory).IsNotRead() {
				files.ReadDir(node)
				self.ShowDir(model.CurrentPath, node, false)
				l.Expand()
			} else if !l.SelectedNode().Expanded && len(l.SelectedNode().Nodes) > 0 {
				l.Expand()
			} else {
				l.ScrollDown()
			}
		case tcell.KeyLeft: // "-":
			if l.SelectedNode().Expanded {
				l.Collapse()
			} else {
				l.ScrollUp()
			}
		}

		switch strings.ToLower(ev.Name()) {
		case "rune[+]", "shift+rune[+]":
			if node.Value.(*model.Directory).IsNotRead() {
				files.ReadDir(node)
				self.ShowDir(model.CurrentPath, node, false)
				l.Expand()
			} else if !l.SelectedNode().Expanded && len(l.SelectedNode().Nodes) > 0 {
				l.Expand()
			}
		case "rune[-]", "shift+rune[-]":
			if l.SelectedNode().Expanded {
				node := l.SelectedNode()
				node.Nodes = nil
				dir := node.Value.(model.Directory)
				dir.AttrB = model.ATTR_NOTREAD | model.ATTR_DIR
				dir.Files = nil
				node.Value = dir
				l.Collapse()
				self.ShowDir(model.CurrentPath, l.SelectedNode(), false)
			}
		case "rune[*]", "shift+rune[*]":
			RefreshTreeNodeRecource(l.SelectedNode())
			l.ExpandRecursive()
		case "rune[ ]":
			//if (node.Value.(model.Directory).Attr & model.ATTR_NOTREAD) > 0 {
			//	files.ReadDir(node)
			//	ShowDir(model.CurrentPath, node, false)
			//	l.Expand()
			//} else if !l.SelectedNode().Expanded && len(l.SelectedNode().Nodes) > 0 {
			//	l.Expand()
			//} else {
			l.ScrollDown()
		//}
		case "alt+rune[f]", "rune[d]":
			self.processNextFileMode()
		case "rune[f]":
			botton_box.RequestFileMask(self.FileMask, self.setFileMask)
		case "rune[b]":
			ViewModeChange(model.VM_FILELIST_1)
			self.ListIsBranch = true
			self.ShowDir(model.CurrentPath, l.SelectedNode(), false)
			self.List.ScrollTop()
		}
		model.LastEvent = ev.Name()
	}

	self.pathCheck()

	return true
}
func (self *TreeAndList) pathCheck() {
	newPath := files.TreeNodeToPath(self.Tree.SelectedNode())
	if model.CurrentPath != newPath {
		model.CurrentPath = newPath

		self.ListIsBranch = false
		self.ShowDir(model.CurrentPath, self.Tree.SelectedNode(), false)
	}
}

func RefreshTreeNodeRecource(node *widgets.TreeNode) {
	files.ReadDir(node)
	for _, n := range node.Nodes {
		RefreshTreeNodeRecource(n)
	}
}

func pressedF6(tree *widgets.Tree) {
	node := tree.SelectedNode()
	if node.Expanded {
		tree.Collapse()
	} else {
		tree.ExpandRecursive()
	}
}

func pressedF5(tree *widgets.Tree) {
	node := tree.SelectedNode()
	allExpanded := true && node.Expanded
	if allExpanded {
		for _, n := range node.Nodes {
			if !n.Expanded {
				allExpanded = false
				break
			}
		}
	}
	if allExpanded {
		tree.CollapseOneLevel()
	} else {
		tree.ExpandRecursive()
	}
}
