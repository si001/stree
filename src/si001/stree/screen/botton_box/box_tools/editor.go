package box_tools

import (
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
	"si001/stree/files"
	"si001/stree/widgets/stuff"
)

type BoxEditor struct {
	InterfaceText1 string
	InterfaceText2 string
	EditorBottom   bool
	Text           string
	Callback       func(result *string)
	Cursor         int
	HistoryId      string
	HistoryWidth   int

	history *historyTool
	startX  int
}

func (box *BoxEditor) Draw(s tcell.Screen) {
	_, h := s.Size()
	style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)
	styleHl := tcell.Style(0).Foreground(tcell.ColorYellow).Background(tcell.ColorDefault)
	box.startX = len([]rune(box.InterfaceText1)) + 2
	if box.EditorBottom {
		stuff.ScreenPrintWithSecondStyleAt(s, 2, h-2, style, styleHl, box.InterfaceText2, '`')
		stuff.ScreenPrintWithSecondStyleAt(s, 2, h-1, style, styleHl, box.InterfaceText1+box.Text, '`')
		//stuff.ScreenPrintAt(s, 2, h-2, style, box.InterfaceText2)
		//stuff.ScreenPrintAt(s, 2, h-1, style, box.InterfaceText1+box.Text)
		s.ShowCursor(box.startX+box.Cursor, h-1)
	} else {
		stuff.ScreenPrintWithSecondStyleAt(s, 2, h-2, style, styleHl, box.InterfaceText1+box.Text, '`')
		stuff.ScreenPrintWithSecondStyleAt(s, 2, h-1, style, styleHl, box.InterfaceText2, '`')
		s.ShowCursor(box.startX+box.Cursor, h-2)
	}
	if box.history != nil {
		box.history.Draw(s)
	}
}

func (box *BoxEditor) ProcessEvent(event tcell.Event) bool {
	if box.history != nil {
		box.history.ProcessEvent(event)
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
		case tcell.KeyBackspace:
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
			box.showHistory()
		}
	}
	return true
}

func (box *BoxEditor) editComplete() {
	if len(box.HistoryId) > 0 {
		items := files.ReadHistory(box.HistoryId)
		var newIts []string
		for _, i := range items {
			if i != box.Text {
				newIts = append(newIts, i)
			}
		}
		newIts = append(newIts, box.Text)
		files.WriteHistory(box.HistoryId, newIts)
	}
	box.Callback(&box.Text)
}

func (box *BoxEditor) showHistory() {
	if len(box.HistoryId) > 0 {
		box.history = &historyTool{
			historyId: box.HistoryId,
			left:      box.startX,
			callback: func(res *string) {
				if res != nil {
					box.Text = *res
					box.Cursor = len(box.Text)
				}
				box.history = nil
			},
		}
		box.history.Init()
	}
}