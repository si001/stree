package screen

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var FilesList1 = widgets.NewList()
var FilesList2 = widgets.NewList()
var DriveInfo = widgets.NewList()
var HeadLeft string = "dir> "     // directory
var HeadRight string = "STree %s" // timer
var Tree1 *widgets.Tree
var Tree2 *widgets.Tree

const VM_TREEVIEW_FILES_1 = 1
const VM_FILELIST_1 = 2

var ViewMode = VM_TREEVIEW_FILES_1

const VC_INFO_WIDTH = 25
const VC_BOTTOM_HEIGHT = 1

var lastEvent termui.Event
