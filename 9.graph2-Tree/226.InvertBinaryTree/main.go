package main

// https://leetcode.com/problems/invert-binary-tree/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC: O(n), SC: O(N)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
