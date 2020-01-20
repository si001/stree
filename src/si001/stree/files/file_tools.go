package files

import (
	"fmt"
	"io"
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

func ReadDir(node *widgets.TreeNode) *model.Directory /*, []model.Directory*/ {
	path := TreeNodeToPath(node)
	if model.PathDivider == "\\" && len(path) == 2 && path[1] == ':' {
		path += model.PathDivider
	}
	dir := node.Value.(*model.Directory)
	dir.Files = []*model.FileInfo{}
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

func FileRename(oldName, newName string) error {
	err := os.Rename(oldName, newName)
	return err
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func FileCopy(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
