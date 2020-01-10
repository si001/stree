package botton_box

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/model"
	"si001/stree/widgets/stuff"
	"strings"
)

type boxOrderBy struct {
	orderBy         byte
	orderByCallback func(order byte)
}

func (self *boxOrderBy) Draw(s tcell.Screen) {
	_, h := s.Size()
	style := tcell.Style(0).Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)

	stuff.ScreenPrintAt(s, 1, h-2, style, "OrderBy: Name, Ext, Size, Date/Time")

	var strPath string
	if self.orderBy&model.OrderByPath > 0 {
		strPath = "On "
	} else {
		strPath = "Off"
	}
	var strAcc string
	if self.orderBy&model.OrderAcs > 0 {
		strAcc = "Ascending"
	} else {
		strAcc = "Descending"
	}
	stuff.ScreenPrintAt(s, 1, h-1, style, fmt.Sprintf("         Path: %s  Order: %s", strPath, strAcc))
}

func (self *boxOrderBy) ProcessEvent(event tcell.Event) bool {
	result := model.OrderByUndefined
	switch ev := event.(type) {
	case *tcell.EventResize:
	case *tcell.EventMouse:
		switch ev.Buttons() {
		case tcell.Button2:
			self.orderByCallback(self.orderBy)
		}
	case *tcell.EventKey:
		switch strings.ToLower(ev.Name()) {
		case "esc":
			NormalBottomBox()
			return true
		case "rune[p]", "shift+rune[p]":
			self.orderBy = self.orderBy ^ model.OrderByPath
			return true
		case "rune[o]", "shift+rune[o]":
			self.orderBy = self.orderBy ^ model.OrderAcs
			return true
		case "rune[a]":
			self.orderBy = self.orderBy | model.OrderAcs
			return true
		case "shift+rune[a]":
			self.orderBy = self.orderBy | model.OrderAcs ^ model.OrderAcs
			return true
		case "rune[s]", "shift+rune[s]":
			self.orderBy = self.orderBy | model.OrderMask ^ model.OrderMask | model.OrderBySize
			result = self.orderBy
		case "rune[n]", "shift+rune[n]":
			self.orderBy = self.orderBy | model.OrderMask ^ model.OrderMask | model.OrderByName
			result = self.orderBy
		case "rune[e]", "shift+rune[e]":
			self.orderBy = self.orderBy | model.OrderMask ^ model.OrderMask | model.OrderByExt
			result = self.orderBy
		case "rune[d]", "shift+rune[d]", "rune[t]", "shift+rune[t]":
			self.orderBy = self.orderBy | model.OrderMask ^ model.OrderMask | model.OrderByDate
			result = self.orderBy
		}
	}
	if result != model.OrderByUndefined {
		self.orderByCallback(result)
		NormalBottomBox()
		return true
	}
	return true
}

func RequestOrderBy(ob byte, callback func(order byte)) {
	model.BottomMode = &boxOrderBy{
		orderBy:         ob,
		orderByCallback: callback,
	}
}
