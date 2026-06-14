package filetree

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type treeNode struct {
	name     string
	isDir    bool
	children []treeNode
}

func (node treeNode) String() string {
	return node.stringFromDepth(0)
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

func parseTreeNode(targetPath string, depth int) (treeNode, error) {
	children, err := walkDirAndParseTreeNode(targetPath, depth)
	rootTreeNode := treeNode{name: ".", isDir: true, children: children}
	return rootTreeNode, err
}

func walkDirAndParseTreeNode(path string, depth int) ([]treeNode, error) {
	if depth < 0 {
		return []treeNode{}, errors.New("invalid depth. please set it to a value greater than 0")
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %w", err)
	}
	var result []treeNode
	for _, entry := range entries {
		node := treeNode{name: entry.Name(), isDir: entry.IsDir()}

		if entry.IsDir() && depth > 0 {
			children, err := walkDirAndParseTreeNode(filepath.Join(path, entry.Name()), depth-1)
			if err != nil {
				return nil, fmt.Errorf("failed to walk dir and parse tree node: %w", err)
			}
			node.children = append(node.children, children...)
		}
		result = append(result, node)
	}
	return result, nil
}
