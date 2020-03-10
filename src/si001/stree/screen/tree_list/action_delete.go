package tree_list

import (
	"fmt"
	"si001/stree/model"
	"si001/stree/screen/botton_box/actions"
	"si001/stree/tools"
	"si001/stree/tools/files"
	"si001/stree/widgets"
)

func (self *TreeAndList) actionRmDir() {
	tn := self.Tree.SelectedNode()
	if (tn.Value).(*model.Directory).Owner != nil {
		path := files.TreeNodeToPath(tn) + model.PathDivider
		if (tn.Value).(*model.Directory).AttrB&model.ATTR_NOTREAD > 0 {
			files.ReadDir(tn)
			self.ShowDir(path, tn, false, false)
		}
		if len(tn.Nodes) > 0 || len(tn.Value.(*model.Directory).Files) > 0 {
			actions.RequestMessageBoxCenter("You can't remove a NOT EMPTY folder!", nil)
			return
		}
		nm := tools.ToBright(tn.Value.(*model.Directory).Name)
		pTxt := files.TreeNodeToPath(tn.Value.(*model.Directory).Owner) + model.PathDivider + string(nm)
		actions.RequestYNMessageBox(fmt.Sprintf("DELETE dir %s ?", pTxt), func(yes bool) {
			if yes {
				err := files.Remove(path)
				if err != nil {
					actions.RequestMessageBoxCenter(fmt.Sprintf("You can't process the operation: %s", err), nil)
				} else {
					own := tn.Value.(*model.Directory).Owner
					var nnn []*widgets.TreeNode
					for _, n := range own.Nodes {
						if n != tn {
							nnn = append(nnn, n)
						}
					}
					own.Nodes = nnn
					self.Tree.PrepareNodes()
					self.Tree.ScrollUp()
				}
			}
		})
	} else {
		actions.RequestMessageBoxCenter("You can't remove a root folder!", nil)
	}
}

func (self *TreeAndList) actionRmFiles() {

}
func (self *TreeAndList) actionRmFile() {
	if fi, ok := (*self.List.SelectedStringer()).(*model.FileInfo); ok {
		path := files.TreeNodeToPath(fi.Owner) + model.PathDivider
		nm := tools.ToBright(fi.Name)
		actions.RequestYNMessageBox(fmt.Sprintf("DELETE file %s ?", nm), func(yes bool) {
			if yes {
				err := files.Remove(path + fi.Name)
				if err != nil {
					actions.RequestMessageBoxCenter(fmt.Sprintf("You can't process the operation: %s", err), nil)
				} else {
					var newFiles []*model.FileInfo
					for _, f := range fi.Owner.Value.(*model.Directory).Files {
						if f.Name != fi.Name {
							newFiles = append(newFiles, f)
						}
					}
					fi.Owner.Value.(*model.Directory).Files = newFiles
					var newBranch []*fmt.Stringer
					for _, f := range self.List.Rows {
						if (*f).(*model.FileInfo).Name != fi.Name {
							newBranch = append(newBranch, f)
						}
					}
					var tc int32 = 0
					var ts int64 = 0
					if fi.IsTagged() {
						tc, ts = 1, fi.Size
					}
					files.ApplyInfoDelta(fi.Owner, -1, -fi.Size, -tc, -ts, true)
					self.List.Rows = newBranch
				}
			}
		})
	}
}
