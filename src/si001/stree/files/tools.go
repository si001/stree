package files

import (
	"github.com/gizak/termui/v3/widgets"
	"io/ioutil"
	"log"
	"si001/stree/model"
	"strings"
)

func GetDirectory(path string) *model.Directory {
	osfiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return nil
	} else {
		dir := model.Directory{
			FileInfo: model.FileInfo{
				Name: path,
				Attr: 1,
			}}
		node := widgets.TreeNode{
			Value: dir,
		}
		for _, file := range osfiles {
			fInfo := model.FileInfo{Name: file.Name(), Size: file.Size(), ModTime: file.ModTime()}
			if file.IsDir() {
				fInfo.Attr = 1
				dInfo := newDirFI(fInfo, &node)
				node.Nodes = append(node.Nodes, dInfo)
			} else {
				dir.Files = append(dir.Files, &fInfo)
			}
		}
		return &dir
	}
}

func TreeNodeToPath(node *widgets.TreeNode) (result string) {
	for node != nil {
		result = node.Value.String() + result
		node = NodeGetParent(node)

		if node != nil {
			result = model.PathDivider + result
		}
	}
	switch {
	case strings.HasPrefix(result, "///"):
		result = result[2:]
	case strings.HasPrefix(result, "//") || strings.HasPrefix(result, "\\"):
		result = result[1:]
	}
	return result
}
