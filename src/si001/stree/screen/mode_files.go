package screen

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/files"
	"si001/stree/model"
	"si001/stree/widgets"
	"strings"
)

func ModefilesDraw(s tcell.Screen, w, h int) {

	//FilesList1.Rows = listData[count%9:]
	FilesList1.SetRect(0, 1, w-VC_INFO_WIDTH+1, h-VC_BOTTOM_HEIGHT)
	//DriveInfo.SetRect(w-VC_INFO_WIDTH, 1, w-1, h-VC_BOTTOM_HEIGHT)
	//p.Text = fmt.Sprintf("%s : %d , %d:%d", p.Text[:20], count, w, h)

	//style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)
	//dt := time.Now()
	//HeadRight = dt.Format("02.01.2006 15:04:05")
	//ScreenPrintAt(s, 1, 0, style, HeadLeft+"   "+lastEvent)
	//ScreenPrintAt(s, w-22, 0, style, HeadRight)

	FilesList1.StyleNumber = FilesMode1
	FilesList1.Draw(s)
	ShowInfoBox(s, FilesMode1, FileMask1, FilePath1)
}

func ShowDir(s string, node *widgets.TreeNode, branch, actualise bool) {
	var rows []*fmt.Stringer
	if node.Value.(model.Directory).IsReadError() {
		errStyle := tcell.StyleDefault.Foreground(tcell.ColorRed)
		var row fmt.Stringer = model.InfoString{"Directory is not read, Read Error!", errStyle, errStyle}
		rows = append(rows, &row)
	} else if node.Value.(model.Directory).IsNotRead() {
		var row fmt.Stringer = model.InfoString{"Directory is not read", tcell.StyleDefault, tcell.StyleDefault}
		rows = append(rows, &row)
	} else if len(node.Value.(model.Directory).Files) == 0 {
		var row fmt.Stringer = model.InfoString{"No files", tcell.StyleDefault, tcell.StyleDefault}
		rows = append(rows, &row)
	} else {
		var nodes []*widgets.TreeNode
		nodes = append(nodes, node)

		rows = append(rows, GetFilesRecourse(nodes)...)
	}

	FilesList1.StyleNumber = FilesMode1
	FilesList1.Rows = rows
}

func GetFilesRecourse(nodes []*widgets.TreeNode) (rows []*fmt.Stringer) {
	for _, node := range nodes {
		files := files.NodeGetFiles(node)
		for _, item := range files {
			var row fmt.Stringer = *item
			rows = append(rows, &row)
		}
		rows = append(rows, GetFilesRecourse(node.Nodes)...)
	}
	return rows
}

func ModefilesPutEvent(event tcell.Event) bool {
	l := FilesList1
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
					ViewMode = VM_TREEVIEW_FILES_1
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
		case tcell.KeyDown: //, "<MouseWheelDown>":
			l.ScrollDown()
		case tcell.KeyUp: //"<Up>", "<MouseWheelUp>":
			l.ScrollUp()
		case tcell.KeyEnter, tcell.KeyEsc:
			ViewMode = VM_TREEVIEW_FILES_1
			l.ScrollTop()
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
		case "alt+rune[f]", "rune[f]":
			if FilesMode1 < 3 {
				FilesMode1++
			} else {
				FilesMode1 = 0
			}
		}
		lastEvent = ev.Name()
	}

	return true
}
