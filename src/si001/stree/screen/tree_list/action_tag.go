package tree_list

import (
	"si001/stree/model"
	"si001/stree/widgets"
)

type nullInterface interface {
}

func setTagFile(tl *TreeAndList, set, step bool) {
	s := tl.List.SelectedStringer()
	if s != nil {
		(*s).(*model.FileInfo).SetTagged(set)
		if step {
			tl.List.ScrollAmount(1)
		}
	}
}

func setTagAllFiles(tl *TreeAndList, set bool) {
	for _, s := range tl.List.Rows {
		if s != nil {
			(*s).(*model.FileInfo).SetTagged(set)
		}
	}
}

func setTagDir(tl *TreeAndList, set, step bool) {
	s := tl.Tree.SelectedNode()
	if s != nil {
		f := *s
		f.Value.(*model.Directory).SetSelected(set)
		if step {
			tl.Tree.ScrollAmount(1)
		}
		tl.pathCheck()
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
