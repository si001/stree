// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	. "github.com/gdamore/tcell"
	"image"
	"si001/stree/widgets/stuff"
	//. "github.com/gdamore/tcell/termbox"
)

type List struct {
	stuff.Block
	Rows             []string
	WrapText         bool
	TextStyle        Style
	SelectedRow      int
	topRow           int
	SelectedRowStyle Style
}

func NewList() *List {
	return &List{
		Block:            *stuff.NewBlock(),
		TextStyle:        stuff.Theme.List.Text,
		SelectedRowStyle: stuff.Theme.List.Text.Foreground(ColorYellow),
	}
}

func (self *List) Draw(s Screen) {
	self.Block.Draw(s)

	point := self.Inner.Min

	// adjusts view into widget
	if self.SelectedRow >= self.Inner.Dy()+self.topRow {
		self.topRow = self.SelectedRow - self.Inner.Dy() + 1
	} else if self.SelectedRow < self.topRow {
		self.topRow = self.SelectedRow
	}

	// draw rows
	for row := self.topRow; row < len(self.Rows) && point.Y < self.Inner.Max.Y; row++ {
		style := self.TextStyle
		if row == self.SelectedRow {
			style = self.SelectedRowStyle
		}
		stuff.ScreenPrintAt(s, point.X+1, point.Y, style, self.Rows[row])
		point = image.Pt(self.Inner.Min.X, point.Y+1)
	}

	// draw UP_ARROW if needed
	if self.topRow > 0 {
		stuff.ScreenPrintAt(s, self.Inner.Min.X-1, self.Inner.Min.Y, self.BorderStyle, string(stuff.UP_ARROW))
	}
	// draw DOWN_ARROW if needed
	if len(self.Rows) > int(self.topRow)+self.Inner.Dy() {
		stuff.ScreenPrintAt(s, self.Inner.Min.X-1, self.Inner.Max.Y-1, self.BorderStyle, string(stuff.DOWN_ARROW))
	}
}

// ScrollAmount scrolls by amount given. If amount is < 0, then scroll up.
// There is no need to set self.topRow, as this will be set automatically when drawn,
// since if the selected item is off screen then the topRow variable will change accordingly.
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
		self.ScrollAmount(-self.Inner.Dy())
	}
}

func (self *List) ScrollPageDown() {
	self.ScrollAmount(self.Inner.Dy())
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
