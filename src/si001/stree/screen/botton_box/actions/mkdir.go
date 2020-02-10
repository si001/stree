package actions

import (
	"fmt"
	"si001/stree/model"
	"si001/stree/screen/botton_box"
	"si001/stree/screen/botton_box/box_tools"
)

type boxMkDir struct {
	box_tools.BoxEditor
	owner          string
	path           string
	fileCallback   func(dirName string) error
	finishCallback func(newName string)
}

func RequestMkDir(path, owner string, cb func(dirName string) error, cbFinish func(newName string)) {
	box := boxMkDir{
		owner: owner,
		BoxEditor: box_tools.BoxEditor{
			InterfaceText1: "            as: ",
			InterfaceText2: "Make directory under: " + owner + "  `↑ history  `D`e`l  `E`s`c  `E`n`t`e`r",
			EditorBottom:   true,
			Text:           "",
			Cursor:         0,
			HistoryId:      "foldername",
		},
		fileCallback:   cb,
		finishCallback: cbFinish,
		path:           path,
	}

	model.BottomMode = &box
	box.BoxEditor.Callback = func(newName *string) {
		botton_box.NormalBottomBox()
		if newName != nil {
			err := box.fileCallback(box.path + *newName)
			if err == nil {
				box.finishCallback(*newName)
			} else {
				RequestMessageBoxCenter(fmt.Sprintf("Error operation: /n%s", err), nil)
			}
		}
	}
}
