package main

// between 2 intervals of each 2 list, we will find the intersection
// if the interval in the list have the end is smaller, we will go to the next interval of that list
//                                i
// [[0,2],[5,10],[13,23],[24,25]]
//                        j
// [[1,5],[8,12],[15,24],[25,26]]
// res = [[1,2],[5,5],[8,10],[15,23],[24,24], [25,25]]

// TC: O(N + M), O(K)
// N: len(firstList)
// M: len(secondList)
// K: number of intersections - worst case: M+N
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var i, j int
	res := make([][]int, 0)
	for i < len(firstList) && j < len(secondList) {
		// check overlap => find intersection
		if firstList[i][0] <= secondList[j][1] && firstList[i][1] >= secondList[j][0] {
			start := max(firstList[i][0], secondList[j][0])
			end := min(firstList[i][1], secondList[j][1])
			res = append(res, []int{start, end})
		}
		if firstList[i][1] > secondList[j][1] {
			j++
		} else {
			i++
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
