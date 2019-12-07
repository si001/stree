package app

import (
	//ui "github.com/gizak/termui/v3"
	//"github.com/nsf/termbox-go"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"time"

	"fmt"
	"log"
	"os"
	"si001/stree/files"
	"si001/stree/model"
	"si001/stree/screen"
	"si001/stree/widgets/stuff"
)

var defStyle tcell.Style

func ShowMain() {
	encoding.Register()

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	defStyle = tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.Clear()
	//defer ui.Close()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	screen.Tree1 = files.BuildTree(dir)

	model.CurrentPath = files.TreeNodeToPath(screen.Tree1.SelectedNode())
	screen.HeadLeft = model.CurrentPath
	screen.ShowDir(model.CurrentPath, screen.Tree1.SelectedNode(), false)

	model.SelectedStyle = tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite).Normal()
	w, h := s.Size()
	model.Divider = int(float32(h-2-screen.VC_BOTTOM_HEIGHT)*0.7) + 2

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

	for {
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
			case ev.Rune() == 'q':
				return
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
}

func processEvent(event tcell.Event) {
	switch screen.ViewMode {
	case screen.VM_TREEVIEW_FILES_1:
		screen.ModetreePutEvent(event)
	case screen.VM_FILELIST_1:
		screen.ModefilesPutEvent(event)
	}
}

var draw = func(s tcell.Screen, count int) {
	w, h := s.Size()
	stuff.ScreenFillBox(s, 0, 0, w, h, tcell.StyleDefault, ' ')
	switch screen.ViewMode {
	case screen.VM_TREEVIEW_FILES_1:
		screen.ModetreeDraw(s, w, h)
	case screen.VM_FILELIST_1:
		screen.ModefilesDraw(s, w, h)
	}
	s.Show()
}
