package filetree

import (
	"fmt"
	"io"
)

type fileTree struct {
	targetPath string
	depth      int
}

func New(targetPath string, depth int) fileTree {
	return fileTree{
		targetPath: targetPath,
		depth:      depth,
	}
}

func (tree fileTree) WriteTo(writer io.Writer) (int64, error) {
	// ターゲットのパスを巡回しパース
	node, err := parseTreeNode(tree.targetPath, tree.depth)
	if err != nil {
		return 0, fmt.Errorf("failed to parse tree data: %w", err)
	}
	n, err := writer.Write([]byte(node.String() + "\n"))
	return int64(n), err
}
