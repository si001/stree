package stuff

import (
	"github.com/gdamore/tcell"
)

// Color is an integer from -1 to 255
// -1 = ColorClear
// 0-255 = Xterm colors

// Modifier represents the possible modifier keys.
type Modifier tcell.ModMask

const (
	// ModifierClear clears any modifiers
	ModifierClear     Modifier = 0
	ModifierBold      Modifier = 1 << 9
	ModifierUnderline Modifier = 1 << 10
	ModifierReverse   Modifier = 1 << 11
)

// StyleClear represents a default Style, with no colors or modifiers
var StyleClear = tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorLightGray)

//Modifier: ModifierClear,
//}

// NewStyle takes 1 to 3 arguments
// 1st argument = Fg
// 2nd argument = optional Bg
// 3rd argument = optional Modifier
func NewStyle(fg tcell.Color, args ...interface{}) tcell.Style {
	bg := tcell.ColorDefault
	//modifier := ModifierClear
	if len(args) >= 1 {
		bg = args[0].(tcell.Color)
	}
	//if len(args) == 2 {
	//	modifier = args[1].(Modifier)
	//}
	return tcell.StyleDefault.Background(bg).Foreground(fg)
}
