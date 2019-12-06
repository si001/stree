// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package stuff

import (
	. "github.com/gdamore/tcell"
)

type Attribute uint16

var StandardColors = []Color{
	ColorRed,
	ColorGreen,
	ColorYellow,
	ColorBlue,
	ColorDarkMagenta,
	ColorDarkCyan,
	ColorWhite,
}

var StandardStyles = []Style{
	NewStyle(ColorRed),
	NewStyle(ColorGreen),
	NewStyle(ColorYellow),
	NewStyle(ColorBlue),
	NewStyle(ColorDarkMagenta),
	NewStyle(ColorDarkCyan),
	NewStyle(ColorWhite),
}

type RootTheme struct {
	Default Style

	Block BlockTheme

	//BarChart        BarChartTheme
	Gauge     GaugeTheme
	Plot      PlotTheme
	List      ListTheme
	Tree      TreeTheme
	Paragraph ParagraphTheme
	//PieChart        PieChartTheme
	//Sparkline       SparklineTheme
	StackedBarChart StackedBarChartTheme
	Tab             TabTheme
	Table           TableTheme
}

type BlockTheme struct {
	Title  Style
	Border Style
}

type BarChartTheme struct {
	Bars   []Color
	Nums   []Style
	Labels []Style
}

type GaugeTheme struct {
	Bar   Color
	Label Style
}

type PlotTheme struct {
	Lines []Attribute
	Axes  Attribute
}

type ListTheme struct {
	Text Style
}

type TreeTheme struct {
	Text      Style
	Collapsed rune
	Expanded  rune
}

type ParagraphTheme struct {
	Text Style
}

//type PieChartTheme struct {
//	Slices []Attribute
//}

//type SparklineTheme struct {
//	Title Style
//	Line  Attribute
//}

type StackedBarChartTheme struct {
	//Bars   []Attribute
	Nums   []Style
	Labels []Style
}

type TabTheme struct {
	Active   Style
	Inactive Style
}

type TableTheme struct {
	Text Style
}

// Theme holds the default Styles and Colors for all widgets.
// You can set default widget Styles by modifying the Theme before creating the widgets.
var Theme = RootTheme{
	Default: NewStyle(ColorWhite),

	Block: BlockTheme{
		Title:  NewStyle(ColorWhite),
		Border: NewStyle(ColorWhite),
	},

	//BarChart: BarChartTheme{
	//	Bars:   StandardColors,
	//	Nums:   StandardStyles,
	//	Labels: StandardStyles,
	//},

	//Paragraph: ParagraphTheme{
	//	Text: NewStyle(ColorWhite),
	//},
	//
	//PieChart: PieChartTheme{
	//	Slices: StandardColors,
	//},

	List: ListTheme{
		Text: NewStyle(ColorWhite),
	},

	Tree: TreeTheme{
		Text:      NewStyle(ColorWhite),
		Collapsed: COLLAPSED,
		Expanded:  EXPANDED,
	},

	//StackedBarChart: StackedBarChartTheme{
	//	Bars:   StandardColors,
	//	Nums:   StandardStyles,
	//	Labels: StandardStyles,
	//},

	Gauge: GaugeTheme{
		Bar:   ColorWhite,
		Label: NewStyle(ColorWhite),
	},

	//Sparkline: SparklineTheme{
	//	Title: NewStyle(ColorWhite),
	//	Line:  ColorWhite,
	//},

	//Plot: PlotTheme{
	//	Lines: StandardColors,
	//	Axes:  ColorWhite,
	//},

	Table: TableTheme{
		Text: NewStyle(ColorWhite),
	},

	Tab: TabTheme{
		Active:   NewStyle(ColorRed),
		Inactive: NewStyle(ColorWhite),
	},
}
