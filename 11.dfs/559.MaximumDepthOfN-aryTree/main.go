package main

import "fmt"

// https://leetcode.com/problems/maximum-depth-of-n-ary-tree/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

// each node, we will have max depth of children
// then return maxDepth + 1

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	d := 0
	for _, child := range root.Children { // child - (maxDepth(3) -> maxDepth(5) => 1)
		d = max(d, maxDepth(child))
		fmt.Println(child.Val, d)
	}

	return d + 1
}
