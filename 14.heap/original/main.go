package main

import "fmt"

func main() {
	nums := []int{5, 3, 2, 1, 4}

	pq := []int{}

	for _, num := range nums {
		pq = push(pq, num)
	}

	for len(pq) > 0 {
		v := top(pq)
		pq = pop(pq)
		fmt.Println(v)
	}
}

func push(pq []int, num int) []int {
	pq = append(pq, num)

	n := len(pq)
	currentIndex := n - 1
	for currentIndex > 0 {
		parentIndex := (currentIndex - 1) / 2
		// max heap
		if pq[parentIndex] >= pq[currentIndex] {
			break
		}

		swap(pq, parentIndex, currentIndex)
		currentIndex = parentIndex
	}

	return pq
}

func pop(pq []int) []int {
	if isEmpty(pq) {
		return nil
	}
	pq[0] = pq[len(pq)-1]
	pq = pq[:len(pq)-1]
	n := len(pq)
	currentIndex := 0
	for {
		leftChildIndex := currentIndex*2 + 1
		rightChildIndex := currentIndex*2 + 2
		largest := currentIndex

		if leftChildIndex < n && pq[leftChildIndex] > pq[largest] {
			largest = leftChildIndex
		}

		if rightChildIndex < n && pq[rightChildIndex] > pq[largest] {
			largest = rightChildIndex
		}

		if largest == currentIndex {
			break
		}

		swap(pq, currentIndex, largest)
		currentIndex = largest
	}

	return pq
}

func top(pq []int) int {
	if isEmpty(pq) {
		panic("heap is empty") // Or return error
	}
	return pq[0]
}

func isEmpty(pq []int) bool {
	return len(pq) == 0
}

func swap(pq []int, i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
