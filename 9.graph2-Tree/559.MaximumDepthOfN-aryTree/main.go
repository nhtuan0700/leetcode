package main

// https://leetcode.com/problems/maximum-depth-of-n-ary-tree/

type Node struct {
	Val      int
	Children []*Node
}

// Find depth of the subtree rooted at the current node.
// => depth of current node = max depth among its children + 1

// TC: O(n), SC: O(n)
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	d := 0 // depth
	for _, child := range root.Children {
		d = max(d, maxDepth(child))
	}

	// +1 to count the current node itself
	return d + 1
}
