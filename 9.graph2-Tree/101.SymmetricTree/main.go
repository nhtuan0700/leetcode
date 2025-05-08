package main

// https://leetcode.com/problems/symmetric-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

// TC: O(n), SC: O(n)
func isMirror(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && isMirror(node1.Left, node2.Right) && isMirror(node1.Right, node2.Left)
}
