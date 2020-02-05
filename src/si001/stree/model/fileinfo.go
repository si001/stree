package model

import (
	"fmt"
	"si001/stree/widgets"
	"strconv"
	"strings"
	"time"
)

type AttrFile byte

const (
	ATTR_NOTREAD AttrFile = 1 << iota
	ATTR_FILE
	ATTR_DIR
	ATTR_ARCH
	ATTR_SELECTED
	ATTR_ERR_MESSAGE = 255
)

const (
	OrderByUndefined byte = 0 + iota
	OrderByName
	OrderByExt
	OrderBySize
	OrderByDate
	OrderAcs    = 0x40
	OrderByPath = 0x80
	OrderMask   = 0x0f
)

type FileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
	AttrB   AttrFile
	Owner   *widgets.TreeNode
}

func (fi *FileInfo) IsDir() bool {
	return fi.AttrB == ATTR_DIR
}

func (fi *FileInfo) IsNotRead() bool {
	return (fi.AttrB & ATTR_NOTREAD) > 0
}

func (fi *FileInfo) IsReadError() bool {
	return fi.AttrB == ATTR_ERR_MESSAGE

}

func (fi *FileInfo) SetTagged(sel bool) {
	if sel {
		fi.AttrB = fi.AttrB | ATTR_SELECTED
	} else {
		fi.AttrB = fi.AttrB | ATTR_SELECTED ^ ATTR_SELECTED
	}
}

func (fi *FileInfo) IsTagged() bool {
	return fi.AttrB&ATTR_SELECTED > 0
}

// as fmt.Stringer
func (fi *FileInfo) String() string {
	return fi.Name
}

// as widgets.ItemStringer
func (fi *FileInfo) ItemString(styleNumber, maxWidth int) (value string, tagged bool) {
	r := ""
	var fs string
	if fi.IsTagged() {
		fs = CharSelector
	} else {
		fs = " "
	}
	switch styleNumber {
	case 0:
		r = fmt.Sprintf("%s%s  ", fs, fi.Name)
	case 1:
		cw := 30
		d := maxWidth / cw
		cw = maxWidth/d - 15
		txt := cropStringTo(fi.Name, cw)
		r = fmt.Sprintf("%s%s% 12s ", fs, txt, ParseSize(fi.Size))
	case 2:
		cw := maxWidth/2 - 41 + 5
		txt := cropStringTo(fi.Name, cw)
		r = fmt.Sprintf("%s%s %12s %19s ", fs, txt, ParseSize(fi.Size), ParseTime(fi.ModTime))
		//r = fmt.Sprintf(" %s %12s %4s %19s ", txt, parseSize(self.Size), parseAttr(self.AttrB), parseTime(self.ModTime))
	case 3:
		cw := maxWidth - 41 + 4
		txt := cropStringTo(fi.Name, cw)
		r = fmt.Sprintf("%s%s %12s  %19s", fs, txt, ParseSize(fi.Size), ParseTime(fi.ModTime))
		//r = fmt.Sprintf(" %s %12s %4s %19s", txt, parseSize(self.Size), parseAttr(self.AttrB), parseTime(self.ModTime))
		//case 4:
		//	r = fmt.Sprintf(" %s%"+strconv.Itoa(maxWidth-len([]rune(self.Name))-2)+"s", self.Name, " ")
	}
	return r, fi.IsTagged()
}

func cropStringTo(str string, maxLen int) string {
	if maxLen < 0 {
		return "►"
	}
	txt := string([]rune(str + strings.Repeat(" ", maxLen))[:maxLen])
	if len([]rune(str)) > maxLen {
		txt += "►" //string(tcell.RuneRArrow)
	} else {
		txt += " "
	}
	return txt
}

func ParseTime(dt time.Time) string {
	return dt.Format("02.01.2006 15:04:05")
}

func parseAttr(a byte) string {
	return strconv.Itoa(int(a))
}

func ParseSize(i int64) (s string) {
	if i > 999999999 {
		i /= 1000
		if i > 999999999 {
			i /= 1000
			s = fmt.Sprintf("%dM", i)
		} else {
			s = fmt.Sprintf("%dK", i)
		}
		l := len(s)
		s = s[:l-7] + "," + s[l-7:]
		s = s[:l-3] + "," + s[l-3:]
	} else if i > 999999 {
		s = fmt.Sprintf("%d", i)
		l := len(s)
		s = s[:l-6] + "," + s[l-6:]
		s = s[:l-2] + "," + s[l-2:]
	} else {
		s = fmt.Sprintf("%d", i)
		if i > 999 {
			l := len(s)
			s = s[:l-3] + "," + s[l-3:]
		}
	}
	return s
}

func ParseSizeMax(i int64) (s string) {
	if i > 999999999999 {
		i /= 1000
		d1 := i % 1000
		d2 := i / 1000 % 1000
		d3 := i / 1000000 % 1000
		d4 := i / 1000000000 % 1000
		s = fmt.Sprintf("%d,%d,%d,%dK", d4, d3, d2, d1)
	} else if i > 999999999 {
		d1 := i % 1000
		d2 := i / 1000 % 1000
		d3 := i / 1000000 % 1000
		d4 := i / 1000000000 % 1000
		s = fmt.Sprintf("%d,%d,%d,%d", d4, d3, d2, d1)
	} else if i > 999999 {
		s = fmt.Sprintf("%d", i)
		l := len(s)
		s = s[:l-6] + "," + s[l-6:]
		s = s[:l-2] + "," + s[l-2:]
	} else {
		s = fmt.Sprintf("%d", i)
		if i > 999 {
			l := len(s)
			s = s[:l-3] + "," + s[l-3:]
		}
	}
	return s
}
