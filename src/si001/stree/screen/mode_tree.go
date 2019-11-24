package screen

import (
	ui "github.com/gizak/termui/v3"
	"github.com/nsf/termbox-go"
	"si001/stree/files"
	"si001/stree/model"
	"time"
)

func ModetreeDraw(w, h int) {

	divider := int(float32(h-2-VC_BOTTOM_HEIGHT)*0.7) + 2
	if divider > h-VC_BOTTOM_HEIGHT-2 {
		divider = h - VC_BOTTOM_HEIGHT
	}
	Tree1.SetRect(0-1, 1, w-VC_INFO_WIDTH+1, divider)
	FilesList1.SetRect(0, Tree1.GetRect().Max.Y-1, w-VC_INFO_WIDTH+1, h-VC_BOTTOM_HEIGHT)
	DriveInfo.SetRect(w-VC_INFO_WIDTH, 1, w, h-VC_BOTTOM_HEIGHT)

	dt := time.Now()
	HeadRight = dt.Format("02.01.2006 15:04:05")
	ScreenPrintAt(1, 0, termbox.ColorDefault, termbox.ColorDefault, HeadLeft+"   "+lastEvent.ID)
	ScreenPrintAt(w-22, 0, termbox.ColorDefault, termbox.ColorDefault, HeadRight)

	if divider == h-VC_BOTTOM_HEIGHT {
		ui.Render(Tree1, DriveInfo)
	} else {
		ui.Render(Tree1, DriveInfo, FilesList1)
	}
}

func ModetreePutEvent(event ui.Event) bool {
	l := Tree1
	switch event.ID {
	case "j", "<Down>", "<MouseWheelDown>":
		l.ScrollDown()
	case "k", "<Up>", "<MouseWheelUp>":
		l.ScrollUp()
	case "<C-d>":
		l.ScrollHalfPageDown()
	case "<C-u>":
		l.ScrollHalfPageUp()
	case "<C-f>", "<PageDown>":
		l.ScrollPageDown()
	case "<C-b>", "<PageUp>":
		l.ScrollPageUp()
	case "<Enter>":
		if !l.SelectedNode().Expanded && len(l.SelectedNode().Nodes) > 0 {
			l.Expand()
		} else {
			ViewMode = VM_FILELIST_1
			FilesList1.ScrollTop()
		}
	case "<Home>":
		l.ScrollTop()
	case "<End>":
		l.ScrollBottom()
	case "E":
		l.ExpandAll()
	case "C":
		l.CollapseAll()
	case "<Right>", "+":
		if !l.SelectedNode().Expanded && len(l.SelectedNode().Nodes) > 0 {
			l.Expand()
		} else {
			l.ScrollDown()
		}
	case "<Left>", "-":
		if l.SelectedNode().Expanded {
			l.Collapse()
		} else {
			l.ScrollUp()
		}

		//case "<Resize>":
		//	x, y := ui.TerminalDimensions()
		//	l.SetRect(0, 0, x, y)
	}

	newPath := files.TreeNodeToPath(l.SelectedNode())
	if model.CurrentPath != newPath {
		model.CurrentPath = newPath
		HeadLeft = model.CurrentPath

		ShowDir(model.CurrentPath, l.SelectedNode(), false)
	}

	lastEvent = event

	return true
}
