package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func sort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
}

// - Traverse from left and compare adjecent elements and the highter one is placed at right side
// - In this way, the largest element is moved to the rightmost end at first
// - This process is then continued to find the second largest and place it and so on until the data is sorted
// TC: O(n^2), SC: O(n)
func main() {
	arr := []int{3, 7, 1, 2, 5, 4}
	sort(arr)
	fmt.Println(arr)
}
