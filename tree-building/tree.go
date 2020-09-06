// Package tree builds a tree data structure from input records.
package tree

import (
	"sort"
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
	nodes := map[int]*Node{}
	for _, record := range records {
		id, parentId := record.ID, record.Parent
		node, ok := nodes[id]
		if !ok {
			node = &Node{ID: id}
			nodes[id] = node
		}
		if id == parentId {
			root = node
			continue
		}
		parent, ok := nodes[parentId]
		if !ok {
			parent = &Node{ID: parentId}
			nodes[parentId] = parent
		}
		parent.Children = append(parent.Children, node)
		sort.Slice(parent.Children, func(i, j int) bool {
			return parent.Children[i].ID < parent.Children[j].ID
		})
	}
	return root, nil
}
