package main

// https://leetcode.com/problems/path-sum-ii/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	stack := make([]int, 0)
	return findPathSum(root, targetSum, stack)
}

// TC: O(n), SC: O(n^2)
func findPathSum(root *TreeNode, targetSum int, stack []int) [][]int {
	if root == nil {
		return nil
	}

	stack = append(stack, root.Val)

	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			temp := make([]int, len(stack))
			copy(temp, stack)
			res := make([][]int, 0)
			res = append(res, temp)
			return res
		}
	}

	leftPaths := findPathSum(root.Left, targetSum-root.Val, stack)
	rightPaths := findPathSum(root.Right, targetSum-root.Val, stack)

	return append(leftPaths, rightPaths...)
}
