package tree_list

import (
	"fmt"
	"si001/stree/files"
	"si001/stree/model"
	"si001/stree/screen/botton_box/actions"
	"si001/stree/widgets"
	"sort"
)

func (self *TreeAndList) actionMkDir() {
	tn := self.Tree.SelectedNode()
	if dir, ok := (tn.Value).(*model.Directory); ok {
		path := files.TreeNodeToPath(tn) + model.PathDivider
		if (tn.Value).(*model.Directory).AttrB&model.ATTR_NOTREAD > 0 {
			files.ReadDir(tn)
			self.ShowDir(path, tn, false, false)
		}
		actions.RequestMkDir(path, dir.Name, files.MkDir, func(nn string) {
			n := files.NewDir(nn, tn)
			sort.Slice(tn.Nodes, func(i, j int) bool {
				return tn.Nodes[i].Value.String() < tn.Nodes[j].Value.String()
			})
			files.ReadDir(n)
			self.Tree.PrepareNodes()
			self.Tree.Expand()
		})
	}
}

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
		n := []rune(tn.Value.(*model.Directory).Name)
		nm := make([]rune, len(n)*2)
		copy(nm, n)
		for i := len(n) - 1; i >= 0; i-- {
			copy(nm[i+1:], nm[i:])
			nm[i] = '`'
		}
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

func (self *TreeAndList) actionRename(isDir bool) {
	if isDir {
		tn := self.Tree.SelectedNode()
		if dir, ok := (tn.Value).(*model.Directory); ok && (tn.Value).(*model.Directory).Owner != nil {
			path := files.TreeNodeToPath(tn.Value.(*model.Directory).Owner) + model.PathDivider
			t1 := "           to: "
			t2 := "Rename folder: " + dir.Name + "  `↑ history  `D`e`l  `E`s`c  `E`n`t`e`r"
			actions.RequestRename(path, dir.Name, "foldername", t1, t2, files.FileRename, func(nn string) {
				dir.Name = nn
			})
		}
	} else {
		if fi, ok := (*self.List.SelectedStringer()).(*model.FileInfo); ok {
			path := files.TreeNodeToPath(fi.Owner)
			t1 := "         to: "
			t2 := "Rename file: " + fi.Name + "  `↑ history  `D`e`l  `E`s`c  `E`n`t`e`r"
			actions.RequestRename(path+model.PathDivider, fi.Name, "filename", t1, t2, files.FileRename, func(nn string) {
				fi.Name = nn
			})
		}
	}
}
