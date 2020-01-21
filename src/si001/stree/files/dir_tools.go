package files

import (
	"io/ioutil"
	"os"
	"runtime"
	"si001/stree/model"
	"si001/stree/widgets"
)

func newDir(nm string, parent *widgets.TreeNode) (dir *widgets.TreeNode) {
	dir = newDirFI(model.FileInfo{
		Name:  nm,
		AttrB: model.ATTR_NOTREAD,
	}, parent)
	return dir
}

func newDirFI(fInfo model.FileInfo, parent *widgets.TreeNode) (dir *widgets.TreeNode) {
	dir = &widgets.TreeNode{
		Value: &model.Directory{
			FileInfo: fInfo,
			Parent:   parent,
		},
	}
	if parent != nil {
		exist := false
		for _, item := range parent.Nodes {
			if item.Value.String() == dir.Value.String() {
				item.Value = dir.Value
				dir = item
				exist = true
				break
			}
		}
		if !exist {
			parent.Nodes = append(parent.Nodes, dir)
		}
	}
	return dir
}

func GetRoot() (r *widgets.TreeNode) {
	r = newDir("", nil)
	if runtime.GOOS == "windows" {
		for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
			_, err := os.Open(string(drive) + ":\\")
			if err == nil {
				r.Nodes = append(r.Nodes, newDir(string(drive)+":", nil))
			}
		}
	} else {
		r.Nodes = append(r.Nodes, newDir("/", nil))
	}
	return r
}

func ReadDir(node *widgets.TreeNode) *model.Directory {
	path := TreeNodeToPath(node)
	if model.PathDivider == "\\" && len(path) == 2 && path[1] == ':' {
		path += model.PathDivider
	}
	dir := node.Value.(*model.Directory)
	oldSize := -dir.Size
	oldCount := -dir.Count
	dir.Count = 0
	dir.Size = 0
	dir.Files = []*model.FileInfo{}
	osfiles, err := ioutil.ReadDir(path)
	if err != nil {
		dir.FileInfo.AttrB = model.ATTR_ERR_MESSAGE
		node.Value = dir
		pullDownFileInfoDelta(node, dir.Count-oldCount, dir.Size-oldSize)
		return dir
	} else {
		dir.FileInfo.AttrB = model.ATTR_DIR
		for _, file := range osfiles {
			fInfo := model.FileInfo{Name: file.Name(), Size: file.Size(), ModTime: file.ModTime(), AttrB: model.ATTR_FILE, Owner: node}
			if file.IsDir() {
				fInfo.AttrB = model.ATTR_NOTREAD
				newDirFI(fInfo, node)
			} else {
				fInfo.AttrB = model.ATTR_FILE
				dir.Files = append(dir.Files, &fInfo)
				dir.Count++
				dir.Size += fInfo.Size
			}
		}
		node.Value = dir
		pullDownFileInfoDelta(node, dir.Count-oldCount, dir.Size-oldSize)
		return dir
	}
}

func CloseDir(node *widgets.TreeNode) {
	node.Nodes = nil
	dir := node.Value.(*model.Directory)
	oldSize := -dir.Size
	oldCount := -dir.Count
	dir.AttrB = model.ATTR_NOTREAD | model.ATTR_DIR
	dir.Files = nil
	dir.Count = 0
	dir.Size = 0
	node.Value = dir
	pullDownFileInfoDelta(node, oldCount, oldSize)
}

func pullDownFileInfoDelta(node *widgets.TreeNode, deltaCount int32, deltaSize int64) {
	if node != nil {
		for {
			node = node.Value.(*model.Directory).Parent
			if node == nil {
				break
			}
			dir := node.Value.(*model.Directory)
			dir.Count += deltaCount
			dir.Size += deltaSize
		}
	}
}

func RefreshTreeNode(node *widgets.TreeNode) {
	//dir := node.Value.(*model.Directory)
	//deltaCount := dir.Count
	//deltaSize := dir.Size
	//for _, f := range dir.Files {
	//	deltaCount--
	//	deltaSize -= f.Size
	//}
	CloseDir(node)
	RefreshTreeNodeRecource(node)
	//pullDownFileInfoDelta(node, -deltaCount, -deltaSize)
}

func RefreshTreeNodeRecource(node *widgets.TreeNode) {
	ReadDir(node)
	for _, n := range node.Nodes {
		RefreshTreeNodeRecource(n)
	}
}
