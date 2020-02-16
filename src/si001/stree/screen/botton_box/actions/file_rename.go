package actions

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"si001/stree/model"
	"si001/stree/screen/botton_box"
	"si001/stree/screen/botton_box/box_tools"
)

type boxFileRename struct {
	box_tools.BoxEditor
	oldName        string
	path           string
	fileCallback   func(old, new string) error
	finishCallback func(newName *string)
}

func RequestRename(path, oldName, history, s1, s2, s3 string, cb func(oldName, newName string) error, cbFinish func(newName *string)) {
	box := boxFileRename{
		oldName: oldName,
		BoxEditor: box_tools.BoxEditor{
			InterfaceText1: s1,
			InterfaceText2: s2,
			InterfaceText3: s3,
			EditorBottom:   true,
			Text:           oldName,
			Cursor:         runewidth.StringWidth(oldName),
			HistoryId:      history,
		},
		fileCallback:   cb,
		finishCallback: cbFinish,
		path:           path,
	}

	model.BottomMode = &box
	box.BoxEditor.Callback = func(newName *string) {
		botton_box.NormalBottomBox()
		if newName != nil {
			err := box.fileCallback(box.path+box.oldName, box.path+*newName)
			if err == nil {
				box.finishCallback(newName)
			} else {
				RequestMessageBoxCenter(fmt.Sprintf("Error operation: %s", err), nil)
				box.finishCallback(nil)
			}
		}
	}
}
