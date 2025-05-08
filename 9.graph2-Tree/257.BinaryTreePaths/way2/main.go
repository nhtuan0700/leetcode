package main

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/binary-tree-paths/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC: O(n), SC: O(n^2)
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	paths := findPaths(root, make([]string, 0))

	res := make([]string, len(paths))
	for i, p := range paths {
		res[i] = strings.Join(p, "->")
	}

	return res
}

func findPaths(root *TreeNode, stack []string) [][]string {
	if root == nil {
		return nil
	}
	stack = append(stack, strconv.Itoa(root.Val))

	if root.Left == nil && root.Right == nil {
		res := make([][]string, 1)
		temp := make([]string, len(stack))
		copy(temp, stack)
		res[0] = temp
		return res
	}

	leftPaths := findPaths(root.Left, stack)
	rightPaths := findPaths(root.Right, stack)

	return append(leftPaths, rightPaths...)
}
