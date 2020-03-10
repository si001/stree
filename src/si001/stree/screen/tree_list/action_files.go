package tree_list

import (
	"si001/stree/model"
	"si001/stree/screen/botton_box/actions"
	"si001/stree/tools/files"
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

func (self *TreeAndList) actionRename(isDir bool) {
	if isDir {
		tn := self.Tree.SelectedNode()
		if dir, ok := (tn.Value).(*model.Directory); ok && (tn.Value).(*model.Directory).Owner != nil {
			path := files.TreeNodeToPath(tn.Value.(*model.Directory).Owner) + model.PathDivider
			t1 := "           to: "
			t2 := "Rename folder: " + dir.Name
			t3 := "`↑ history  `D`e`l  `E`s`c  `E`n`t`e`r"
			actions.RequestRename(path, dir.Name, "foldername", t1, t2, t3, files.FileRename, func(nn *string) {
				if nn != nil {
					dir.Name = *nn
				}
			})
		}
	} else {
		if fi, ok := (*self.List.SelectedStringer()).(*model.FileInfo); ok {
			path := files.TreeNodeToPath(fi.Owner)
			t1 := "         to: "
			t2 := "Rename file: " + fi.Name
			t3 := "`↑ history  `D`e`l  `E`s`c  `E`n`t`e`r"
			actions.RequestRename(path+model.PathDivider, fi.Name, "filename", t1, t2, t3, files.FileRename, func(nn *string) {
				if nn != nil {
					fi.Name = *nn
				}
			})
		}
	}
}
