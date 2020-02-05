package tree_list

import (
	"si001/stree/files"
	"si001/stree/screen/botton_box/actions"
)

func (self *TreeAndList) actionLog() {
	var rs []string
	for _, r := range self.Tree.GetRoot() {
		rs = append(rs, r.Value.String())
	}

	actions.RequestLog(rs, func(logPath string) {
		//root := self.Tree.GetRoot()

		files.LogTree(logPath, self.Tree)
	})
}
