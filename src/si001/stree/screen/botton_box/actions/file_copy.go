package actions

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
	"si001/stree/model"
	"si001/stree/screen/botton_box"
	"si001/stree/screen/botton_box/box_tools"
	"si001/stree/tools"
)

type boxFileCopyAs struct {
	box_tools.BoxEditor
}

type boxFileCopyTo struct {
	box_tools.BoxEditor
}

func RequestCopy(path, oldName, s1, s2, s3, c1, c2, c3 string, startSelectFolder func(s func(result *string)) (func(s tcell.Screen), func(event tcell.Event) bool), cbFinish func(newPath, newName *string)) {
	box := boxFileCopyAs{
		BoxEditor: box_tools.BoxEditor{
			InterfaceText1: s1,
			InterfaceText2: fmt.Sprintf(s2, oldName),
			InterfaceText3: s3,
			EditorBottom:   true,
			Text:           oldName,
			Cursor:         runewidth.StringWidth(oldName),
			HistoryId:      "filename",
		},
	}
	model.BottomMode = &box
	box.BoxEditor.Callback = func(newName *string) {
		if newName != nil {
			c22 := fmt.Sprintf("%s as %s", tools.ToBright(oldName), tools.ToBright(*newName))
			box2 := boxFileCopyTo{
				BoxEditor: box_tools.BoxEditor{
					InterfaceText1: c1,
					InterfaceText2: fmt.Sprintf(c2, c22),
					InterfaceText3: c3,
					EditorBottom:   true,
					Text:           "",
					Cursor:         0,
					HistoryId:      "foldername",
					SpecialF2Func:  startSelectFolder,
				},
			}
			box2.BoxEditor.Callback = func(newPath *string) {
				botton_box.NormalBottomBox()
				cbFinish(newPath, newName)
			}
			model.BottomMode = &box2
		} else {
			botton_box.NormalBottomBox()
			cbFinish(nil, nil)
		}
	}
}
