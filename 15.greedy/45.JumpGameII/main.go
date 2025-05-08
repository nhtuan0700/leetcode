package main

// https://leetcode.com/problems/jump-game-ii/description/

// Greedy: At each index, choose the next jump range that reaches as far as possible
// In other words, the jump range corresponds to the current step.

// Approach
// - At each index, update the farthest reachable index for the next step.
// - When we reach the end of the current range, increment the jump count and update the range to the farthest reachable

// 2, 3, 1, 1, 4
// 0, 1, 1, 2, 2
//         Range
// step 0: 0 - 0
// step 1: 1 - 2
// step 2: 3 - 4

// step i: start_i -> end_i
//  for u: start_id -> end_i
//      - end_(i+1)=max(u + nums[u])
// step i+1: end_i + 1 -> end_(i+1)

// TC: O(n), SC: O(1)
func jump(nums []int) int {
	n := len(nums)
	cnt := 0
	curEnd := 0
	maxPos := 0
	i := 0

	for i < n-1 {
		// calculate maxPos for nextStep
		maxPos = max(maxPos, i+nums[i])
		if i == curEnd {
			cnt++
			curEnd = maxPos
		}
		i++
	}

	return cnt
}
