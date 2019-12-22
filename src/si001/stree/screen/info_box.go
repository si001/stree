package screen

import (
	"github.com/gdamore/tcell"
	"si001/stree/model"
	"si001/stree/widgets/stuff"
	"strconv"
	"time"
)

func ShowInfoBox(s tcell.Screen, fileMode int, fileMask, path string) {
	w, h := s.Size()
	style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)

	stuff.ScreenDrawBox(s, w-VC_INFO_WIDTH, 1, w-1, h-VC_BOTTOM_HEIGHT, style, ' ')

	var cp = model.CurrentPath
	if FileList1_IsBranch {
		cp += "(BRANCH)"
	}
	stuff.ScreenPrintAt(s, 1, 0, style, cp)
	//stuff.ScreenPrintAt(s, w-111, 0, style, lastEvent+" : "+tmpValue)
	dt := time.Now()
	tm := dt.Format("02.01.2006 15:04:05")
	stuff.ScreenPrintAt(s, w-22, 0, style, tm)

	// file mask
	stuff.ScreenPrintAt(s, w-23, 2, style, "MASK "+FileMask1)
	stuff.ScreenPrintAt(s, w-5, 2, style, "F:"+strconv.Itoa(FilesMode1))

	stuff.ScreenDrawLine(s, w-VC_INFO_WIDTH, 3, w-1, 3, style, tcell.RuneLTee, tcell.RuneHLine, tcell.RuneRTee)

	// logged bytes

	stuff.ScreenPrintAt(s, w-23, 10, style, lastEvent)

}
