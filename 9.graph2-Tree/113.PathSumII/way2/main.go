package main

// https://leetcode.com/problems/path-sum-ii/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC: O(n), SC: O(n^2)
func pathSum(root *TreeNode, targetSum int) [][]int {
	stack := make([]int, 0)
	pathList := make([][]int, 0)

	var findPathSum func(root *TreeNode, targetSum int)

	findPathSum = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}

		stack = append(stack, root.Val)

		if root.Left == nil && root.Right == nil {
			if root.Val == targetSum {
				temp := make([]int, len(stack))
				copy(temp, stack)
				pathList = append(pathList, temp)
			}
			stack = stack[:len(stack)-1]
			return
		}

		findPathSum(root.Left, targetSum-root.Val)
		findPathSum(root.Right, targetSum-root.Val)
		stack = stack[:len(stack)-1]
	}

	return pathList
}
