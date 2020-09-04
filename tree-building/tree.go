// Package tree builds a tree data structure from input records.
package tree

import (
	"fmt"
)

// Node is an element of the tree that may have child nodes.
type Node struct {
	ID       int
	Children []*Node
}

// Record is an item of input data used to build up a tree of nodes.
type Record struct {
	ID     int
	Parent int
}

// Build creates a tree of nodes from a list of input records.
func Build(records []Record) (*Node, error) {
	var root *Node = nil
	nodes := map[int]Node{}
	for _, record := range records {
		id, parentId := record.ID, record.Parent
		node, ok := nodes[id]
		if !ok {
			node = Node{ID: id}
			nodes[id] = node
		}
		if id == parentId {
			root = &node
			continue
		}
		parent, ok := nodes[parentId]
		if !ok {
			parent = Node{ID: parentId}
			nodes[parentId] = parent
		}
		parent.Children = append(parent.Children, &node)
		fmt.Printf("%v\n", parent)
	}
	fmt.Printf("%v\n", nodes)
	fmt.Printf("%v\n", root)
	return root, nil
}
