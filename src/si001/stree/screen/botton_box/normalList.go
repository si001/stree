package botton_box

import (
	"github.com/gdamore/tcell"
	"si001/stree/widgets/stuff"
	"strings"
)

type normalList struct {
	actions []Action
}

func (box *normalList) Draw(s tcell.Screen) {
	_, h := s.Size()
	style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)

	acts := ""
	for _, act := range box.actions {
		acts += "  " + act.Name()
	}
	stuff.ScreenPrintAt(s, 1, h-2, style, acts)
	stuff.ScreenPrintAt(s, 1, h-1, style, "` `<Esc>")
	s.HideCursor()
}

func (box *normalList) ProcessEvent(event tcell.Event) bool {
	switch ev := event.(type) {
	case *tcell.EventResize:
	case *tcell.EventMouse:
	case *tcell.EventKey:
		for _, act := range box.actions {
			if act.Key() == strings.ToLower(ev.Name()) {
				act.Doing()()
				return true
			}
		}
	}
	return false
}
