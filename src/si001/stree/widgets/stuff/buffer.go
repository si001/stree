// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package stuff

import . "github.com/gdamore/tcell"

// Cell represents a viewable terminal cell
type Cell struct {
	Rune  rune
	Style Style
}

//type Attribute uint16
//
//
//var CellClear = tb.Cell{
//	Ch: ' ',
//	Fg: tb.ColorDefault,
//	Bg: tb.CColorDefault,
//}

// NewCell takes 1 to 2 arguments
// 1st argument = rune
// 2nd argument = optional style
//func NewCell(rune rune, args ...interface{}) tb.Cell {
//	style := StyleClear
//	if len(args) == 1 {
//		style = args[0].(Style)
//	}
//	return tb.Cell
//	{
//	Ch:
//		rune,
//			Fg: style.Fg,
//		Bg: style.Bg,
//	}
//}

//// Buffer represents a section of a terminal and is a renderable rectangle of cells.
//type Buffer struct {
//	image.Rectangle
//	CellMap map[image.Point]tb.Cell
//}
//
//func NewBuffer(r image.Rectangle) *Buffer {
//	buf := &Buffer{
//		Rectangle: r,
//		CellMap:   make(map[image.Point]Cell),
//	}
//	buf.Fill(CellClear, r) // clears out area
//	return buf
//}
//
//func (self *Buffer) GetCell(p image.Point) tb.Cell {
//	return self.CellMap[p]
//}
//
//func (self *Buffer) SetCell(c tb.Cell, p image.Point) {
//	self.CellMap[p] = c
//}
//
//func (self *Buffer) Fill(c tb.Cell, rect image.Rectangle) {
//	for x := rect.Min.X; x < rect.Max.X; x++ {
//		for y := rect.Min.Y; y < rect.Max.Y; y++ {
//			self.SetCell(c, image.Pt(x, y))
//		}
//	}
//}
//
//func (self *Buffer) SetString(s string, style Style, p image.Point) {
//	runes := []rune(s)
//	x := 0
//	for _, char := range runes {
//		self.SetCell(Cell{char, style.Fg, style.Bg}, image.Pt(p.X+x, p.Y))
//		x += rw.RuneWidth(char)
//	}
//}
