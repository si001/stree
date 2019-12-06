package screen

import (
	"github.com/gdamore/tcell"
	//"github.com/gdamore/tcell/encoding"

	"si001/stree/files"
	"si001/stree/model"
	"si001/stree/widgets"
	"time"
)

func ModefilesDraw(s tcell.Screen, w, h int) {

	//FilesList1.Rows = listData[count%9:]
	FilesList1.SetRect(0, 1, w-VC_INFO_WIDTH+1, h-VC_BOTTOM_HEIGHT)
	DriveInfo.SetRect(w-VC_INFO_WIDTH, 1, w-1, h-VC_BOTTOM_HEIGHT)
	//p.Text = fmt.Sprintf("%s : %d , %d:%d", p.Text[:20], count, w, h)

	style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)
	dt := time.Now()
	HeadRight = dt.Format("02.01.2006 15:04:05")
	ScreenPrintAt(s, 1, 0, style, HeadLeft+"   "+lastEvent)
	ScreenPrintAt(s, w-22, 0, style, HeadRight)

	FilesList1.Draw(s)
	DriveInfo.Draw(s)
}

func ShowDir(s string, node *widgets.TreeNode, actualise bool) {
	var rows []string
	if node.Value.(model.Directory).Attr == model.ATTR_NOTREAD {
		rows = append(rows, "Directory is not read")
	} else {
		files := files.NodeGetFiles(node)

		for _, item := range files {
			rows = append(rows, item.Name)
		}
	}

	FilesList1.Rows = rows
}

func ModefilesPutEvent(event tcell.Event) bool {
	l := FilesList1
	switch ev := event.(type) {
	case *tcell.EventResize:
		//s.Sync()
		//st := tcell.StyleDefault.Background(tcell.ColorRed)
		//s.SetContent(w-1, h-1, 'R', nil, st)
	case *tcell.EventMouse:
		//x, y := ev.Position()
		//button := ev.Buttons()
		//s.SetContent(w-1, h-1, 'R', nil, st)
		//processEvent(event)
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
		case tcell.KeyRight:
			l.ScrollPageDown()
		case tcell.KeyLeft:
			l.ScrollPageUp()

			//case "<Resize>":
			//	x, y := ui.TerminalDimensions()
			//	l.SetRect(0, 0, x, y)
		}
		lastEvent = ev.Name()
	}

	return true
}
