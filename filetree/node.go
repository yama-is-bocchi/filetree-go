package filetree

type treeNode struct {
	name     string
	isDir    bool
	children []treeNode
}

func (node treeNode) String() string {
	return node.stringFromDepth(0) + "\n"
}
func (node treeNode) stringFromDepth(depth int) string {
	beforeLine := ""
	for range depth {
		beforeLine += "│   "
	}
	current := node.name
	for i, child := range node.children {
		if len(node.children)-1 == i {
			current += "\n" + beforeLine + "└── " + child.stringFromDepth(depth+1)
			continue
		}
		current += "\n" + beforeLine + "├── " + child.stringFromDepth(depth+1)
	}
	return current
}
