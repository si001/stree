package screen

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/nsf/termbox-go"
	"si001/stree/files"
	"si001/stree/model"
	"time"
)

func ModefilesDraw(w, h int) {

	//FilesList1.Rows = listData[count%9:]
	FilesList1.SetRect(0, 1, w-24, h-1)
	DriveInfo.SetRect(w-25, 1, w, h-1)
	//p.Text = fmt.Sprintf("%s : %d , %d:%d", p.Text[:20], count, w, h)

	dt := time.Now()
	HeadRight = dt.Format("02.01.2006 15:04:05")
	ScreenPrintAt(1, 0, termbox.ColorDefault, termbox.ColorDefault, HeadLeft+"   "+lastEvent.ID)
	ScreenPrintAt(w-22, 0, termbox.ColorDefault, termbox.ColorDefault, HeadRight)

	ui.Render(Tree1, DriveInfo, FilesList1)
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

func ModefilesPutEvent(event ui.Event) bool {
	l := FilesList1
	switch event.ID {
	case "<Down>", "<MouseWheelDown>":
		l.ScrollDown()
	case "<Up>", "<MouseWheelUp>":
		l.ScrollUp()
	case "<Enter>", "<Escape>":
		ViewMode = VM_TREEVIEW_FILES_1
		l.ScrollTop()
	case "<Home>":
		l.ScrollTop()
	case "<End>":
		l.ScrollBottom()
	case "<Right>":
		l.ScrollPageDown()
	case "<Left>":
		l.ScrollPageUp()

		//case "<Resize>":
		//	x, y := ui.TerminalDimensions()
		//	l.SetRect(0, 0, x, y)
	}

	lastEvent = event

	return true
}
