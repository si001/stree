package model

import (
	"github.com/gdamore/tcell"
)

type InfoString struct {
	Name           string
	FStyle         tcell.Style
	FSelectedStyle tcell.Style
}

func (self InfoString) String() string {
	return self.Name
}

func (self InfoString) FormattedString(stl, width int) string {
	return self.Name
}

func (self InfoString) Style() tcell.Style {
	return self.FStyle
}

func (self InfoString) SelectedStyle() tcell.Style {
	return self.FSelectedStyle
}
