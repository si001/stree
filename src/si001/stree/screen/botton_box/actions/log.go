package actions

import (
	"fmt"
	"si001/stree/model"
	"si001/stree/screen/botton_box"
	"si001/stree/screen/botton_box/box_tools"
	"si001/stree/widgets"
)

//type LogBox struct {
//	box_tools.HistoryTool
//}
//
//func RequestLog(callback func(string)) {
//	box := &LogBox{
//		HistoryTool: box_tools.HistoryTool{
//			HistoryId: "log",
//			Width:     40,
//		},
//	}
//	box.Callback = func(res *string) {
//		if res != nil {
//			model.LastEvent = *res
//			callback(*res)
//		}
//		botton_box.NormalBottomBox()
//	}
//	box.Init()
//	model.BottomMode = box
//}

type BoxLog struct {
	box_tools.BoxEditor
	logCallback func(mask string)
}

//func (box *BoxLog) Draw(s tcell.Screen) {
//
//}

func RequestLog(roots []string, callback func(string)) {
	box := BoxLog{
		BoxEditor: box_tools.BoxEditor{
			InterfaceText1: "Log path: ",
			InterfaceText2: "",
			Text:           "",
			Cursor:         0,
			HistoryWidth:   80,
			HistoryId:      "log",
		},
		logCallback: callback,
	}
	model.BottomMode = &box
	box.Callback = func(path *string) {
		botton_box.NormalBottomBox()
		if path != nil {
			box.logCallback(*path)
		}
	}
	box.ShowHistory()
	list := box.History.GetList()
	var values []*fmt.Stringer
	for _, elt := range list.Rows {
		e := (*elt).String()
		exist := false
		for _, r := range roots {
			if e == r {
				exist = true
				break
			}
		}
		if !exist {
			var v fmt.Stringer = widgets.SimpleStringer{Value: " " + e}
			values = append(values, &v)
		}
	}
	for _, r := range roots {
		var v fmt.Stringer = widgets.SimpleStringer{Value: " " + r}
		values = append(values, &v)
	}
	list.Rows = values
}
