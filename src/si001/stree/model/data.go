package model

import (
	"github.com/gdamore/tcell"
)

const CharSelector = "●"

var PathDivider = "?"
var CurrentPath string
var Root Directory
var SelectedStyle tcell.Style
var ScreenWidth, ScreenHeight int

var DEBUG = true

type VmType int

const (
	VM_TREEVIEW_FILES_1 VmType = 1 + iota
	VM_FILELIST_1
)

var viewMode = VM_TREEVIEW_FILES_1

func ViewModeChng(typ VmType) {
	viewMode = typ
}
func ViewMode() VmType {
	return viewMode
}

const (
	BM_NORMAL int = 0 + iota
	BM_FILE_MASK
	BM_ORDER
	BM_COPY
	BM_MOVE
	BM_DELETE
	BM_SEARCH
)

const VC_INFO_WIDTH = 25
const VC_BOTTOM_HEIGHT = 2

var LastEvent string

var BottomMode BottomBox = nil

type BottomBox interface {
	Draw(s tcell.Screen)
	ProcessEvent(event tcell.Event) bool
}
