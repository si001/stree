package files

import (
	"github.com/gizak/termui/v3/widgets"
	"io/ioutil"
	"os"
	"runtime"
	"si001/stree/model"
)

func newDir(nm string, parent *widgets.TreeNode) (dir *widgets.TreeNode) {
	dir = newDirFI(model.FileInfo{
		Name: nm,
		Attr: model.ATTR_NOTREAD,
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
		dir.FileInfo.Name = err.Error()
		dir.FileInfo.Attr = model.ATTR_ERR_MESSAGE
		return dir
	} else {
		dir.FileInfo.Attr = model.ATTR_DIR
		for _, file := range osfiles {
			fInfo := model.FileInfo{Name: file.Name(), Size: file.Size(), ModTime: file.ModTime(), Attr: model.ATTR_FILE}
			if file.IsDir() {
				fInfo.Attr = model.ATTR_NOTREAD
				newDirFI(fInfo, node)
			} else {
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
