package tree_list

import (
	"github.com/gdamore/tcell"
	"math"
	"si001/stree/files"
	"si001/stree/model"
	"si001/stree/screen/info_box"
)

func (self *TreeAndList) Draw(s tcell.Screen, viewMode model.VmType, w, h int) {
	switch viewMode {
	case model.VM_TREEVIEW_FILES_1:
		self.drawTreeAndList(s, w, h)
	case model.VM_FILELIST_1:
		self.drawListOnly(s, w, h)
	}
}

func (self *TreeAndList) drawTreeAndList(s tcell.Screen, w, h int) {
	if self.Divider < 2 {
		self.Divider = 1
	}
	self.Tree.SetRect(0, 1, w-model.VC_INFO_WIDTH, h-self.Divider)
	self.List.SetRect(0, self.Tree.GetRect().Max.Y, w-model.VC_INFO_WIDTH+1, h-model.VC_BOTTOM_HEIGHT-1)
	//DriveInfo.SetRect(w-VC_INFO_WIDTH, 1, w-1, h-VC_BOTTOM_HEIGHT)

	if self.Divider <= 2 {
		self.Tree.Draw(s)
	} else {
		self.List.StyleNumber = self.FileMode
		self.List.Draw(s)
		self.Tree.Draw(s)
	}
	info_box.ShowInfoBox(s, self.Tree.GetRoot(), self.Tree.SelectedNode(), self.FileMask, model.CurrentPath)
}

func (self *TreeAndList) drawListOnly(s tcell.Screen, w, h int) {

	self.List.SetRect(0, 1, w-model.VC_INFO_WIDTH+1, h-model.VC_BOTTOM_HEIGHT-1)

	self.List.StyleNumber = self.FileMode
	self.List.Draw(s)
	fileS := self.List.SelectedStringer()
	var fileName string
	if fileS != nil {
		fileS := *fileS
		var path string
		fi, ok := fileS.(*model.FileInfo)
		if !ok {
			path = ""
		} else {
			path = files.TreeNodeToPath(fi.Owner)
		}
		fileName = path
		length := len([]rune(fileName))
		if length > w-30 && length > 12 {
			fileName = string([]rune(fileName)[:7]) + "***" + string(([]rune(fileName)[int(math.Max(10, float64(35+length-w))):]))
		}
	} else {
		fileName = ""
	}
	info_box.ShowInfoBox(s, self.Tree.GetRoot(), self.Tree.SelectedNode(), self.FileMask, model.CurrentPath)

	//stuff.ScreenPrintAt(s, 2, h-3, tcell.StyleDefault, "OB:"+strconv.Itoa(int(self.OrderBy)))
}
