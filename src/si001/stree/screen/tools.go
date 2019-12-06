package screen

import (
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

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
