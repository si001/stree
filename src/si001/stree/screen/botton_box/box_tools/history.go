package box_tools

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/tools/files/settings"
	"si001/stree/widgets"
)

type HistoryTool struct {
	HistoryId             string
	Callback              func(res *string)
	Left                  int
	Width                 int
	list                  *widgets.List
	mouseLastEvent        *tcell.EventMouse
	mouseClickTimerForDbl int64
	mouseLastSelectedRow  int
	mouseTagging          bool
}

type historyItem string

func (hi historyItem) String() string {
	return string(hi)
}

func (hst *HistoryTool) Draw(s tcell.Screen) {
	w, h := s.Size()
	//style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)
	x, y := hst.Left, 4
	if x == 0 {
		x = (w - hst.Width) / 2
	}
	hst.list.SetRect(x, y, x+hst.Width, h-3)
	hst.list.Draw(s)
	s.HideCursor()
}

func (hst *HistoryTool) ProcessEvent(event tcell.Event) bool {
	switch ev := event.(type) {
	case *tcell.EventResize:
	case *tcell.EventMouse:
		toLastEvent := ev
		switch ev.Buttons() {
		case tcell.Button1:
			if ev.Buttons()&hst.mouseLastEvent.Buttons() == 0 {
				var ms int64 = ev.When().UnixNano() / 1000000
				var result string = (*hst.list.SelectedStringer()).String()
				if ms-hst.mouseClickTimerForDbl < 400 {
					hst.Callback(&result)
				}
				hst.mouseClickTimerForDbl = ms
				_, y := ev.Position()
				if hst.list.CheckIn(ev.Position()) && y != hst.list.Min.Y && y != hst.list.Max.Y {
					hst.list.ScrollToMouse(ev.Position())
				} else {
					hst.Callback(&result)
				}
			} else {
				if hst.list.CheckIn(ev.Position()) || hst.list.CheckIn(hst.mouseLastEvent.Position()) {
					hst.list.ScrollToMouse(ev.Position())
					toLastEvent = hst.mouseLastEvent
				}
			}
		case tcell.Button2:
			hst.Callback(nil)
		case tcell.ButtonNone:
			hst.mouseLastSelectedRow = -1
		case tcell.WheelUp:
			hst.list.ScrollUp()
		case tcell.WheelDown:
			hst.list.ScrollDown()
		}
		//x, y := ev.Position()
		//model.LastEvent = fmt.Sprintf("%d:%d / %s : %s", x, y, string(ev.Buttons()), ev.Modifiers())
		hst.mouseLastEvent = toLastEvent
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEsc:
			hst.Callback(nil)
		case tcell.KeyEnter:
			var s string
			sr := hst.list.SelectedStringer()
			if sr != nil {
				s = (*sr).String()
			} else {
				s = ""
			}
			hst.Callback(&s)
		case tcell.KeyDelete:
			if hst.list.SelectedRow >= 0 && hst.list.SelectedRow < len(hst.list.Rows) {
				hst.list.Rows = append(hst.list.Rows[:hst.list.SelectedRow], hst.list.Rows[hst.list.SelectedRow+1:]...)
			}
		case tcell.KeyUp:
			hst.list.ScrollUp()
		case tcell.KeyDown:
			hst.list.ScrollDown()
		case tcell.KeyHome:
			hst.list.ScrollTop()
		case tcell.KeyEnd:
			hst.list.ScrollBottom()
		case tcell.KeyPgDn:
			hst.list.ScrollPageDown()
		case tcell.KeyPgUp:
			hst.list.ScrollPageUp()
		}
	}
	return true
}

func (hst *HistoryTool) Init() {
	hst.list = widgets.NewList()
	hst.list.SingleColumn = true
	s := settings.ReadHistory(hst.HistoryId)
	if hst.Width == 0 {
		hst.Width = 30
	}
	for _, i := range s {
		var hi fmt.Stringer = historyItem(i)
		hst.list.Rows = append(hst.list.Rows, &hi)
		if hst.Width < len([]rune(i))+2 {
			hst.Width = len([]rune(i)) + 2
		}
	}
	hst.list.ScrollBottom()
}

func (hst *HistoryTool) GetList() *widgets.List {
	return hst.list
}
