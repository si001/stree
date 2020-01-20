package botton_box

import (
	"github.com/gdamore/tcell"
	"strings"
)

type normalList struct {
	actions []Action
}

func (box *normalList) Draw(s tcell.Screen) {
	ActionsDraw(s, box.actions)
}

func (box *normalList) ProcessEvent(event tcell.Event) bool {
	switch ev := event.(type) {
	case *tcell.EventResize:
	case *tcell.EventMouse:
	case *tcell.EventKey:
		for _, act := range box.actions {
			if act.Key() == strings.ToLower(ev.Name()) || "shift+"+act.Key() == strings.ToLower(ev.Name()) {
				act.Doing()()
				return true
			}
		}
	}
	return false
}
