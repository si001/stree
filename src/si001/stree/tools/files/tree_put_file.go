package files

import (
	"fmt"
	"si001/stree/model"
	"si001/stree/widgets"
	"strings"
)

func PutFileToPath(path string, file *model.FileInfo, tree *widgets.Tree) (resNode *widgets.TreeNode, resStr *string) {
	ph := strings.Split(path, model.PathDivider)
	var nodePath []*widgets.TreeNode
	var dir *widgets.TreeNode
	for i, dirNm := range ph {
		if len([]rune(strings.TrimSpace(dirNm))) == 0 {
			continue
		}
		dirNm = UpcaseIfWindows(dirNm)
		if i == 0 {
			if dirNm == "" {
				dirNm = model.PathDivider
			}
			read := GetRoot()
			tRoot := &widgets.TreeNode{
				Nodes: tree.GetRoot(),
			}
			_, dir = checkNodes(tRoot, read, nil, dirNm)
		} else {
			//dir = newDir(dirNm, dir)
			old := dir
			for _, d := range dir.Nodes {
				if UpcaseIfWindows(d.Value.String()) == dirNm {
					dir = d
					break
				}
			}
			if dir == old || dir == nil {
				s := fmt.Sprintf("Path '%s' is incorrect!", path)
				return nil, &s
			}
		}
		if dir == nil {
			s := fmt.Sprintf("Path '%s' is incorrect!", path)
			return nil, &s
		}
		nodePath = append(nodePath, dir)
		var old widgets.TreeNode
		(old.Value) = &model.Directory{
			FileInfo: (dir.Value).(*model.Directory).FileInfo,
		}
		(old.Value).(*model.Directory).FileInfo.Owner = nil
		path := TreeNodeToPath(dir)
		old.Value = ReadDirPath(&old, path)
		(old.Value).(*model.Directory).FileInfo.Owner = (dir.Value).(*model.Directory).FileInfo.Owner
		_, _ = checkNodes(dir, &old, dir, "")
	}

	node := nodePath[len(nodePath)-1]
	file.Owner = node
	node.Value.(*model.Directory).Files = append(node.Value.(*model.Directory).Files, file)
	PushDownFileInfoDelta(node, 1, file.Size, 0, 0)
	return node, nil
}
