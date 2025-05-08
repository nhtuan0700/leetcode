package main

// https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/

// S1 = "0"
// S2 = "011"
// S3 = "0111001" 7 => 1, 6 => 2, 5 => 3
// S4 = "011100110110001"
// 0111001
// 0110001

// lenghth of n = 2^n - 1

// biến đổi về một nửa bên phải, nếu k > mid thì invert + reverse
// mid = length / 2 + 1
// mid == k => return
// k < mid => do nothing
// else => reverse + invert, k = symmetry of k, invert=!invert

// n = 3, k = 6
// f(n, invert)
// f(3, false): k = 7, length = 7, mid = 4, k = 1
// f(2, true): k = 0, length = 3, mid = 1, k = 0
// f(1, true): k = 0, n = 1 => invert ? 1 : 0

// n = 3, k = 5
// f(n, invert)
// f(3, false): k = 5, length = 7, mid = 4, k = 3, invert = true
// f(2, true): k = 3, length = 3, mid = 1, k = 1, invert = false
// f(1) => 0

// n = 4, k = 4
// f(n, invert)
// f(3, false): k = 4, length = 7, mid = 3, k = 2
// f(2, true): k = 1, length = 3, mid = 1 => invert ? 0 : 1

// recursion n -1 until (n == 1 => 0)
// key: mỗi ký tự tại k của ở 1 nửa sẽ đối nghịch tương ứng với nửa còn lại

func findKthBit(n int, k int) byte {
	return runFindKthBit(n, k, false)
}

func runFindKthBit(n int, k int, invert bool) byte {
	length := (1 << n) - 1
	mid := length/2 + 1

	if n == 1 {
		if invert {
			return '1'
		}

		return '0'
	}

	if mid == k {
		if invert {
			return '0'
		}

		return '1'
	}

	// k -= mid
	// k = mid - k
	// k = mid - (k - mid) <=> k = 2mid - k <=> k = len + 1 - k

	if k > mid {
		k = length - k + 1
		invert = !invert
	}

	return runFindKthBit(n-1, k, invert)
}
