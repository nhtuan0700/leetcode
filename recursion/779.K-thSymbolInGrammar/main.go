package main

// S1 = "0"
// S2 = "01"
// S3 = "0110"
// S4 = "01101001"
// S5 = "0110100110010110"

// n = 5, k = 8, mid = 9, symmetricK = k - mid + 1
func kthGrammar(n int, k int) int {
	length := 1 << (n - 1)
	mid := length/2 + 1

	if n == 1 {
		return 0
	}

	if k < mid {
		return kthGrammar(n-1, k)
	}

	return 1 - kthGrammar(n-1, k-mid+1)
}
