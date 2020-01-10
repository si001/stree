package botton_box

import (
	"github.com/gdamore/tcell"
	"si001/stree/model"
	"si001/stree/widgets/stuff"
)

type boxFileMask struct {
	fileMask     string
	maskCallback func(mask string)
	cursor       int
}

func (box *boxFileMask) Draw(s tcell.Screen) {
	_, h := s.Size()
	style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)

	stuff.ScreenPrintAt(s, 1, h-1, style, "Mask: "+box.fileMask)
	s.ShowCursor(7+box.cursor, h-1)
}

func (box *boxFileMask) ProcessEvent(event tcell.Event) bool {
	switch ev := event.(type) {
	case *tcell.EventResize:
	case *tcell.EventMouse:
		switch ev.Buttons() {
		case tcell.Button2:
			box.finish(nil)
		}
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEsc:
			box.finish(nil)
		case tcell.KeyEnter:
			box.finish(&box.fileMask)
		case tcell.KeyLeft:
			if box.cursor > 0 {
				box.cursor--
			}
		case tcell.KeyRight:
			if box.cursor < len(box.fileMask) {
				box.cursor++
			}
		case tcell.KeyBackspace:
			if box.cursor > 0 {
				box.fileMask = box.fileMask[0:box.cursor-1] + box.fileMask[box.cursor:]
				box.cursor--
			}
		case tcell.KeyDelete:
			if box.cursor < len(box.fileMask) {
				box.fileMask = box.fileMask[0:box.cursor] + box.fileMask[box.cursor+1:]
			}
		case tcell.KeyHome:
			box.cursor = 0
		case tcell.KeyEnd:
			box.cursor = len(box.fileMask)
		case tcell.KeyRune:
			//if box.cursor >= len(box.fileMask) {
			//	box.fileMask = box.fileMask[0:box.cursor] + string(ev.Rune())
			//} else {
			box.fileMask = box.fileMask[0:box.cursor] + string(ev.Rune()) + box.fileMask[box.cursor:]
			//}
			box.cursor++
		}
	}
	return true
}

func (box *boxFileMask) finish(mask *string) {
	NormalBottomBox()
	if mask != nil {
		box.maskCallback(*mask)
	}
}

func RequestFileMask(mask string, cb func(mask string)) {
	model.BottomMode = &boxFileMask{
		fileMask:     mask,
		maskCallback: cb,
		cursor:       0,
	}
}
