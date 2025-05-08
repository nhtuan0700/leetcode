package main

// https://leetcode.com/problems/minimum-time-to-make-rope-colorful/description/

// problem: remove same consecutive ballons and keep only one from them that make sure the removed time is minimum
// solution:
// - keep the maxTime from the same consecutive ballons
// - if currentTime < maxtime
//   - true:  removing current ballon (removedTime += currentTime)
//   - false: removedTime += maxTime, update maxTime = currentTime
//
// greedy: remove ballons that have time smaller than max time and keep maxTime <=> sum - tmax
// we assume there is any solution that sum - ti <= sum - tmax <=> ti >= tmax

// Giả sử phản chứng:
// - Gọi G là lời giải của greedy: giữ lại quả bóng có thời gian xóa lớn nhất tmax, tổng chi phí là:
// cost(G) = sum - tmax
// - Giả sử tồn tại lời giải O khác tốt hơn: giữ lại bóng có thời gian ti < tmax, tổng chi phí là:
// cost(O) = sum - ti
// - Ta có cost(O) < cost(G)
// 	⟹ Vì ti < tmax ⇒ cost(O) > cost(G)
// 	⟹ Mâu thuẫn với giả sử O tối ưu hơn.

// TC: O(N), SC: O(1)
func minCost(colors string, neededTime []int) int {
	res := 0
	n := len(colors)
	for i := 0; i < n; {
		j := i
		sum := neededTime[j]
		maxTime := neededTime[j]
		for j+1 < n && colors[j] == colors[j+1] {
			j++
			sum += neededTime[j]
			maxTime = max(maxTime, neededTime[j])
		}

		res += sum - maxTime
		i = j + 1
	}
	return res
}

func minCost1(colors string, neededTime []int) int {
	res := 0
	maxTime := 0
	// 3 4 2
	for i := 0; i < len(colors)-1; i++ {
		if colors[i] == colors[i+1] {
			if maxTime == 0 {
				maxTime = neededTime[i]
			}
			if maxTime < neededTime[i+1] {
				res += maxTime
				maxTime = neededTime[i+1]
			} else {
				res += neededTime[i+1]
			}
		} else {
			maxTime = 0
		}
	}
	return res
}

// this section is shorter than minCost1
func minCost2(colors string, neededTime []int) int {
	res := 0
	maxTime := neededTime[0]
	// 3 4 2
	for i := 1; i < len(colors); i++ {
		if colors[i] == colors[i-1] {
			res += min(maxTime, neededTime[i])
			maxTime = max(maxTime, neededTime[i])
		} else {
			maxTime = neededTime[i]
		}
	}
	return res
}
