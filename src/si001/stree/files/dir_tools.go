package files

import (
	"io/ioutil"
	"os"
	"runtime"
	"si001/stree/model"
	"si001/stree/widgets"
	"sort"
)

func newDir(nm string, parent *widgets.TreeNode) (dir *widgets.TreeNode) {
	dir = newDirFI(model.FileInfo{
		Name:  nm,
		AttrB: model.ATTR_NOTREAD,
		Owner: parent,
	}, parent)
	return dir
}

func newDirFI(fInfo model.FileInfo, parent *widgets.TreeNode) (dir *widgets.TreeNode) {
	dir = &widgets.TreeNode{
		Value: &model.Directory{
			FileInfo: fInfo,
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
	return ReadDirPath(node, path)
}
func ReadDirPath(node *widgets.TreeNode, path string) *model.Directory {
	if model.PathDivider == "\\" && len(path) == 2 && path[1] == ':' {
		path += model.PathDivider
	}
	dir := node.Value.(*model.Directory)
	ct, sz, tc, ts := calcCountSize(node)
	oldSize := dir.Size
	oldCount := dir.Count
	oldTagSize := -ts
	oldTagCount := -tc
	dir.Count -= ct
	dir.Size -= sz
	dir.Files = []*model.FileInfo{}
	osfiles, err := ioutil.ReadDir(path)
	if err != nil {
		dir.FileInfo.AttrB = model.ATTR_ERR_MESSAGE
		node.Value = dir
		pullDownFileInfoDelta(node, dir.Count-oldCount, dir.Size-oldSize, oldTagCount, oldTagSize)
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
		sort.Slice(node.Nodes, func(i, j int) bool {
			return UpcaseIfWindows(node.Nodes[i].Value.String()) < UpcaseIfWindows(node.Nodes[j].Value.String())
		})

		node.Value = dir
		pullDownFileInfoDelta(node, dir.Count-oldCount, dir.Size-oldSize, oldTagCount, oldTagSize)
		return dir
	}
}

func calcCountSize(node *widgets.TreeNode) (int32, int64, int32, int64) {
	var sz, ts int64 = 0, 0
	var tc int32 = 0
	for _, f := range node.Value.(*model.Directory).Files {
		sz += f.Size
		if f.IsTagged() {
			tc++
			ts += f.Size
		}
	}
	return int32(len(node.Value.(*model.Directory).Files)), sz, tc, ts
}

func CloseDir(node *widgets.TreeNode) {
	node.Nodes = nil
	dir := node.Value.(*model.Directory)
	oldSize := -dir.Size
	oldCount := -dir.Count
	oldTagSize := -dir.TagSize
	oldTagCount := -dir.TagCount
	dir.AttrB = model.ATTR_NOTREAD | model.ATTR_DIR
	dir.Files = nil
	dir.Count = 0
	dir.Size = 0
	dir.TagCount = 0
	dir.TagSize = 0
	node.Value = dir
	pullDownFileInfoDelta(node, oldCount, oldSize, oldTagCount, oldTagSize)
}

func PullDownFileInfoDeltaTag(node *widgets.TreeNode, deltaTagCount int32, deltaTagSize int64) {
	pullDownFileInfoDelta(node, 0, 0, deltaTagCount, deltaTagSize)
}

func pullDownFileInfoDelta(node *widgets.TreeNode, deltaCount int32, deltaSize int64, deltaTagCount int32, deltaTagSize int64) {
	if node != nil {
		for {
			node = node.Value.(*model.Directory).FileInfo.Owner
			if node == nil {
				break
			}
			dir := node.Value.(*model.Directory)
			dir.Count += deltaCount
			dir.Size += deltaSize
			dir.TagCount += deltaTagCount
			dir.TagSize += deltaTagSize
		}
	}
}

func SizeFiles(files []*model.FileInfo) int64 {
	var sz int64 = 0
	for _, fi := range files {
		sz += fi.Size
	}
	return sz
}

func RefreshTreeNode(node *widgets.TreeNode) {
	CloseDir(node)
	RefreshTreeNodeRecource(node)
}

func RefreshTreeNodeRecource(node *widgets.TreeNode) {
	ReadDir(node)
	for _, n := range node.Nodes {
		RefreshTreeNodeRecource(n)
	}
}
