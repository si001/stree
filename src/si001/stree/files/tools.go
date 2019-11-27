package files

import (
	"github.com/gizak/termui/v3/widgets"
	"si001/stree/model"
	"strings"
)

func TreeNodeToPath(node *widgets.TreeNode) (result string) {
	for {
		result = model.PathDivider + node.Value.String() + result
		node = node.Value.(model.Directory).Parent
		if node == nil {
			break
		}
	}

	switch {
	case strings.HasPrefix(result, "///"):
		result = result[2:]
	case strings.HasPrefix(result, "//") || strings.HasPrefix(result, "\\"):
		result = result[1:]
	}
	return result
}
