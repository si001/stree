// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package stuff

import (
	"github.com/gdamore/tcell"
	"image"
	"sync"
)

// Block is the base struct inherited by most widgets.
// Block manages size, position, border, and title.
// It implements all 3 of the methods needed for the `Drawable` interface.
// Custom widgets will override the Draw method.
type Block struct {
	Border      bool
	BorderStyle tcell.Style

	BorderLeft, BorderRight, BorderTop, BorderBottom bool

	PaddingLeft, PaddingRight, PaddingTop, PaddingBottom int

	image.Rectangle
	Inner image.Rectangle

	Title      string
	TitleStyle tcell.Style

	sync.Mutex
}

func NewBlock() *Block {
	return &Block{
		Border:       true,
		BorderStyle:  Theme.Block.Border,
		BorderLeft:   true,
		BorderRight:  true,
		BorderTop:    true,
		BorderBottom: true,

		TitleStyle: Theme.Block.Title,
	}
}

func (self *Block) drawBorder(s tcell.Screen) {
	ScreenDrawBox(s, self.Min.X, self.Min.Y, self.Max.X, self.Max.Y, self.BorderStyle, ' ')
}

// Draw implements the Drawable interface.
func (self *Block) Draw(s tcell.Screen) {
	if self.Border {
		self.drawBorder(s)
	}
	ScreenPrintAt(s, self.Min.X+2, self.Min.Y, self.TitleStyle, self.Title)
}

// SetRect implements the Drawable interface.
func (self *Block) SetRect(x1, y1, x2, y2 int) {
	self.Rectangle = image.Rect(x1, y1, x2, y2)
	self.Inner = image.Rect(
		self.Min.X+1+self.PaddingLeft,
		self.Min.Y+1+self.PaddingTop,
		self.Max.X-1-self.PaddingRight,
		self.Max.Y-1-self.PaddingBottom,
	)
}

// GetRect implements the Drawable interface.
func (self *Block) GetRect() image.Rectangle {
	return self.Rectangle
}

func (self *Block) CheckIn(x, y int) bool {
	return x >= self.Min.X && x <= self.Max.X && y >= self.Min.Y && y <= self.Max.Y
}
