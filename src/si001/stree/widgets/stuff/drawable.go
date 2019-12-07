package stuff

import (
	"github.com/gdamore/tcell"
	runewidth "github.com/mattn/go-runewidth"
)

type Drawable interface {
	TCellDraw(s *tcell.Screen)
}

func ScreenPrintAt(s tcell.Screen, x, y int, style tcell.Style, str string) {
	var comb []rune
	for _, c := range str {
		comb = nil
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}
func ScreenPrintStr(s tcell.Screen, x, y int, style tcell.Style, str string) int {
	var comb []rune
	for _, c := range str {
		comb = nil
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
	return x
}
func ScreenPrintAtTween(s tcell.Screen, x, y, xTo int, style1 tcell.Style, style2 tcell.Style, str1, str2 string) {
	x = ScreenPrintStr(s, x, y, style1, str1)
	x = ScreenPrintStr(s, x, y, style2, str2)
	for i := x; i <= xTo; i++ {
		s.SetContent(i, y, ' ', nil, style2)
	}
}

func ScreenDrawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, r rune) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}
	if y1 != y2 && x1 != x2 {
		// Only add corners if we need to
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}
	if r >= 0 {
		for row := y1 + 1; row < y2; row++ {
			for col := x1 + 1; col < x2; col++ {
				s.SetContent(col, row, r, nil, style)
			}
		}
	}
}

func ScreenFillBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, r rune) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, r, nil, style)
		}
	}
}

func ScreenDrawScrolled(s tcell.Screen, x, y1, y2, y3 int, style tcell.Style) {
	for i := y1; i <= y3; i++ {
		if i == y2 {
			s.SetContent(x, i, tcell.RuneBlock, nil, style)
		} else {
			s.SetContent(x, i, tcell.RuneBoard, nil, style)
		}
	}
}
