package info_box

import (
	"github.com/gdamore/tcell"
	"si001/stree/model"
	"si001/stree/widgets/stuff"
	"time"
)

func ShowInfoBox(s tcell.Screen, fileMode int, fileMask, path string) {
	w, h := s.Size()
	style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)

	stuff.ScreenDrawBox(s, w-model.VC_INFO_WIDTH, 1, w-1, h-model.VC_BOTTOM_HEIGHT-1, style, ' ')

	var cp = path
	//if FileList1_IsBranch {
	//	cp = cp
	//}
	stuff.ScreenPrintAt(s, 1, 0, style, cp)
	//stuff.ScreenPrintAt(s, w-111, 0, style, lastEvent+" : "+tmpValue)
	dt := time.Now()
	tm := dt.Format("02.01.2006 15:04:05")
	stuff.ScreenPrintAt(s, w-22, 0, style, tm)

	// file mask
	stuff.ScreenPrintAt(s, w-23, 2, style, ""+fileMask)
	//stuff.ScreenPrintAt(s, w-5, 2, style, "F:"+strconv.Itoa(fileMode))

	stuff.ScreenDrawLine(s, w-model.VC_INFO_WIDTH, 3, w-1, 3, style, tcell.RuneLTee, tcell.RuneHLine, tcell.RuneRTee)

	// logged bytes

	//if model.DEBUG {
	//	stuff.ScreenPrintAt(s, 80, h-3, style, model.LastEvent)
	//}
	stuff.ScreenPrintAt(s, w-13, h-1, style, "STreeGo V0.1")
}
