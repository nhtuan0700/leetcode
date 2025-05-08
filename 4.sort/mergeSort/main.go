package main

import "fmt"

// Main function that sorts arr[l...r] using
// merge()
func mergeSort(arr []int, l, r int) {
	if l < r {
		// Find the middle point
		m := l + (r-l)/2

		// Sort first and second halves
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)

		// Merge the sorted halves
		merge(arr, l, m, r)
	}
}

// Merge two sub array of arr[]
// First subarray is arr[l...m]
// Second subarray is arr[m+1...r]
func merge(arr []int, l, m, r int) {
	// Find size of two subarray to be merged
	n1 := m - l + 1
	n2 := r - m

	// Create temp arrays
	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temp arrays
	for i := 0; i < n1; i++ {
		L[i] = arr[i+l]
	}
	for i := 0; i < n2; i++ {
		R[i] = arr[i+m+1]
	}

	// Merge the temp arrays
	// Initital indices of first and second subarrays
	var i, j int
	// Initial index of merged subarray to array
	k := l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}

		k++
	}

	// Copy remaining elements of L[] if any
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	// Copy remaining elements of R[] if any
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

// - Divide the array into 2 halves
// - Sort each half
// - Then merge the sorted halves back together
// 	 + We have 2 pointers for the 2 arrays
// 	 + We compare the 2 elements
//	 + We put the smaller element into the result and increment the pointer
// - This proccess is repeated until the entire array is sorted
// TC: O(nlogn) SC: O(n)
func main() {
	arr := []int{3, 7, 1, 2}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
