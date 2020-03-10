package tree_list

import (
	"si001/stree/model"
	"si001/stree/screen/botton_box"
	"si001/stree/screen/botton_box/actions"
)

func (self *TreeAndList) actionsList() {
	botton_box.SetListActions([]botton_box.Action{
		botton_box.Action{
			ActName: "`Tag",
			ActKey:  "rune[t]",
			Callback: func() {
				setTagFile(self, true, true)
			},
		},
		botton_box.Action{
			ActName: "`Untag",
			ActKey:  "rune[u]",
			Callback: func() {
				setTagFile(self, false, true)
			},
		},
		botton_box.Action{
			ActName: "`Filespec",
			ActKey:  "rune[f]",
			Callback: func() {
				actions.RequestFileMask(self.FileMask, self.setFileMask)
			},
		},
		botton_box.Action{
			ActName: "`Alt+`File display",
			ActKey:  "alt+rune[f]",
			Callback: func() {
				self.processNextFileMode()
			},
		},
		botton_box.Action{
			ActName: "`Alt+`Sort criteria",
			ActKey:  "alt+rune[s]",
			Callback: func() {
				actions.RequestOrderBy(self.OrderBy, self.setOrderBy)
			},
		},
		botton_box.Action{
			ActName: "",
			ActKey:  "rune[ ]",
			Callback: func() {
				self.List.ScrollDown()
			},
		},
		botton_box.Action{
			ActName: "`Rename",
			ActKey:  "rune[r]",
			Callback: func() {
				self.actionRename(false)
			},
		},
		botton_box.Action{
			ActName: "`Copy",
			ActKey:  "rune[c]",
			Callback: func() {
				self.actionCopy()
			},
		},
		botton_box.Action{
			ActName: "`Ctl+`Copy tagged",
			ActKey:  "ctrl+c",
			Callback: func() {
				self.actionCopyTagged()
			},
		},
		botton_box.Action{
			ActName: "`Delete",
			ActKey:  "rune[d]",
			Callback: func() {
				self.actionRmFile()
			},
		},
		botton_box.Action{
			ActName: "`Ctl+`Delete",
			ActKey:  "ctrl+d",
			Callback: func() {
				self.actionRmFiles()
			},
		},
		botton_box.Action{
			ActName: "`Log path",
			ActKey:  "rune[l]",
			Callback: func() {
				self.actionLog()
				ViewModeChange(model.VM_TREEVIEW_FILES_1)
			},
		},
		botton_box.Action{
			ActName: "",
			ActKey:  "",
			Callback: func() {
			},
		},
		botton_box.Action{
			ActName: "`Ctl+`Tag all",
			ActKey:  "ctrl+t",
			Callback: func() {
				setTagAllFiles(self, true)
			},
		},
		botton_box.Action{
			ActName: "`Ctl+`Untag all",
			ActKey:  "ctrl+u",
			Callback: func() {
				setTagAllFiles(self, false)
			},
		},
		botton_box.Action{
			ActName: "`Quit",
			ActKey:  "rune[q]",
			Callback: func() {
				actionQuit()
			},
		},
	})
}
