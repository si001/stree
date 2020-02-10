package actions

import (
	"github.com/gdamore/tcell"
	"si001/stree/model"
	"si001/stree/screen/botton_box"
	"si001/stree/widgets/stuff"
	"strings"
)

type boxMessage struct {
	isCenter     bool
	message      string
	needYNAnswer bool
	callback     func(result bool)
}

func (box *boxMessage) Draw(s tcell.Screen) {
	if box.isCenter {
		box.DrawCenter(s)
	} else {
		box.DrawBottom(s)
	}
}

func (box *boxMessage) DrawBottom(s tcell.Screen) {
	_, h := s.Size()
	style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)
	style2 := tcell.Style(0).Foreground(tcell.ColorYellow).Background(tcell.ColorDefault)

	stuff.ScreenPrintWithSecondStyleAt(s, 1, h-2, style, style2, box.message, '`')
	stuff.ScreenPrintWithSecondStyleAt(s, 2, h-1, style, style2, "Press `Yes or `No key.", '`')
}

func (box *boxMessage) DrawCenter(s tcell.Screen) {
	w, h := s.Size()
	style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)
	style2 := tcell.Style(0).Foreground(tcell.ColorYellow).Background(tcell.ColorDefault)

	cx, cy := w/2, h/2
	tw := len([]rune(box.message)) / 2
	sx, sy := tw+5, 3

	if w > sx*2+5 {
		stuff.ScreenDrawBox(s, cx-sx, cy-sy, cx+sx, cy+sy, style, ' ')
		stuff.ScreenPrintAt(s, cx-tw, cy, style, box.message)
	} else {
		rs := []rune(box.message)
		s1 := string(rs[:w*2/3])
		s2 := string(rs[w*2/3:])
		tw = len([]rune(s1)) / 2
		sx = tw + 5
		stuff.ScreenDrawBox(s, cx-sx, cy-sy, cx+sx, cy+sy, style, ' ')
		stuff.ScreenPrintAt(s, cx-tw, cy, style, s1)
		stuff.ScreenPrintAt(s, cx-tw, cy+1, style, s2)
	}

	if box.needYNAnswer {
		stuff.ScreenPrintWithSecondStyleAt(s, 2, h-2, style, style2, "Press `Yes or `No key.", '`')
	} else {
		stuff.ScreenPrintWithSecondStyleAt(s, 2, h-2, style, style2, "Please press any key", '`')
	}
}

func (box *boxMessage) ProcessEvent(event tcell.Event) bool {
	switch ev := event.(type) {
	case *tcell.EventResize:
	case *tcell.EventMouse:
		switch ev.Buttons() {
		case tcell.Button1:
			box.success(false)
		case tcell.Button2:
			box.success(false)
		}
	case *tcell.EventKey:
		switch strings.ToLower(ev.Name()) {
		case "esc":
			box.success(false)
		case "rune[n]":
			box.success(false)
		case "rune[y]":
			box.success(true)
		case "enter":
			box.success(true)
		}
	}
	return true
}

func (box *boxMessage) success(success bool) {
	botton_box.NormalBottomBox()
	if box.callback != nil {
		box.callback(success)
	}
}

func RequestYNMessageBoxCenter(message string, callback func(result bool)) {
	model.BottomMode = &boxMessage{
		message:      message,
		callback:     callback,
		isCenter:     true,
		needYNAnswer: true,
	}
}

func RequestYNMessageBox(message string, callback func(result bool)) {
	model.BottomMode = &boxMessage{
		message:      message,
		callback:     callback,
		isCenter:     false,
		needYNAnswer: true,
	}
}

func RequestMessageBoxCenter(message string, callback func(result bool)) {
	model.BottomMode = &boxMessage{
		message:      message,
		callback:     callback,
		isCenter:     true,
		needYNAnswer: false,
	}
}
