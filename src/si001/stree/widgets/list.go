package widgets

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/widgets/stuff"
	"sort"
	"strings"
)

type List struct {
	stuff.Block
	Rows             []*fmt.Stringer
	TextStyle        tcell.Style
	SelectedRow      int
	topRow           int
	columns          int
	SelectedRowStyle tcell.Style
	TaggedRowStyle   tcell.Style
	StyleNumber      int
	SingleColumn     bool
}

type SimpleStringer struct {
	Value string
}

func (s SimpleStringer) String() string {
	return s.Value
}

type ItemStringer interface {
	ItemString(styleNumber, maxWidth int) (value string, tagged bool)
}

type Styler interface {
	Style() tcell.Style
	SelectedStyle() tcell.Style
}

func NewList() *List {
	return &List{
		Block:            *stuff.NewBlock(),
		TextStyle:        stuff.StyleClear,
		SelectedRowStyle: stuff.Theme.Tree.Text.Foreground(tcell.ColorWhite).Background(tcell.ColorBlue),
		TaggedRowStyle:   stuff.Theme.Tree.Text.Foreground(tcell.ColorYellow).Background(tcell.ColorDefault),
	}
}

func (self *List) ScrollToMouse(x, y int) {
	if self.columns <= 0 {
		return
	}
	sel := y - self.Min.Y + self.topRow - 1
	if sel < 0 {
		sel = 0
	} else if sel >= len(self.Rows) {
		sel = len(self.Rows) - 1
	}
	cw := self.Inner.Dx() / self.columns
	ww := x - self.Min.X
	cl := ww / cw
	sel += cl * (self.Inner.Dy() + 1)
	if sel >= len(self.Rows) {
		sel = len(self.Rows) - 1
	}
	self.SelectedRow = sel
}

func (self *List) Draw(s tcell.Screen) {
	self.Block.Draw(s)

	if len(self.Rows) == 0 {
		return
	}
	point := self.Inner.Min

	if self.topRow > len(self.Rows)-self.Inner.Dy() {
		self.topRow = len(self.Rows) - self.Inner.Dy()
	}
	if self.SelectedRow >= (self.Inner.Dy()+1)*self.columns+self.topRow {
		self.topRow = self.SelectedRow - (self.Inner.Dy()+1)*self.columns + 1
	} else if self.SelectedRow < self.topRow {
		self.topRow = self.SelectedRow
	}
	if self.topRow < 0 {
		self.topRow = 0
	}

	row := self.topRow
	x := point.X
	cellWidth := 1
	col := 0
	for x < self.Inner.Max.X-cellWidth*3/4 {
		cellWidth = 1
		// draw rows
		point.Y = self.Inner.Min.Y
		for ; row < len(self.Rows) && point.Y <= self.Inner.Max.Y; row++ {
			r := self.Rows[row]
			var style tcell.Style
			if stl, ok := (*r).(Styler); ok {
				if row == self.SelectedRow {
					style = stl.SelectedStyle()
				} else {
					style = stl.Style()
				}
			} else {
				if row == self.SelectedRow {
					style = self.SelectedRowStyle
				} else {
					style = self.TextStyle
				}
			}
			var itemStr string
			if itemStrgr, ok := (*r).(ItemStringer); ok {
				var tagged bool
				itemStr, tagged = itemStrgr.ItemString(self.StyleNumber, self.Inner.Dx())
				if tagged && row != self.SelectedRow {
					style = self.TaggedRowStyle
				}
			} else {
				itemStr = (*r).String()
			}
			lng := len([]rune(itemStr))
			if self.SingleColumn && lng < self.Inner.Dx() {
				itemStr += strings.Repeat(" ", self.Inner.Dx()-lng)
				lng = self.Inner.Dx()
			}
			if cellWidth < lng {
				cellWidth = lng
			}
			stuff.ScreenPrintAtTween(s, x, point.Y, x+len([]rune(itemStr)), style, style, "", itemStr)
			point.Y++
		}
		x += cellWidth
		if cellWidth > 1 { // i.e. strings are exists
			col++
		}
	}
	if col < 1 {
		col = 1
	}
	self.columns = col

	dy2 := float32(self.SelectedRow) / float32(len(self.Rows)-1)
	stuff.ScreenDrawScrolled(s, self.Inner, dy2, self.topRow > 0, len(self.Rows) > int(self.topRow)+self.Inner.Dy()+1, self.BorderStyle)
}

func (self *List) ScrollAmount(amount int) {
	if len(self.Rows)-int(self.SelectedRow) <= amount {
		self.SelectedRow = len(self.Rows) - 1
	} else if int(self.SelectedRow)+amount < 0 {
		self.SelectedRow = 0
	} else {
		self.SelectedRow += amount
	}
}

func (self *List) ScrollUp() {
	self.ScrollAmount(-1)
}

func (self *List) ScrollDown() {
	self.ScrollAmount(1)
}

func (self *List) ScrollPageUp() {
	// If an item is selected below top row, then go to the top row.
	if self.SelectedRow > self.topRow {
		if self.columns > 1 {
			self.ScrollAmount(-self.Inner.Dy() - 1)
		} else {
			self.SelectedRow = self.topRow
		}
	} else {
		self.ScrollAmount(-self.Inner.Dy())
	}
}

func (self *List) ScrollPageDown() {
	if self.columns > 1 {
		self.ScrollAmount(self.Inner.Dy() + 1)
	} else {
		self.ScrollAmount(self.Inner.Dy())
	}
}

func (self *List) ScrollHalfPageUp() {
	self.ScrollAmount(-int(stuff.FloorFloat64(float64(self.Inner.Dy()) / 2)))
}

func (self *List) ScrollHalfPageDown() {
	self.ScrollAmount(int(stuff.FloorFloat64(float64(self.Inner.Dy()) / 2)))
}

func (self *List) ScrollTop() {
	self.SelectedRow = 0
}

func (self *List) ScrollBottom() {
	self.SelectedRow = len(self.Rows) - 1
	self.topRow = self.SelectedRow - (self.Inner.Dy()+1)*self.columns + 1
}

func (self *List) SelectedStringer() (res *fmt.Stringer) {
	if len(self.Rows) <= self.SelectedRow || self.SelectedRow < 0 {
		return nil
	} else {
		return self.Rows[self.SelectedRow]
	}
}

func (self *List) SetOrderComparator(comparator func(val1, val2 int) bool) {
	sort.Slice(self.Rows, comparator)
}
