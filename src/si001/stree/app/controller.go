package app

import (
	"github.com/gdamore/tcell"
	"si001/stree/model"
	"si001/stree/screen"
	"si001/stree/screen/botton_box"
	"si001/stree/widgets/stuff"
)

func processEvent(event tcell.Event) {
	if !botton_box.ProcessEvent(event) {
		screen.TreeAndList1.PutEvent(event)
	}
}

func draw(s tcell.Screen, count int) {
	w, h := s.Size()
	model.ScreenWidth, model.ScreenHeight = w, h
	stuff.ScreenFillBox(s, 0, 0, w, h, tcell.StyleDefault, ' ')
	screen.TreeAndList1.Draw(s, model.ViewMode(), w, h)
	botton_box.Draw(s)
	s.Show()
}
