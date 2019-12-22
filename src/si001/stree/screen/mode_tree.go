package screen

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/files"
	"si001/stree/model"
	"si001/stree/widgets"
	"strings"
)

var mouseLastEvent *tcell.EventMouse
var mouseClickTimerForDbl int64

var tmpValue string = ""

func ModetreeDraw(s tcell.Screen, w, h int) {

	if model.Divider < 2 {
		model.Divider = 1
	}
	Tree1.SetRect(0, 1, w-VC_INFO_WIDTH, h-model.Divider)
	FilesList1.SetRect(0, Tree1.GetRect().Max.Y, w-VC_INFO_WIDTH+1, h-VC_BOTTOM_HEIGHT)
	//DriveInfo.SetRect(w-VC_INFO_WIDTH, 1, w-1, h-VC_BOTTOM_HEIGHT)

	if model.Divider <= 2 {
		Tree1.Draw(s)
	} else {
		FilesList1.StyleNumber = FilesMode1
		FilesList1.Draw(s)
		Tree1.Draw(s)
	}
	ShowInfoBox(s, FilesMode1, FileMask1, FilePath1)
}

func ModetreePutEvent(event tcell.Event) bool {
	l := Tree1
	node := l.SelectedNode()
	switch ev := event.(type) {
	case *tcell.EventResize:
		//x, y := ev.Size()
		//l.SetRect(0, 0, x, y)
	case *tcell.EventMouse:
		toLastEvent := ev
		switch ev.Buttons() {
		case tcell.Button1:
			if ev.Buttons()&mouseLastEvent.Buttons() == 0 {
				var ms int64 = ev.When().UnixNano() / 1000000
				if ms-mouseClickTimerForDbl < 400 {
					if node.Value.(model.Directory).IsNotRead() {
						files.ReadDir(node)
						ShowDir(model.CurrentPath, node, false, false)
						l.Expand()
					} else if len(node.Value.(model.Directory).Files) > 0 {
						ViewMode = VM_FILELIST_1
						FilesList1.ScrollTop()
					}
				}
				mouseClickTimerForDbl = ms
				if l.CheckIn(ev.Position()) {
					l.ScrollToMouse(ev.Position())
				} else if FilesList1.CheckIn(ev.Position()) {
					if node.Value.(model.Directory).IsNotRead() {
						files.ReadDir(node)
						ShowDir(model.CurrentPath, node, false, false)
						l.Expand()
					}
					ViewMode = VM_FILELIST_1
					FilesList1.ScrollTop()
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
		lastEvent = fmt.Sprintf("%d:%d / %s : %s", x, y, string(ev.Buttons()), ev.Modifiers())
		mouseLastEvent = toLastEvent
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyDown: //"<Down>", "<MouseWheelDown>", "<Space>":
			if ev.Modifiers() == tcell.ModAlt {
				if model.Divider > 1 {
					model.Divider--
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
				if model.Divider < model.ScreenHeight {
					model.Divider++
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
			ShowDir(model.CurrentPath, node, false, false)
			l.Expand()
		case tcell.KeyF5:
			pressedF5(l)
		case tcell.KeyF6:
			pressedF6(l)
		case tcell.KeyEnter:
			if node.Value.(model.Directory).IsNotRead() {
				files.ReadDir(node)
				ShowDir(model.CurrentPath, node, false, false)
				l.Expand()
			} else if !node.Expanded && len(node.Nodes) > 0 {
				l.Expand()
			} else {
				ViewMode = VM_FILELIST_1
				FilesList1.ScrollTop()
			}
		case tcell.KeyHome:
			l.ScrollTop()
		case tcell.KeyEnd:
			l.ScrollBottom()
		case tcell.KeyRight: //, "+":
			if node.Value.(model.Directory).IsNotRead() {
				files.ReadDir(node)
				ShowDir(model.CurrentPath, node, false, false)
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
			if node.Value.(model.Directory).IsNotRead() {
				files.ReadDir(node)
				ShowDir(model.CurrentPath, node, false, false)
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
				ShowDir(model.CurrentPath, l.SelectedNode(), false, false)
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
		case "alt+rune[f]", "rune[f]":
			if FilesMode1 < 3 {
				FilesMode1++
			} else {
				FilesMode1 = 0
			}
		case "rune[b]":
			ViewMode = VM_FILELIST_1
			FileList1_IsBranch = true
			ShowDir(model.CurrentPath, l.SelectedNode(), true, false)
			FilesList1.ScrollTop()
		}
		lastEvent = ev.Name()
	}

	newPath := files.TreeNodeToPath(l.SelectedNode())
	if model.CurrentPath != newPath {
		model.CurrentPath = newPath

		ShowDir(model.CurrentPath, l.SelectedNode(), false, false)
		FileList1_IsBranch = false
	}

	return true
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
