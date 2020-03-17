package app

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"log"
	"math"
	"os"
	"runtime"
	"si001/stree/model"
	"si001/stree/screen"
	"si001/stree/screen/tree_list"
	"si001/stree/tools/files"
	"si001/stree/tools/files/settings"
	"si001/stree/widgets"
	"si001/stree/widgets/stuff"
	"time"
)

var defStyle tcell.Style

func ShowMain() {
	encoding.Register()
	screen.TreeAndList1 = tree_list.TreeAndList{
		List:         widgets.NewList(),
		Tree:         nil,
		FileMode:     0,
		FileMask:     "*",
		ListIsBranch: false,
		OrderBy:      model.OrderAcs | model.OrderByName,
	}

	if runtime.GOOS == "windows" {
		screen.TreeAndList1.FileMask = "*.*"
		model.PathDivider = "\\"
	} else {
		model.PathDivider = "/"
	}

	settings.ReadSettings()
	screen.TreeAndList1.Init()

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	defStyle = tcell.StyleDefault.
		Background(tcell.ColorDefault).
		Foreground(tcell.ColorLightGray)
	//tcell.StyleDefault = defStyle
	stuff.StyleClear = defStyle
	s.SetStyle(defStyle)
	s.EnableMouse()
	//s.Clear()
	//defer ui.Close()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	screen.TreeAndList1.Tree = files.LogTree(dir, widgets.NewTree())

	model.CurrentPath = files.TreeNodeToPath(screen.TreeAndList1.Tree.SelectedNode())
	//screen.HeadLeft = model.CurrentPath
	screen.TreeAndList1.ShowDir(model.CurrentPath, screen.TreeAndList1.Tree.SelectedNode(), false, false)

	model.SelectedStyle = tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite).Normal()
	w, h := s.Size()
	screen.TreeAndList1.Divider = int(math.Max(5, float64(h-model.VC_BOTTOM_HEIGHT+4)*0.25))

	tickerCount := 0
	draw(s, tickerCount)
	previousKey := ""
	//uiEvents := s.PollEvent()
	var tim *time.Timer
	callback := func() {
		var evt tcell.Event
		s.PostEvent(evt)
		tim.Reset(time.Second)
	}
	tim = time.AfterFunc(time.Second, callback)

	for !model.AppFinished {
		//select {
		//case ev := <- s.PollEvent():
		event := s.PollEvent()
		switch ev := event.(type) {
		case *tcell.EventResize:
			s.Sync()
			st := tcell.StyleDefault.Background(tcell.ColorRed)
			s.SetContent(w-1, h-1, 'R', nil, st)
		case *tcell.EventMouse:
			//x, y := ev.Position()
			//button := ev.Buttons()
			//s.SetContent(w-1, h-1, 'R', nil, st)
			processEvent(event)
		case *tcell.EventKey:
			switch {
			default:
				processEvent(event)
			}
			if previousKey == "g" {
				previousKey = ""
			} else {
				previousKey = ev.Name()
			}
		}
		//case <-ticker:
		//	tickerCount++
		//}

		draw(s, tickerCount)
	}
	settings.WriteSettings()
}
