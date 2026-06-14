package filetree

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type treeData struct {
	name     string
	isDir    bool
	children []treeData
}

func parseTreeData(targetPath string, depth int) (treeData, error) {
	children, err := walkDirAndParseTreeData(targetPath, depth)
	rootTreeData := treeData{name: ".", isDir: true, children: children}
	return rootTreeData, err
}

func walkDirAndParseTreeData(path string, depth int) ([]treeData, error) {
	if depth < 0 {
		return []treeData{}, errors.New("invalid depth. please set it to a value greater than 0")
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %w", err)
	}
	var result []treeData
	for _, entry := range entries {
		node := treeData{name: entry.Name(), isDir: entry.IsDir()}

		if entry.IsDir() && depth > 0 {
			children, err := walkDirAndParseTreeData(filepath.Join(path, entry.Name()), depth-1)
			if err != nil {
				return nil, fmt.Errorf("failed to walk dir and parse tree data: %w", err)
			}
			node.children = append(node.children, children...)
		}
		result = append(result, node)
	}
	return result, nil
}
