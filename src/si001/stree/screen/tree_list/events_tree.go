package tree_list

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/model"
	"si001/stree/tools/files"
	"si001/stree/widgets"
	"strings"
)

var mouseLastEvent *tcell.EventMouse
var mouseClickTimerForDbl int64
var mouseLastSelectedRow int
var mouseTagging bool

func (self *TreeAndList) PutEventTreeList(event tcell.Event) bool {
	l := self.Tree
	node := l.SelectedNode()
	switch ev := event.(type) {
	case *tcell.EventResize:
	case *tcell.EventMouse:
		toLastEvent := ev
		switch ev.Buttons() {
		case tcell.Button1:
			if mouseLastEvent != nil && ev.Buttons()&mouseLastEvent.Buttons() == 0 {
				var ms int64 = ev.When().UnixNano() / 1000000
				if ms-mouseClickTimerForDbl < 400 {
					if node.Value.(*model.Directory).IsNotRead() {
						files.ReadDir(node)
						self.ShowDir(model.CurrentPath, node, false, false)
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
						self.ShowDir(model.CurrentPath, node, false, false)
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
		case tcell.Button2:
			_, y := ev.Position()
			if l.CheckIn(ev.Position()) && y != l.Min.Y && y != l.Max.Y {
				if mouseLastSelectedRow < 0 {
					mouseLastSelectedRow = y
					l.ScrollToMouse(ev.Position())
					mouseTagging = checkIsTagged(l.SelectedNode())
				} else {
					l.ScrollToMouse(ev.Position())
				}
				setTagDir(self, !mouseTagging, false)
			}
		case tcell.ButtonNone:
			mouseLastSelectedRow = -1
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
		case tcell.KeyEnter:
			if node.Value.(*model.Directory).IsNotRead() {
				files.ReadDir(node)
				self.ShowDir(model.CurrentPath, node, false, false)
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
				self.ShowDir(model.CurrentPath, node, false, false)
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
				node := l.SelectedNode()
				parent := node.Value.(*model.Directory).Owner
				for node.Value.(*model.Directory).Owner != nil && parent != node {
					l.ScrollUp()
					node = l.SelectedNode()
				}
			}
		}
		if strings.HasPrefix(ev.Name(), "Shift+Rune[") {
			r := []rune(strings.ToLower(ev.Name()))[11]
			start := l.SelectedNode()
			n := l.SelectedNode()
			var old *widgets.TreeNode = nil
			for ([]rune(strings.ToLower(n.Value.String()))[0] != r && l.SelectedNode() != start) || old == nil {
				l.ScrollDown()
				old = n
				n = l.SelectedNode()
				if n == old {
					l.ScrollTop()
					n = l.SelectedNode()
				}
			}
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
		self.ShowDir(model.CurrentPath, self.Tree.SelectedNode(), false, false)
	}
}

func (self *TreeAndList) pressedF6() {
	node := self.Tree.SelectedNode()
	if node.Expanded {
		self.Tree.Collapse()
	} else {
		self.Tree.ExpandRecursive()
	}
}

func (self *TreeAndList) pressedF5() {
	node := self.Tree.SelectedNode()
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
		self.Tree.CollapseOneLevel()
	} else {
		self.Tree.ExpandRecursive()
	}
}
