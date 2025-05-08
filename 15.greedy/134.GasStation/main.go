package main

// https://leetcode.com/problems/gas-station/

// 2, 3
// 4, 2
// gas:  1,2,3,4,5
// cost: 3,4,5,1,2
// diff: -2,-2,-2,3,3
// Greedy:
// - sum(gas) >= sum(cost) -> there is exactly one solution
// - Iterate from left to the right:
//   - keep track currentTank = gas[i] - cost[i]
//   - If at any point with current tank < 0, it means at that point we can move to the next one, so try starting from the next position and reset current tank = 0

// TC: O(n), SC: O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	totalGas := 0
	totalCost := 0
	tank := 0
	start := 0
	for i := 0; i < n; i++ {
		totalGas += gas[i]
		totalCost += cost[i]

		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}

	return start
}
