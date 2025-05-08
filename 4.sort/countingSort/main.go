package main

import "fmt"

func countingSort(arr []int) {
	max := 0
	for _, num := range arr {
		if max < num {
			max = num
		}
	}

	countArr := make([]int, max+1)
	for _, num := range arr {
		countArr[num]++
	}

	for i := 1; i <= max; i++ {
		countArr[i] += countArr[i-1]
	}

	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		result[countArr[arr[i]] - 1] = arr[i]
		countArr[arr[i]]--
	}

	fmt.Println(result)
	// fmt.Println(countArr)
}

// https://www.geeksforgeeks.org/counting-sort/
// 1. find max value in arr
// 2. init countArr with len = max + 1
// 3. counting value of countArr by arr
// 4. Store cumulative sum (prefixSum) of the countArr
// 5. Loop arr: 
// 	result[count[arr[i]] - 1] = arr[i]; count[arr[i]]--
func main() {
	arr := []int{3, 10, 7, 20, 1, 2, 3, 10, 7}
	countingSort(arr)
	fmt.Println(arr)
}
