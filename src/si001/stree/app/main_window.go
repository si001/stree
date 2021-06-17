package app

import (
	ui "github.com/gizak/termui/v3"
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"si001/stree/files"
	"si001/stree/model"
	"si001/stree/screen"
	"time"
)

func ShowMain() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	//screen.FilesList1.Title = "Files"
	//screen.FilesList1.Rows = listData
	screen.FilesList1.TextStyle.Fg = ui.ColorYellow

	//screen.DriveInfo.Roqws = listData
	screen.DriveInfo.TextStyle.Fg = ui.ColorYellow

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	screen.Tree1 = files.BuildTree(dir)

	model.CurrentPath = files.TreeNodeToPath(screen.Tree1.SelectedNode())
	screen.HeadLeft = model.CurrentPath
	screen.ShowDir(model.CurrentPath, screen.Tree1.SelectedNode(), false)

	tickerCount := 0
	draw(tickerCount)
	previousKey := ""
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q":
				return
			default:
				processEvent(e)
			}
			if previousKey == "g" {
				previousKey = ""
			} else {
				previousKey = e.ID
			}
		case <-ticker:
			tickerCount++
		}

		draw(tickerCount)
	}
}

func processEvent(event ui.Event) {
	switch screen.ViewMode {
	case screen.VM_TREEVIEW_FILES_1:
		screen.ModetreePutEvent(event)
	case screen.VM_FILELIST_1:
		screen.ModefilesPutEvent(event)
	}

}

var draw = func(count int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, h := termbox.Size()

	switch screen.ViewMode {
	case screen.VM_TREEVIEW_FILES_1:
		screen.ModetreeDraw(w, h)
	case screen.VM_FILELIST_1:
		screen.ModefilesDraw(w, h)
	}
}
