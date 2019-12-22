package screen

import (
	"si001/stree/widgets"
)

var FilePath1 string
var FilesMode1 int = 0
var FileMask1 string
var Tree1 *widgets.Tree
var FilesList1 = widgets.NewList()
var FileList1_IsBranch bool

var FilePath2 string
var FilesMode2 int = 0
var FileMask2 string
var Tree2 *widgets.Tree
var FilesList2 = widgets.NewList()

var DriveInfo = widgets.NewList()

const VM_TREEVIEW_FILES_1 = 1
const VM_FILELIST_1 = 2

var ViewMode = VM_TREEVIEW_FILES_1

const VC_INFO_WIDTH = 25
const VC_BOTTOM_HEIGHT = 1

var lastEvent string
