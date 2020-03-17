package box_tools

import (
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
	"si001/stree/tools"
	"si001/stree/tools/files/settings"
	"si001/stree/widgets/stuff"
	"strings"
)

type BoxEditor struct {
	InterfaceText1 string
	InterfaceText2 string
	InterfaceText3 string
	EditorBottom   bool
	Text           string
	Callback       func(result *string)
	Cursor         int
	HistoryId      string
	HistoryWidth   int
	SpecialF2Func  func(func(result *string)) (func(s tcell.Screen), func(event tcell.Event) bool)

	History         *HistoryTool
	startX          int
	subDraw         func(s tcell.Screen)
	subProcessEvent func(event tcell.Event) bool
}

func (box *BoxEditor) Draw(s tcell.Screen) {
	w, h := s.Size()
	style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)
	styleHl := tcell.Style(0).Foreground(tcell.ColorYellow).Background(tcell.ColorDefault)
	box.startX = len([]rune(box.InterfaceText1)) + 2
	if tools.BrightLen(box.InterfaceText2) > w/3-5 {
		w = tools.BrightLen(box.InterfaceText2) + 4
	} else {
		w = w / 3
	}
	cx, cy := box.startX+box.Cursor, h-1
	if box.EditorBottom {
		stuff.ScreenPrintWithSecondStyleAt(s, 2, h-2, style, styleHl, box.InterfaceText2, '`')
		stuff.ScreenPrintWithSecondStyleAt(s, w, h-2, style, styleHl, box.InterfaceText3, '`')
		stuff.ScreenPrintWithSecondStyleAt(s, 2, h-1, style, styleHl, box.InterfaceText1+box.Text, '`')
	} else {
		stuff.ScreenPrintWithSecondStyleAt(s, 2, h-2, style, styleHl, box.InterfaceText1+box.Text, '`')
		stuff.ScreenPrintWithSecondStyleAt(s, 2, h-1, style, styleHl, box.InterfaceText2, '`')
		stuff.ScreenPrintWithSecondStyleAt(s, w, h-1, style, styleHl, box.InterfaceText3, '`')
		cy = h - 2
	}
	if box.subDraw != nil {
		s.HideCursor()
		box.subDraw(s)
	} else {
		s.ShowCursor(cx, cy)
	}
}

func (box *BoxEditor) ProcessEvent(event tcell.Event) bool {
	if box.subProcessEvent != nil {
		box.subProcessEvent(event)
		return true
	}

	switch ev := event.(type) {
	case *tcell.EventResize:
	case *tcell.EventMouse:
		switch ev.Buttons() {
		case tcell.Button2:
			box.Callback(nil)
		}
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEsc:
			box.Callback(nil)
		case tcell.KeyEnter:
			box.editComplete()
		case tcell.KeyLeft:
			if box.Cursor > 0 {
				box.Cursor--
			}
		case tcell.KeyRight:
			if box.Cursor < runewidth.StringWidth(box.Text) {
				box.Cursor++
			}
		case tcell.KeyBackspace, tcell.KeyBackspace2:
			if box.Cursor > 0 {
				box.Text =
					string(([]rune(box.Text))[0:box.Cursor-1]) + string(([]rune(box.Text))[box.Cursor:])
				box.Cursor--
			}
		case tcell.KeyDelete:
			if box.Cursor < runewidth.StringWidth(box.Text) {
				box.Text = string(([]rune(box.Text))[0:box.Cursor]) + string(([]rune(box.Text))[box.Cursor+1:])
			}
		case tcell.KeyHome:
			box.Cursor = 0
		case tcell.KeyEnd:
			box.Cursor = runewidth.StringWidth(box.Text)
		case tcell.KeyRune:
			box.Text = string(([]rune(box.Text))[0:box.Cursor]) + string(ev.Rune()) + string(([]rune(box.Text))[box.Cursor:])
			box.Cursor++
		case tcell.KeyUp:
			box.ShowHistory()
		case tcell.KeyF2:
			box.SpecialRun()
		}
	}
	return true
}

func (box *BoxEditor) editComplete() {
	if len(box.HistoryId) > 0 {
		items := settings.ReadHistory(box.HistoryId)
		var newIts []string
		for _, i := range items {
			if i != box.Text {
				newIts = append(newIts, i)
			}
		}
		newIts = append(newIts, box.Text)
		settings.WriteHistory(box.HistoryId, newIts)
	}
	box.Callback(&box.Text)
}

func (box *BoxEditor) ShowHistory() {
	if len(box.HistoryId) > 0 {
		box.History = &HistoryTool{
			HistoryId: box.HistoryId,
			Left:      box.startX,
			Width:     box.HistoryWidth,
			Callback: func(res *string) {
				if res != nil {
					box.Text = strings.Trim(*res, " ")
					box.Cursor = len([]rune(box.Text))
				}
				box.History = nil
				box.subDraw = nil
				box.subProcessEvent = nil
			},
		}
		box.History.Init()
		box.subDraw = box.History.Draw
		box.subProcessEvent = box.History.ProcessEvent
	}
}

func (box *BoxEditor) SpecialRun() {
	if box.SpecialF2Func != nil {
		box.subDraw, box.subProcessEvent = box.SpecialF2Func(func(result *string) {
			box.subDraw = nil
			box.subProcessEvent = nil
			if result != nil {
				box.Text = *result
				box.Cursor = len([]rune(string(box.Text)))
			}
		})
	}
}
