package widgets

import (
	"fmt"
	"github.com/gdamore/tcell"
	"si001/stree/widgets/stuff"
)

type List struct {
	stuff.Block
	Rows             []*fmt.Stringer
	TextStyle        tcell.Style
	SelectedRow      int
	topRow           int
	columns          int
	SelectedRowStyle tcell.Style
	StyleNumber      int
}

type ItemStringer interface {
	ItemString(styleNumber, maxWidth int) string
}

type Styler interface {
	Style() tcell.Style
	SelectedStyle() tcell.Style
}

func NewList() *List {
	return &List{
		Block:            *stuff.NewBlock(),
		TextStyle:        stuff.Theme.Tree.Text,
		SelectedRowStyle: stuff.Theme.Tree.Text.Foreground(tcell.ColorWhite).Background(tcell.ColorBlue),
	}
}

func (self *List) ScrollToMouse(x, y int) {
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
	self.SelectedRow = sel
}

func (self *List) Draw(s tcell.Screen) {
	self.Block.Draw(s)

	point := self.Inner.Min

	if self.SelectedRow >= (self.Inner.Dy()+1)*self.columns+self.topRow {
		self.topRow = self.SelectedRow - (self.Inner.Dy()+1)*self.columns + 1
	} else if self.SelectedRow < self.topRow {
		self.topRow = self.SelectedRow
	}

	row := self.topRow
	x := point.X
	cellWidth := 1
	col := 0
	for x < self.Inner.Max.X-cellWidth/2 {
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
				itemStr = itemStrgr.ItemString(self.StyleNumber, self.Inner.Dx())
			} else {
				itemStr = (*r).String()
			}
			lng := len([]rune(itemStr))
			if cellWidth < lng {
				cellWidth = lng
			}
			stuff.ScreenPrintAtTween(s, x, point.Y, x+len(itemStr), style, style, "", itemStr)
			point.Y++
		}
		x += cellWidth
		col++
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
		self.SelectedRow = self.topRow
	} else {
		if self.columns > 1 {
			self.ScrollAmount(-self.Inner.Dy() - 1)
		} else {
			self.ScrollAmount(-self.Inner.Dy())
		}
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
}
