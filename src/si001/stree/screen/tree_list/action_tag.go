package tree_list

import (
	"si001/stree/files"
	"si001/stree/model"
	"si001/stree/widgets"
)

type nullInterface interface {
}

func setTagFile(tl *TreeAndList, set, step bool) {
	s := tl.List.SelectedStringer()
	if s != nil {
		fi := (*s).(*model.FileInfo)
		setTagFileInfo(fi, set)
		if step {
			tl.List.ScrollAmount(1)
		}
	}
}

func setTagAllFiles(tl *TreeAndList, set bool) {
	for _, s := range tl.List.Rows {
		if s != nil {
			setTagFileInfo((*s).(*model.FileInfo), set)
		}
	}
}

func setTagDir(tl *TreeAndList, set, step bool) {
	s := tl.Tree.SelectedNode()
	if s != nil {
		f := *s
		for _, f := range f.Value.(*model.Directory).Files {
			setTagFileInfo(f, set)
		}
		if step {
			tl.Tree.ScrollAmount(1)
		}
		tl.pathCheck()
	}
}

func setTagFileInfo(fi *model.FileInfo, set bool) {
	if fi.IsTagged() != set {
		node := fi.Owner
		dir := node.Value.(*model.Directory)
		fi.SetTagged(set)
		if set {
			dir.TagCount++
			dir.TagSize += fi.Size
			files.PullDownFileInfoDeltaTag(node, 1, fi.Size)
		} else {
			dir.TagCount--
			dir.TagSize -= fi.Size
			files.PullDownFileInfoDeltaTag(node, -1, -fi.Size)
		}
	}
}

func checkIsTagged(node *widgets.TreeNode) bool {
	d, ok := node.Value.(*model.Directory)
	c := 0
	if ok {
		for _, f := range d.Files {
			if f.IsTagged() {
				c++
			}
		}
		if len(d.Files) < c*2 {
			return true
		}
	}
	return false
}
