package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC: O(n), SC: O(n)
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return height(root) != -1
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := height(root.Left)
	right := height(root.Right)
	if left == -1 || right == -1 {
		return -1
	}

	if abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
