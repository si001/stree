package tree_list

import (
	"si001/stree/screen/botton_box"
)

func (self *TreeAndList) actionsList() {
	botton_box.SetListActions([]botton_box.Action{
		botton_box.Action{
			ActName: "Tag",
			ActKey:  "rune[t]",
			Callback: func() {
				setTagFile(self, true)
			},
		},
		botton_box.Action{
			ActName: "^Tag All",
			ActKey:  "ctrl+t",
			Callback: func() {
				setTagAllFiles(self, true)
			},
		},
		botton_box.Action{
			ActName: "Untag",
			ActKey:  "rune[u]",
			Callback: func() {
				setTagFile(self, false)
			},
		},
		botton_box.Action{
			ActName: "^Untag All",
			ActKey:  "ctrl+u",
			Callback: func() {
				setTagAllFiles(self, false)
			},
		},
	})
}
