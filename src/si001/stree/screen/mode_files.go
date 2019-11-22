package screen

import (
	ui "github.com/gizak/termui/v3"
	"github.com/nsf/termbox-go"
	"time"
)

func ModefilesDraw(w, h int) {

	//FilesList1.Rows = listData[count%9:]
	FilesList1.SetRect(0, 1, w-24, h-2)
	DriveInfo.SetRect(w-25, 1, w, h-2)
	//p.Text = fmt.Sprintf("%s : %d , %d:%d", p.Text[:20], count, w, h)

	dt := time.Now()
	HeadRight = dt.Format("01-02-2006 15:04:05")
	ScreenPrintAt(1, 0, termbox.ColorDefault, termbox.ColorDefault, HeadLeft)
	ScreenPrintAt(w-22, 0, termbox.ColorDefault, termbox.ColorDefault, HeadRight)

	ui.Render(Tree1, DriveInfo, FilesList1)
}
