package main

import "fmt"

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func partition(arr []int, l, r int) int {
	// choosing the pivot
	pivot := arr[r]

	// Index of next smaller element position
	i := l

	for j := l; j < r; j++ {
		// if the current element is smaller than the pivot
		if arr[j] < pivot {
			// swap this element to the current index
			swap(arr, i, j)
			i++
		}
	}

	swap(arr, i, r)

	return i
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	// pI is partioning index
	pI := partition(arr, l, r)

	// Separately sort elements before parition and after partition
	quickSort(arr, l, pI-1)
	quickSort(arr, pI+1, r)
}

// - Array is divided into subarrays by selecting the pivot element (element selected from the array)
// - Position the pivot in such a way that:
//   + elements less than pivot are kept on the left side
//   + elements greater than pivot are kept on the right side
// - the left and right subarrays are also divdied using the same approach. This process continues until each subarray contains a single elemnts
// - At this point, elements are already sorted. Finally, elments are combined to form a sorted array
// * partition: 
// - keep a pointer pointing to the pivot element (last element)
// - Keep another pointer "i" keep track of the smaller elements of pivot
// - Iterate from left to right
// - If we encounter elements < pivot, swap the current element with the one in index "i"
//   Then increment i
// - Finally, we swap the pivot element with the element in index "i"
// TC: Best/Average: O(nlogn), worst: O(n^2)
// SC: O(n)
func main() {
	arr := []int{3, 10, 7, 20, 1, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
