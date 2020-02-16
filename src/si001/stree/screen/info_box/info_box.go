package info_box

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/model"
	"si001/stree/tools/files"
	"si001/stree/widgets"
	"si001/stree/widgets/stuff"
	"strings"
	"time"
)

func ShowInfoBox(s tcell.Screen, root []*widgets.TreeNode, selected *widgets.TreeNode, fileMask, path string) {
	w, h := s.Size()
	style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)
	style2 := tcell.Style(0).Foreground(tcell.ColorYellow).Background(tcell.ColorDefault)

	stuff.ScreenDrawBox(s, w-model.VC_INFO_WIDTH, 1, w-1, h-model.VC_BOTTOM_HEIGHT-1, style, ' ')

	lt := files.TreeNodeToPath(selected)
	if strings.Compare(lt, path) == 0 {
		stuff.ScreenPrintAt(s, 1, 0, style, path)
	} else {
		lr := []rune(lt)
		pr := []rune(path)
		c := 0
		for c < len(pr) && c < len(lr) && pr[c] == lr[c] {
			c++
		}
		x := stuff.ScreenPrintStr(s, 1, 0, style, string(lr[:c]))
		stuff.ScreenPrintStr(s, x, 0, style2, string(pr[c:]))
	}

	//	stuff.ScreenPrintAt(s, w-40, h-2, style, p)

	//stuff.ScreenPrintAt(s, w-111, 0, style, lastEvent+" : "+tmpValue)
	dt := time.Now()
	tm := dt.Format("02.01.2006 15:04:05")
	stuff.ScreenPrintAt(s, w-22, 0, style, tm)

	// file mask
	stuff.ScreenPrintAt(s, w-23, 2, style, ""+fileMask)
	//stuff.ScreenPrintAt(s, w-5, 2, style, "F:"+strconv.Itoa(fileMode))

	stuff.ScreenDrawLine(s, w-model.VC_INFO_WIDTH, 3, w-1, 3, style, tcell.RuneLTee, tcell.RuneHLine, tcell.RuneRTee)

	row := 4
	for _, r := range root {
		dir := r.Value.(*model.Directory)
		stuff.ScreenPrintAt(s, w-23, row, style, fmt.Sprintf("%s   count % 10s", dir.Name, model.ParseSize(int64(dir.Count))))
		row++
		stuff.ScreenPrintAt(s, w-23, row, style, fmt.Sprintf(" bytes% 15s", model.ParseSize(dir.Size)))
		row++
	}

	dir := selected.Value.(*model.Directory)
	stuff.ScreenPrintAt(s, w-23, row, style, fmt.Sprintf("Current:"))
	row++
	stuff.ScreenPrintAt(s, w-23, row, style, fmt.Sprintf(" files count% 9s", model.ParseSize(int64(len(dir.Files)))))
	row++
	stuff.ScreenPrintAt(s, w-23, row, style, fmt.Sprintf(" bytes% 15s", model.ParseSize(files.SizeFiles(dir.Files))))
	row++

	stuff.ScreenPrintAt(s, w-23, row, style, fmt.Sprintf("Branch:"))
	row++
	stuff.ScreenPrintAt(s, w-23, row, style, fmt.Sprintf(" files count% 9s", model.ParseSize(int64(dir.Count))))
	row++
	stuff.ScreenPrintAt(s, w-23, row, style, fmt.Sprintf(" bytes% 15s", model.ParseSize(dir.Size)))
	row++

	tc, ts := calcTagged(root)
	stuff.ScreenPrintAt(s, w-23, row, style, fmt.Sprintf("Tagged:"))
	row++
	stuff.ScreenPrintAt(s, w-23, row, style, fmt.Sprintf(" files count% 9s", model.ParseSize(int64(tc))))
	row++
	stuff.ScreenPrintAt(s, w-23, row, style, fmt.Sprintf(" bytes% 15s", model.ParseSize(ts)))
	row++

	stuff.ScreenPrintAt(s, w-13, h-1, style, "STreeGo V0.1")

	//stuff.ScreenPrintAt(s, w-20, h-1, style, model.LastEvent)
}

func calcTagged(root []*widgets.TreeNode) (int32, int64) {
	var tc int32
	var ts int64
	for _, n := range root {
		tc += n.Value.(*model.Directory).TagCount
		ts += n.Value.(*model.Directory).TagSize
	}
	return tc, ts
}
