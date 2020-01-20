package botton_box

import (
	"github.com/gdamore/tcell"
	"strings"
)

type normalTree struct {
	actions []Action
}

func (box *normalTree) Draw(s tcell.Screen) {
	ActionsDraw(s, box.actions)
}

func (box *normalTree) ProcessEvent(event tcell.Event) bool {
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
