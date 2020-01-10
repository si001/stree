package tree_list

import (
	"si001/stree/screen/botton_box"
)

func (self *TreeAndList) actionsTree() {
	var actions = []botton_box.Action{
		botton_box.Action{
			ActName: "TagD",
			ActKey:  "rune[t]",
			Callback: func() {
				setTagDir(self, true)
			},
		},
		botton_box.Action{
			ActName: "UntagD",
			ActKey:  "rune[u]",
			Callback: func() {
				setTagDir(self, false)
			},
		},
	}
	botton_box.SetTreeActions(actions)
}
