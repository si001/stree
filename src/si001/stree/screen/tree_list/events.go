package tree_list

import (
	"github.com/gdamore/tcell"
	"si001/stree/model"
)

func (self *TreeAndList) PutEvent(event tcell.Event) bool {
	switch model.ViewMode() {
	case model.VM_TREEVIEW_FILES_1:
		return self.PutEventTreeList(event)
	case model.VM_FILELIST_1:
		return self.PutEventList(event)
	}
	return false
}
