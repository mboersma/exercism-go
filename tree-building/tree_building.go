// Package tree builds a tree data structure from input records.
package tree

import (
	"fmt"
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
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for i, record := range records {
		id, parentId := record.ID, record.Parent
		if i != id {
			return nil, fmt.Errorf("nodes must be continuous")
		}
		if id == parentId && id != 0 {
			return nil, fmt.Errorf("non-root node cannot have itself as parent")
		}
		_, ok := nodes[id]
		if ok {
			return nil, fmt.Errorf("duplicate node")
		}
		node := &Node{ID: id}
		nodes[id] = node
		if id == 0 {
			if parentId != 0 {
				return nil, fmt.Errorf("root node cannot have a parent")
			}
			root = node
			continue
		}
		parent, ok := nodes[parentId]
		if !ok {
			parent = &Node{ID: parentId}
			nodes[parentId] = parent
		}
		parent.Children = append(parent.Children, node)
	}
	if root == nil && len(nodes) > 0 {
		return nil, fmt.Errorf("no root node")
	}
	return root, nil
}
