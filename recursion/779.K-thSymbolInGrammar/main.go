package main

import (
	"fmt"
	"math"
)

// S1 = "0"
// S2 = "01"
// S3 = "0110"
// S4 = "01101001"
// S5 = "0110100110010110"

// n = 5, k = 8, mid = 9, symmetricK = k - mid + 1
func kthGrammar1(n int, k int) int {
	length := 1 << (n - 1)
	mid := length/2 + 1

	if n == 1 {
		return 0
	}

	if k < mid {
		return kthGrammar1(n-1, k)
	}

	return 1 - kthGrammar1(n-1, k-mid+1)
}

func kthGrammar2(n int, k int) int {
	len := int(math.Pow(2, float64(n-1)))

	midPoint := len / 2
	if n == 1 {
		return 0
	}

	symetryPos := k
	if k > midPoint {
		symetryPos = k - midPoint
	}

	preResult := kthGrammar2(n-1, symetryPos)

	if k > midPoint {
		if preResult == 0 {
			preResult += 1
		} else {
			preResult -= 1
		}
	}

	fmt.Println("k: ", k)
	fmt.Println("preResult: ", preResult)

	return preResult
}
