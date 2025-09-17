package visualizer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Run(root string, depth int, showAll bool, jsonOut bool) error {
	if jsonOut {
		return printJSON(root, depth, showAll)
	}
	return printTree(root, "", depth, showAll, 0)
}

func printTree(path, prefix string, maxDepth int, showAll bool, level int) error {
	if maxDepth != -1 && level > maxDepth {
		return nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for i, entry := range entries {
		if !showAll && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		connector := "├── "
		nextPrefix := prefix + "│   "
		if i == len(entries)-1 {
			connector = "└── "
			nextPrefix = prefix + "    "
		}

		fmt.Println(prefix + connector + entry.Name())

		if entry.IsDir() {
			subPath := filepath.Join(path, entry.Name())
			printTree(subPath, nextPrefix, maxDepth, showAll, level+1)
		}
	}
	return nil
}

type Node struct {
	Name     string `json:"name"`
	IsDir    bool   `json:"is_dir"`
	Children []Node `json:"children,omitempty"`
}

func buildTree(path string, depth int, showAll bool, level int) (Node, error) {
	info, err := os.Stat(path)
	if err != nil {
		return Node{}, err
	}

	node := Node{Name: info.Name(), IsDir: info.IsDir()}

	if info.IsDir() {
		if depth != -1 && level > depth {
			return node, nil
		}

		entries, err := os.ReadDir(path)
		if err != nil {
			return node, err
		}

		for _, entry := range entries {
			if !showAll && strings.HasPrefix(entry.Name(), ".") {
				continue
			}
			childPath := filepath.Join(path, entry.Name())
			child, err := buildTree(childPath, depth, showAll, level+1)
			if err == nil {
				node.Children = append(node.Children, child)
			}
		}
	}

	return node, nil
}

func printJSON(root string, depth int, showAll bool) error {
	tree, err := buildTree(root, depth, showAll, 0)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(tree)
}
