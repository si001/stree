package files

import (
	"si001/stree/model"
	"si001/stree/widgets"
	"strings"
)

func TreeNodeToPath(node *widgets.TreeNode) (result string) {
	for {
		result = model.PathDivider + node.Value.String() + result
		node = node.Value.(*model.Directory).Parent
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
