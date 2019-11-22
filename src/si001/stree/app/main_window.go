package app

import (
	ui "github.com/gizak/termui/v3"
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"si001/stree/screen"
	"time"
)

func ShowMain() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	listData := []string{
		"[0] gizak/termui",
		"[1] editbox.go",
		"[2] interrupt.go",
		"[3] keyboard.go",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
		"[2] interrupt.go",
		"[3] keyboard.go",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
		"[7] nsf/termbox-go",
	}

	screen.FilesList1.Title = "Files"
	screen.FilesList1.Rows = listData
	screen.FilesList1.TextStyle.Fg = ui.ColorYellow

	screen.DriveInfo.Rows = listData
	screen.DriveInfo.TextStyle.Fg = ui.ColorYellow

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	screen.Tree1 = buildTree(dir)

	//HeadLeft.

	draw := func(count int) {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		w, h := termbox.Size()

		switch screen.ViewMode {
		case screen.VM_TREEVIEW_FILES_1:
			{
				screen.ModetreeDraw(w, h)
			}
		case screen.VM_FILELIST_1:
			{
				screen.ModefilesDraw(w, h)
			}

		}
	}

	tickerCount := 1
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
				screen.ModetreePutEvent(e)
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
