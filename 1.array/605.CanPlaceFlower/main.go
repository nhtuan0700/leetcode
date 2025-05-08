package main

// https://leetcode.com/problems/can-place-flowers/description/

// [1, 0, 0], [0]
// TC: O(n)

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if n == 0 {
			return true
		}

		if flowerbed[i] == 1 {
			i++
			continue
		}

		// pre, next Plot
		prePlot := max(i-1, 0)
		nextPlot := min(i+1, len(flowerbed)-1)
		if flowerbed[nextPlot] == 1 || flowerbed[prePlot] == 1 {
			continue
		}

		flowerbed[i] = 1
		i++
		n--
	}

	return n == 0
}

// ====2====
// [1,0,0,0,1]

// Suy ngược: cần tìm tổng tối đa trồng được bao nhiêu bông hoa ở các diện tích được phép trồng
// => maximum = (n + 1) / 2
// TC: O(n)
func maximumPlaceNewFlowers(flowerbed []int) int {
	result := 0

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}

		j := i
		// find sub array containing only value 0
		for j+1 < len(flowerbed) && flowerbed[j+1] == 0 {
			j++
		}

		totalPlacableRange := j - i + 1
		if i > 0 {
			totalPlacableRange--
		}

		if j+1 < len(flowerbed) {
			totalPlacableRange--
		}

		result += (totalPlacableRange + 1) / 2
		i = j
	}

	return result
}

func canPlaceFlowers2(flowerbed []int, n int) bool {
	return maximumPlaceNewFlowers(flowerbed) >= n
}
