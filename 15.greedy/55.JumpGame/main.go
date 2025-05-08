package main

// https://leetcode.com/problems/jump-game/

// [2,2,2,0,4]
// idea: at each step, we will keep the maximum jump into the the last index
// Greedy:
// - at each step, keep track of the farthest reachable index
// - if current index is within farthest, upadate farthest
// - if at any point, we can reach index i return false

// jump = nums[0]
// for i: 1 -> n - 1:
// - check jump > i && i < n:
//   - true: jump = max(jump, i + nums[i])
//   - false: return jump >= n - 1

// TC: O(n), SC: O(1)
func canJump(nums []int) bool {
	n := len(nums)
	maxPos := nums[0]
	for i := 1; i < n; i++ {
		if maxPos >= n-1 || i > maxPos {
			break
		}
		maxPos = max(maxPos, i+nums[i])
	}

	return maxPos >= n-1
}

func canJump2(nums []int) bool {
	n := len(nums)
	maxPos := 0
	for i := 0; i < n-1; i++ {
		maxPos = max(maxPos, i+nums[i])
		if i == maxPos {
			return false
		}
	}

	return true
}
