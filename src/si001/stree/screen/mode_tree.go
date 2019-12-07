package screen

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/files"
	"si001/stree/model"
	"si001/stree/widgets"
	"si001/stree/widgets/stuff"
	"time"
)

func ModetreeDraw(s tcell.Screen, w, h int) {

	if model.Divider > h-VC_BOTTOM_HEIGHT-1 {
		model.Divider = h - VC_BOTTOM_HEIGHT
	}
	Tree1.SetRect(0, 1, w-VC_INFO_WIDTH, model.Divider)
	FilesList1.SetRect(0, Tree1.GetRect().Max.Y-1, w-VC_INFO_WIDTH+1, h-VC_BOTTOM_HEIGHT)
	DriveInfo.SetRect(w-VC_INFO_WIDTH, 1, w-1, h-VC_BOTTOM_HEIGHT)

	if model.Divider == h-VC_BOTTOM_HEIGHT {
		Tree1.Draw(s)
		DriveInfo.Draw(s)
	} else {
		Tree1.Draw(s)
		FilesList1.Draw(s)
		DriveInfo.Draw(s)
	}

	style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault).Bold(true)
	dt := time.Now()
	HeadRight = dt.Format("02.01.2006 15:04:05")
	stuff.ScreenPrintAt(s, 1, 0, style, HeadLeft+"   "+lastEvent)
	stuff.ScreenPrintAt(s, w-22, 0, style, HeadRight)
}

func ModetreePutEvent(event tcell.Event) bool {
	l := Tree1
	switch ev := event.(type) {
	case *tcell.EventResize:
		//x, y := ev.Size()
		//l.SetRect(0, 0, x, y)
	case *tcell.EventMouse:
		switch ev.Buttons() {
		case tcell.WheelUp:
			l.ScrollUp()
		case tcell.WheelDown:
			l.ScrollDown()
		}
		x, y := ev.Position()
		lastEvent = fmt.Sprintf("%d:%d / %s : %s", x, y, string(ev.Buttons()), ev.Modifiers())
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyDown, tcell.KeyCtrlSpace: //"<Down>", "<MouseWheelDown>", "<Space>":
			if ev.Modifiers() == tcell.ModCtrl {
				//w, h := tcell.Screen.Size()
				//if model.Divider < h-1 {
				model.Divider++
				//}
			} else {
				l.ScrollDown()
			}
		case tcell.KeyUp: //, "<MouseWheelUp>":
			if ev.Modifiers() == tcell.ModCtrl {
				if model.Divider > 0 {
					model.Divider--
				}
			} else {
				l.ScrollUp()
			}
		//case "<C-d>":
		//	l.ScrollHalfPageDown()
		//case "<C-u>":
		//	l.ScrollHalfPageUp()
		case tcell.KeyPgDn:
			l.ScrollPageDown()
		case tcell.KeyPgUp:
			l.ScrollPageUp()
		case tcell.KeyF3:
			node := l.SelectedNode()
			files.ReadDir(node)
			ShowDir(model.CurrentPath, node, false)
			l.Expand()
		case tcell.KeyF5:
			pressedF5(l)
		case tcell.KeyF6:
			pressedF6(l)
		case tcell.KeyEnter:
			node := l.SelectedNode()
			if node.Value.(model.Directory).Attr == model.ATTR_NOTREAD {
				files.ReadDir(node)
				ShowDir(model.CurrentPath, node, false)
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
			if !l.SelectedNode().Expanded && len(l.SelectedNode().Nodes) > 0 {
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

		switch ev.Name() {
		case "Rune[+]", "Shift+Rune[+]":
			if !l.SelectedNode().Expanded && len(l.SelectedNode().Nodes) > 0 {
				l.Expand()
			}
		case "Rune[-]", "Shift+Rune[-]":
			if l.SelectedNode().Expanded {
				node := l.SelectedNode()
				node.Nodes = nil
				dir := node.Value.(model.Directory)
				dir.Attr = model.ATTR_NOTREAD
				dir.Files = nil
				node.Value = dir
				l.Collapse()
				ShowDir(model.CurrentPath, l.SelectedNode(), false)
			}
		case "Rune[*]", "Shift+Rune[*]":
			RefreshTreeNodeRecource(l.SelectedNode())
			l.ExpandRecursive()
		}
		lastEvent = ev.Name()
	}

	newPath := files.TreeNodeToPath(l.SelectedNode())
	if model.CurrentPath != newPath {
		model.CurrentPath = newPath
		HeadLeft = model.CurrentPath

		ShowDir(model.CurrentPath, l.SelectedNode(), false)
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
