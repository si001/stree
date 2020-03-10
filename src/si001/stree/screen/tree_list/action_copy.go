package tree_list

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/model"
	"si001/stree/screen/botton_box/actions"
	"si001/stree/tools/files"
	"si001/stree/widgets"
	"strings"
	"time"
)

func (self *TreeAndList) actionCopy() {
	if fi, ok := (*self.List.SelectedStringer()).(*model.FileInfo); ok {
		path := files.TreeNodeToPath(fi.Owner)
		t1 := "       as: "
		t2 := "Copy file: %s"
		t3 := "`↑ history  `D`e`l  `E`s`c  `E`n`t`e`r"
		s1 := "       to: "
		s2 := "Copy file: %s"
		s3 := "`↑ history  `D`e`l  `E`s`c  `E`n`t`e`r  `F`2 folder"
		actions.RequestCopy(fi.Name, fi.Name, t1, t2, t3, s1, s2, s3, self.startSelectFolder, func(newPath, newName *string) {
			if newPath != nil && len(*newPath) > 0 {
				err, size := files.FileCopy(path+model.PathDivider+fi.Name, *newPath+model.PathDivider+*newName)
				if err == nil {
					fileInfo := model.FileInfo{
						Name:    *newName,
						Size:    size,
						ModTime: time.Now(),
						AttrB:   0,
						Owner:   nil,
					}
					node, _ := files.PutFileToPath(*newPath, &fileInfo, self.Tree)
					if strings.Compare(path, *newPath) == 0 {
						self.ShowDir("", node, false, false)
					}
				}
			}
		})
	}
}

func (self *TreeAndList) actionCopyTagged() {
	var fls []*model.FileInfo
	for _, tgd := range self.List.Rows {
		if fi, ok := (*tgd).(*model.FileInfo); ok && fi.IsTagged() {
			fls = append(fls, fi)
		}
	}
	if len(fls) > 0 {
		//path := files.TreeNodeToPath(fi.Owner)
		t1 := "       as: "
		t2 := "Copy %s tagged files"
		t3 := "`↑ history  `D`e`l  `E`s`c  `E`n`t`e`r"
		s1 := "       to: "
		s2 := "Copy %s files"
		s3 := "`↑ history  `D`e`l  `E`s`c  `E`n`t`e`r  `F`2 folder"
		actions.RequestCopy(fmt.Sprintf("%d", len(fls)), "*", t1, t2, t3, s1, s2, s3, self.startSelectFolder, func(newPath, newName *string) {
			if newPath != nil && len(*newPath) > 0 {
				for _, fi := range fls {
					path := files.TreeNodeToPath(fi.Owner)
					newName = &fi.Name
					err, size := files.FileCopy(path+model.PathDivider+fi.Name, *newPath+model.PathDivider+*newName)
					if err == nil {
						fileInfo := model.FileInfo{
							Name:    *newName,
							Size:    size,
							ModTime: time.Now(),
							AttrB:   0,
							Owner:   nil,
						}
						//node, _ :=
						files.PutFileToPath(*newPath, &fileInfo, self.Tree)
						//if strings.Compare(path, *newPath) == 0 {
						//	self.ShowDir("", node, false, false)
						//}
					}
				}
			}
		})
	}
}

func (self *TreeAndList) startSelectFolder(doResult func(result *string)) (func(s tcell.Screen), func(event tcell.Event) bool) {
	tree := widgets.NewTree()
	tree.SetNodes(self.Tree.GetRoot())
	tree.SelectedRow = self.Tree.SelectedRow
	y := self.Tree.SelectedRow - self.Tree.Dy()/2 + 1
	if y < 0 {
		y = 0
	}
	tree.TopRowSet(y)
	draw := func(s tcell.Screen) {
		w, h := s.Size()
		tree.SetRect(w/8, 1, w-model.VC_INFO_WIDTH, h-2)
		tree.Draw(s)
	}
	eventProcess := func(event tcell.Event) bool {
		l := tree
		node := l.SelectedNode()
		path := files.TreeNodeToPath(node)
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
							l.Expand()
						} else {
							doResult(&path)
						}
					}
					mouseClickTimerForDbl = ms
					if l.CheckIn(ev.Position()) {
						l.ScrollToMouse(ev.Position())
					}
				} else {
					if l.CheckIn(ev.Position()) || l.CheckIn(mouseLastEvent.Position()) {
						l.ScrollToMouse(ev.Position())
						toLastEvent = mouseLastEvent
					}
				}
			case tcell.Button2:
				doResult(nil)
			case tcell.ButtonNone:
				mouseLastSelectedRow = -1
			case tcell.WheelUp:
				l.ScrollUp()
			case tcell.WheelDown:
				l.ScrollDown()
			}
			mouseLastEvent = toLastEvent
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyDown: //"<Down>", "<MouseWheelDown>", "<Space>":
				if ev.Modifiers() == tcell.ModCtrl {
					l.ScrollScreenDown()
				} else {
					l.ScrollDown()
				}
			case tcell.KeyUp: //, "<MouseWheelUp>":
				if ev.Modifiers() == tcell.ModCtrl {
					l.ScrollScreenUp()
				} else {
					l.ScrollUp()
				}
			case tcell.KeyPgDn:
				l.ScrollPageDown()
			case tcell.KeyPgUp:
				l.ScrollPageUp()
			case tcell.KeyEnter:
				doResult(&path)
			case tcell.KeyEsc:
				doResult(nil)
			case tcell.KeyHome:
				l.ScrollTop()
			case tcell.KeyEnd:
				l.ScrollBottom()
			case tcell.KeyRight: //, "+":
				if node.Value.(*model.Directory).IsNotRead() {
					files.ReadDir(node)
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
		return true
	}
	return draw, eventProcess
}
