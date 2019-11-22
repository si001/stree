package app

import (
	"github.com/gizak/termui/v3/widgets"
	"io/ioutil"
	"log"
	"si001/stree/model"
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
			fInfo := model.FileInfo{file.Name(), file.Size(), file.ModTime(), 0}
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

type nodeValue string
