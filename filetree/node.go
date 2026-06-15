package filetree

type treeNode struct {
	name     string
	isDir    bool
	children []treeNode
}

func (node treeNode) String() string {
	return node.stringFromDepth("") + "\n"
}
func (node treeNode) stringFromDepth(frontString string) string {
	current := node.name
	for i, child := range node.children {
		// 最後尾
		if len(node.children)-1 == i {
			current += "\n" + frontString + "└── " + child.stringFromDepth(frontString+"    ")
			continue
		}
		current += "\n" + frontString + "├── " + child.stringFromDepth(frontString+"│   ")
	}
	return current
}
