package botton_box

import (
	"github.com/gdamore/tcell"
	"si001/stree/model"
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
