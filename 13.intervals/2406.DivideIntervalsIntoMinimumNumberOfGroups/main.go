package main

// https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/
import "sort"

// Mỗi khi có một interval bắt đầu → ta coi như mở một phòng → +1.
// Mỗi khi interval kết thúc → đóng phòng lại → -1.
// Tổng số phòng đang mở tại một thời điểm = số interval đang hoạt động (overlap).
// [1,4] [2,3] [3,4] => 3

// TC: O(nlogn), SC: O(n)
func minGroups(intervals [][]int) int {
	n := len(intervals)
	events := make([][2]int, 2*n)

	for i, in := range intervals {
		indexOfEvent := i * 2
		events[indexOfEvent] = [2]int{in[0], 1}    // open
		events[indexOfEvent+1] = [2]int{in[1], -1} // close
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] > events[j][1] // open must be before close
		}
		return events[i][0] < events[j][0]
	})

	opening := 0
	overlap := 0

	for _, e := range events {
		opening += e[1]
		overlap = max(overlap, opening)
	}

	return overlap
}
