package tree_list

import (
	"si001/stree/model"
)

type nullInterface interface {
}

func setTagFile(tl *TreeAndList, set bool) {
	s := tl.List.SelectedStringer()
	if s != nil {
		(*s).(*model.FileInfo).SetTagged(set)
		tl.List.ScrollAmount(1)
	}
}

func setTagAllFiles(tl *TreeAndList, set bool) {
	for _, s := range tl.List.Rows {
		if s != nil {
			(*s).(*model.FileInfo).SetTagged(set)
		}
	}
}

func setTagDir(tl *TreeAndList, set bool) {
	s := tl.Tree.SelectedNode()
	if s != nil {
		f := *s
		f.Value.(*model.Directory).SetSelected(set)
		tl.Tree.ScrollAmount(1)
		tl.pathCheck()
	}
}
