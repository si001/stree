package tree_list

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/model"
	"si001/stree/tools/files"
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
		case tcell.Button2:
			_, y := ev.Position()
			if l.CheckIn(ev.Position()) && y != l.Min.Y && y != l.Max.Y {
				if mouseLastSelectedRow < 0 {
					mouseLastSelectedRow = y
					l.ScrollToMouse(ev.Position())
					mouseTagging = (*l.SelectedStringer()).(*model.FileInfo).IsTagged()
				} else {
					l.ScrollToMouse(ev.Position())
				}
				setTagFile(self, !mouseTagging, false)
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
		case tcell.KeyDown: //, "<MouseWheelDown>":
			l.ScrollDown()
		case tcell.KeyUp: //"<Up>", "<MouseWheelUp>":
			l.ScrollUp()
		case tcell.KeyEnter, tcell.KeyEsc:
			ViewModeChange(model.VM_TREEVIEW_FILES_1)
			l.ScrollTop()
			self.ListIsBranch = false
			node := self.Tree.SelectedNode()
			self.ShowDir(model.CurrentPath, node, false, false)
		case tcell.KeyHome:
			l.ScrollTop()
		case tcell.KeyEnd:
			l.ScrollBottom()
		case tcell.KeyPgDn, tcell.KeyRight:
			l.ScrollPageDown()
		case tcell.KeyLeft, tcell.KeyPgUp:
			l.ScrollPageUp()
		}
		if strings.HasPrefix(ev.Name(), "Shift+Rune[") {
			r := []rune(strings.ToLower(ev.Name()))[11]
			start := l.SelectedStringer()
			var old *fmt.Stringer = nil
			n := l.SelectedStringer()
			for ([]rune(strings.ToLower((*n).String()))[0] != r && l.SelectedStringer() != start) || old == nil {
				l.ScrollDown()
				old = n
				n = l.SelectedStringer()
				if n == old {
					l.ScrollTop()
					n = l.SelectedStringer()
				}
			}
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
	self.ShowDir(model.CurrentPath, self.Tree.SelectedNode(), false, false)
}

func (self *TreeAndList) ReSort() {
	self.List.SetOrderComparator(func(val1, val2 int) bool {
		fi1 := (*self.List.Rows[val1]).(*model.FileInfo)
		fi2 := (*self.List.Rows[val2]).(*model.FileInfo)
		var compResult bool = true
		switch self.OrderBy & model.OrderMask {
		case model.OrderByName:
			compResult = files.UpcaseIfWindows(fi1.Name) > files.UpcaseIfWindows(fi2.Name)
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
