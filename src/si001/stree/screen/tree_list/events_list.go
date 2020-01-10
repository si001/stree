package tree_list

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/files"
	"si001/stree/model"
	"si001/stree/screen/botton_box"
	"strings"
)

func (self *TreeAndList) PutEventList(event tcell.Event) bool {
	l := self.List
	switch ev := event.(type) {
	case *tcell.EventResize:
		//s.Sync()
		//st := tcell.StyleDefault.Background(tcell.ColorRed)
		//s.SetContent(w-1, h-1, 'R', nil, st)
	case *tcell.EventMouse:
		toLastEvent := ev
		switch ev.Buttons() {
		case tcell.Button1:
			if ev.Buttons()&mouseLastEvent.Buttons() == 0 {
				var ms int64 = ev.When().UnixNano() / 1000000
				if ms-mouseClickTimerForDbl < 400 {
					//if (node.Value.(model.Directory).Attr & model.ATTR_NOTREAD) > 0 {
					//	files.ReadDir(node)
					//	ShowDir(model.CurrentPath, node, false)
					//	l.Expand()
					//} else if len(node.Value.(model.Directory).Files) > 0 {
					//	ViewMode = VM_FILELIST_1
					//	FilesList1.ScrollTop()
					//}
				}
				mouseClickTimerForDbl = ms
				_, y := ev.Position()
				if l.CheckIn(ev.Position()) && y != l.Min.Y && y != l.Max.Y {
					l.ScrollToMouse(ev.Position())
				} else {
					ViewModeChange(model.VM_TREEVIEW_FILES_1)
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
		case tcell.KeyDown: //, "<MouseWheelDown>":
			l.ScrollDown()
		case tcell.KeyUp: //"<Up>", "<MouseWheelUp>":
			l.ScrollUp()
		case tcell.KeyEnter, tcell.KeyEsc:
			ViewModeChange(model.VM_TREEVIEW_FILES_1)
			l.ScrollTop()
			self.ListIsBranch = false
			node := self.Tree.SelectedNode()
			self.ShowDir(model.CurrentPath, node, false)
		case tcell.KeyHome:
			l.ScrollTop()
		case tcell.KeyEnd:
			l.ScrollBottom()
		case tcell.KeyPgDn, tcell.KeyRight:
			l.ScrollPageDown()
		case tcell.KeyLeft, tcell.KeyPgUp:
			l.ScrollPageUp()

			//case "<Resize>":
			//	x, y := ui.TerminalDimensions()
			//	l.SetRect(0, 0, x, y)
		}
		switch strings.ToLower(ev.Name()) {
		case "alt+rune[s]", "rune[s]":
			botton_box.RequestOrderBy(self.OrderBy, self.setOrderBy)
		case "alt+rune[f]", "rune[d]":
			self.processNextFileMode()
		case "rune[f]":
			botton_box.RequestFileMask(self.FileMask, self.setFileMask)
		}
		model.LastEvent = ev.Name()
	}

	return true
}

func (self *TreeAndList) setOrderBy(orderBy byte) {
	if orderBy == model.OrderByUndefined {
		return
	}
	self.OrderBy = orderBy
	self.ReSort()
}

func (self *TreeAndList) setFileMask(mask string) {
	if len(mask) == 0 {
		mask = "*"
	}
	self.FileMask = mask
	self.ShowDir(model.CurrentPath, self.Tree.SelectedNode(), false)
}

func (self *TreeAndList) ReSort() {
	self.List.SetOrderComparator(func(val1, val2 int) bool {
		fi1 := (*self.List.Rows[val1]).(*model.FileInfo)
		fi2 := (*self.List.Rows[val2]).(*model.FileInfo)
		var compResult bool = true
		switch self.OrderBy & model.OrderMask {
		case model.OrderByName:
			compResult = fi1.Name > fi2.Name
		case model.OrderByExt:
			compResult = fi1.Name > fi2.Name
		case model.OrderBySize:
			compResult = fi1.Size > fi2.Size
		case model.OrderByDate:
			compResult = fi1.ModTime.Unix() > fi2.ModTime.Unix()
		}
		if self.OrderBy&model.OrderByPath > 0 {
			p1 := files.TreeNodeToPath(fi1.Owner)
			p2 := files.TreeNodeToPath(fi2.Owner)
			if p1 != p2 {
				compResult = p1 > p2
			}
		}
		if self.OrderBy&model.OrderAcs > 0 {
			compResult = !compResult
		}
		return compResult
	})
}
