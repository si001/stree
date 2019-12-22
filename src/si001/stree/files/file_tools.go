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
		Value: model.Directory{
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

func ReadDir(node *widgets.TreeNode) model.Directory /*, []model.Directory*/ {
	path := TreeNodeToPath(node)
	if model.PathDivider == "\\" && len(path) == 2 && path[1] == ':' {
		path += model.PathDivider
	}
	dir := node.Value.(model.Directory)
	osfiles, err := ioutil.ReadDir(path)
	if err != nil {
		dir.FileInfo.AttrB = model.ATTR_ERR_MESSAGE
		node.Value = dir
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
			}
		}
		node.Value = dir
		return dir
	}

}

//func GetDirectory(path string) *widgets.TreeNode{
//	osfiles, err := ioutil.ReadDir(path)
//	if err != nil {
//		log.Fatal(err)
//		return nil
//	} else {
//		dir := model.Directory{
//			FileInfo: model.FileInfo{
//				Name: path,
//				Attr: 1,
//			}}
//		node := widgets.TreeNode{
//			Value: dir,
//		}
//		for _, file := range osfiles {
//			fInfo := model.FileInfo{Name: file.Name(), Size: file.Size(), ModTime: file.ModTime()}
//			if file.IsDir() {
//				fInfo.Attr = 1
//				dInfo := newDirFI(fInfo, &node)
//				node.Nodes = append(node.Nodes, dInfo)
//			} else {
//				dir.Files = append(dir.Files, &fInfo)
//			}
//		}
//		return &dir
//	}
//}
