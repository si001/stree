package widgets

import (
	"fmt"
	"github.com/gdamore/tcell"
	"image"
	"si001/stree/widgets/stuff"
	"strings"
)

const TreeIndent = "  "

// TreeNode is a tree node.
type TreeNode struct {
	Value    fmt.Stringer
	Expanded bool
	Nodes    []*TreeNode
	// level stores the node level in the tree.
	Level int
}

type Patcher interface {
	ParsePatch(node *TreeNode) (path, value string)
}

type NodeValue struct {
	Name string
}

func (self NodeValue) String() string {
	return self.Name
}

func (self NodeValue) ParsePatch(node *TreeNode) string {
	var sb strings.Builder
	if len(node.Nodes) == 0 {
		sb.WriteString(strings.Repeat(TreeIndent, node.Level+1))
	} else {
		sb.WriteString(strings.Repeat(TreeIndent, node.Level))
		if node.Expanded {
			sb.WriteRune(stuff.Theme.Tree.Expanded)
		} else {
			sb.WriteRune(stuff.Theme.Tree.Collapsed)
		}
		sb.WriteRune(stuff.Theme.Tree.Expanded)
	}
	sb.WriteString(self.Name)
	return sb.String()
}

func (self *TreeNode) parseStyles() string {
	var sb strings.Builder
	if len(self.Nodes) == 0 {
		sb.WriteString(strings.Repeat(TreeIndent, self.Level+1))
	} else {
		sb.WriteString(strings.Repeat(TreeIndent, self.Level))
		if self.Expanded {
			sb.WriteRune(stuff.Theme.Tree.Expanded)
		} else {
			sb.WriteRune(stuff.Theme.Tree.Collapsed)
		}
		sb.WriteRune(stuff.Theme.Tree.Expanded)
	}
	sb.WriteString(self.Value.String())
	return sb.String()
}

// TreeWalkFn is a function used for walking a Tree.
// To interrupt the walking process function should return false.
type TreeWalkFn func(*TreeNode) bool

// Tree is a tree widget.
type Tree struct {
	stuff.Block
	TextStyle        tcell.Style
	SelectedRowStyle tcell.Style
	WrapText         bool
	SelectedRow      int

	nodes []*TreeNode
	// rows is flatten nodes for rendering.
	rows   []*TreeNode
	topRow int
}

// NewTree creates a new Tree widget.
func NewTree() *Tree {
	return &Tree{
		Block:            *stuff.NewBlock(),
		TextStyle:        stuff.Theme.Tree.Text,
		SelectedRowStyle: stuff.Theme.Tree.Text.Foreground(tcell.ColorWhite).Background(tcell.ColorBlue),
		WrapText:         true,
	}
}

func (self *Tree) SetNodes(nodes []*TreeNode) {
	self.nodes = nodes
	self.prepareNodes()
}

/**
*  Need to process with node.Nodes manipulate
 */
func (self *Tree) PrepareNodes() {
	self.prepareNodes()
}

func (self *Tree) prepareNodes() {
	self.rows = make([]*TreeNode, 0)
	for _, node := range self.nodes {
		self.prepareNode(node, 0)
	}
}

func (self *Tree) prepareNode(node *TreeNode, level int) {
	self.rows = append(self.rows, node)
	node.Level = level

	if node.Expanded {
		for _, n := range node.Nodes {
			self.prepareNode(n, level+1)
		}
	}
}

func (self *Tree) Walk(fn TreeWalkFn) {
	for _, n := range self.nodes {
		if !self.walk(n, fn) {
			break
		}
	}
}

func (self *Tree) walk(n *TreeNode, fn TreeWalkFn) bool {
	if !fn(n) {
		return false
	}

	for _, node := range n.Nodes {
		if !self.walk(node, fn) {
			return false
		}
	}

	return true
}

func (self *Tree) ScrollToMouse(x, y int) {
	sel := y - self.Min.Y + self.topRow - 1
	if sel < 0 {
		sel = 0
	} else if sel >= len(self.rows) {
		sel = len(self.rows) - 1
	}
	self.SelectedRow = sel
}

func (self *Tree) Draw(s tcell.Screen) {
	self.Block.Draw(s)
	point := self.Inner.Min

	// adjusts view into widget
	if self.SelectedRow >= self.Inner.Dy()+self.topRow {
		self.topRow = self.SelectedRow - self.Inner.Dy()
	} else if self.SelectedRow < self.topRow {
		self.topRow = self.SelectedRow
	}

	var nodePath string
	// draw rows
	for row := self.topRow; row < len(self.rows) && point.Y <= self.Inner.Max.Y; row++ {
		style := self.TextStyle
		if row == self.SelectedRow {
			style = self.SelectedRowStyle
		}
		node := self.rows[row]
		//if (node.Value insta)
		nodeString := node.parseStyles()
		if patcher, ok := node.Value.(Patcher); ok {
			nodePath, nodeString = patcher.ParsePatch(node)
		}
		stuff.ScreenPrintAtTween(s, point.X, point.Y, self.Inner.Max.X, self.TextStyle, style, nodePath, nodeString)
		point = image.Pt(self.Inner.Min.X, point.Y+1)
	}

	dy2 := float32(self.SelectedRow) / float32(len(self.rows))
	stuff.ScreenDrawScrolled(s, self.Inner, dy2, self.topRow > 0, len(self.rows) > int(self.topRow)+self.Inner.Dy()+1, self.BorderStyle)
}

// ScrollAmount scrolls by amount given. If amount is < 0, then scroll up.
// There is no need to set self.topRow, as this will be set automatically when drawn,
// since if the selected item is off screen then the topRow variable will change accordingly.
func (self *Tree) ScrollAmount(amount int) {
	if len(self.rows)-int(self.SelectedRow) <= amount {
		self.SelectedRow = len(self.rows) - 1
	} else if int(self.SelectedRow)+amount < 0 {
		self.SelectedRow = 0
	} else {
		self.SelectedRow += amount
	}
}

func (self *Tree) SelectedNode() *TreeNode {
	if len(self.rows) == 0 {
		return nil
	}
	return self.rows[self.SelectedRow]
}

func (self *Tree) ScrollScreenUp() {
	if self.topRow > 0 {
		self.topRow--
		if self.SelectedRow-self.topRow > self.Inner.Dy() {
			self.SelectedRow--
		}
	}
}

func (self *Tree) ScrollUp() {
	self.ScrollAmount(-1)
}

func (self *Tree) ScrollDown() {
	self.ScrollAmount(1)
}

func (self *Tree) ScrollScreenDown() {
	if self.topRow < len(self.rows)-1 {
		self.topRow++
		if self.topRow > self.SelectedRow {
			self.SelectedRow++
		}
	}
}

func (self *Tree) ScrollPageUp() {
	// If an item is selected below top row, then go to the top row.
	if self.SelectedRow > self.topRow {
		self.SelectedRow = self.topRow
	} else {
		self.ScrollAmount(-self.Inner.Dy())
	}
}

func (self *Tree) ScrollPageDown() {
	self.ScrollAmount(self.Inner.Dy())
}

func (self *Tree) ScrollHalfPageUp() {
	self.ScrollAmount(-int(stuff.FloorFloat64(float64(self.Inner.Dy()) / 2)))
}

func (self *Tree) ScrollHalfPageDown() {
	self.ScrollAmount(int(stuff.FloorFloat64(float64(self.Inner.Dy()) / 2)))
}

func (self *Tree) ScrollTop() {
	self.SelectedRow = 0
}

func (self *Tree) ScrollBottom() {
	self.SelectedRow = len(self.rows) - 1
}

func (self *Tree) Collapse() {
	self.rows[self.SelectedRow].Expanded = false
	self.prepareNodes()
}

func (self *Tree) Expand() {
	node := self.rows[self.SelectedRow]
	if len(node.Nodes) > 0 {
		self.rows[self.SelectedRow].Expanded = true
	}
	self.prepareNodes()
}

func (self *Tree) CollapseOneLevel() {
	node := self.rows[self.SelectedRow]
	for _, n := range node.Nodes {
		n.Expanded = false
	}
	self.prepareNodes()
}

func (self *Tree) ExpandRecursive() {
	node := self.rows[self.SelectedRow]
	if len(node.Nodes) > 0 {
		node.Expanded = true
		for _, n := range node.Nodes {
			expandRec(n)
		}
	}
	self.prepareNodes()
}

func expandRec(node *TreeNode) {
	node.Expanded = true
	for _, n := range node.Nodes {
		if len(n.Nodes) > 0 {
			expandRec(n)
		}
	}
}

func (self *Tree) ToggleExpand() {
	node := self.rows[self.SelectedRow]
	if len(node.Nodes) > 0 {
		node.Expanded = !node.Expanded
	}
	self.prepareNodes()
}

func (self *Tree) ExpandAll() {
	self.Walk(func(n *TreeNode) bool {
		if len(n.Nodes) > 0 {
			n.Expanded = true
		}
		return true
	})
	self.prepareNodes()
}

func (self *Tree) CollapseAll() {
	self.Walk(func(n *TreeNode) bool {
		n.Expanded = false
		return true
	})
	self.prepareNodes()
}
