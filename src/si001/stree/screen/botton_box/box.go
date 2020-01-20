package botton_box

import (
	"github.com/gdamore/tcell"
	"si001/stree/model"
	"si001/stree/widgets/stuff"
)

var normTree = &normalTree{}
var normList = &normalList{}

func NormalBottomBox() {
	if model.ViewMode() == model.VM_FILELIST_1 {
		model.BottomMode = normList
	} else {
		model.BottomMode = normTree
	}
}

func checkBottomMode() {
	if model.BottomMode == nil {
		model.BottomMode = normTree
	}
}

func Draw(s tcell.Screen) {
	checkBottomMode()
	model.BottomMode.Draw(s)
}

func ProcessEvent(event tcell.Event) bool {
	checkBottomMode()
	return model.BottomMode.ProcessEvent(event)
}

func SetListActions(acts []Action) {
	normList.actions = acts
}

func SetTreeActions(acts []Action) {
	normTree.actions = acts
}

func ActionsDraw(s tcell.Screen, actions []Action) {
	_, h := s.Size()
	style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)
	style2 := tcell.Style(0).Foreground(tcell.ColorYellow).Background(tcell.ColorDefault)

	acts := ""
	h -= 2
	for _, act := range actions {
		if len(act.ActName) == 0 && len(act.ActKey) == 0 {
			stuff.ScreenPrintWithSecondStyleAt(s, 0, h, style, style2, acts, '`')
			h++
			acts = ""
		} else if len(act.ActName) > 0 {
			acts += "  " + act.Name()
		}
	}
	if len([]rune(acts)) > 0 {
		stuff.ScreenPrintWithSecondStyleAt(s, 0, h, style, style2, acts, '`')
	}
	s.HideCursor()
}
