package actions

import (
	"github.com/mattn/go-runewidth"
	"si001/stree/model"
	"si001/stree/screen/botton_box"
	"si001/stree/screen/botton_box/box_tools"
)

type boxFileMask struct {
	box_tools.BoxEditor
	maskCallback func(mask string)
}

func RequestFileMask(mask string, cb func(mask string)) {
	box := boxFileMask{
		BoxEditor: box_tools.BoxEditor{
			InterfaceText1: "Filespec: ",
			InterfaceText2: "",
			Text:           mask,
			Cursor:         runewidth.StringWidth(mask),
			HistoryId:      "filespec",
		},
		maskCallback: cb,
	}
	model.BottomMode = &box
	box.Callback = func(mask *string) {
		botton_box.NormalBottomBox()
		if mask != nil {
			box.maskCallback(*mask)
		}
	}
}
