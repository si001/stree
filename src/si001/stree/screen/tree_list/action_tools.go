package tree_list

import (
	"si001/stree/model"
	"si001/stree/screen/botton_box/actions"
)

func actionQuit() {
	actions.RequestYNMessageBoxCenter("Do you really want to exit?", func(result bool) {
		if result {
			model.AppFinished = true
		}
	})
}
