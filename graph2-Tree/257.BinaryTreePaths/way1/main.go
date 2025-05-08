package main

import "strconv"

// https://leetcode.com/problems/binary-tree-paths/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC: O(N), SC: O(N)
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	return findPaths(root, "")
}

func findPaths(root *TreeNode, path string) []string {
	if root == nil {
		return nil
	}

	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil { // leaf node
		return []string{path}
	}

	path += "->"
	leftPaths := findPaths(root.Left, path)
	rightPaths := findPaths(root.Right, path)

	return append(leftPaths, rightPaths...)
}
