package model

import (
	"fmt"
	"si001/stree/widgets"
	"strconv"
	"time"
)

const (
	ATTR_NOTREAD     byte = 1
	ATTR_FILE        byte = 2
	ATTR_DIR         byte = 4
	ATTR_ARCH        byte = 16
	ATTR_SELECTED    byte = 32
	ATTR_ERR_MESSAGE byte = 255
)

type FileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
	AttrB   byte
	Owner   *widgets.TreeNode
}

func (self FileInfo) IsDir() bool {
	return self.AttrB == ATTR_DIR
}

func (self FileInfo) IsNotRead() bool {
	return (self.AttrB & ATTR_NOTREAD) > 0
}

func (self FileInfo) IsReadError() bool {
	return self.AttrB == ATTR_ERR_MESSAGE
}

// as fmt.Stringer
func (self FileInfo) String() string {
	return self.Name
}

// as widgets.ItemStringer
func (self FileInfo) ItemString(styleNumber, maxWidth int) string {
	r := ""
	switch styleNumber {
	case 0:
		r = fmt.Sprintf(" %s%"+strconv.Itoa(maxWidth-len([]rune(self.Name))-39)+"s%12s %4s %19s", self.Name, " ", parseSize(self.Size), parseAttr(self.AttrB), parseTime(self.ModTime))
	case 1:
		r = fmt.Sprintf(" %s%"+strconv.Itoa(maxWidth/2-len([]rune(self.Name))-14)+"s%12s ", self.Name, " ", parseSize(self.Size))
	case 2:
		r = fmt.Sprintf(" %s%"+strconv.Itoa(maxWidth/3-len([]rune(self.Name))-2)+"s", self.Name, " ")
	case 3:
		r = fmt.Sprintf(" %s%"+strconv.Itoa(maxWidth-len([]rune(self.Name))-2)+"s", self.Name, " ")
	}
	return r
}

func parseTime(dt time.Time) string {
	return dt.Format("02.01.2006 15:04:05")
}

func parseAttr(a byte) string {

	return strconv.Itoa(int(a))
}

func parseSize(i int64) (s string) {
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
	} else if i > 999 {
		s = fmt.Sprintf("%d", i)
		l := len(s)
		s = s[:l-3] + "," + s[l-3:]
	}
	return s
}
