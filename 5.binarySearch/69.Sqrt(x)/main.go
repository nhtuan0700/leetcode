package main

// https://leetcode.com/problems/sqrtx/

// x: 4 -> 2, 16 -> 4
// x: 8 -> 2

// largest <= num
func mySqrt1(x int) int {
	l := 1
	h := x
	res := 0
	for l <= h {
		m := l + (h-l)/2
		square := m * m
		if square == x {
			return m
		}

		if square <= x {
			res = m
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return res
}

// x: 4 -> 2, 16 -> 4
// x: 8 -> 2

// largest <= num
func mySqrt2(x int) int {
	l := 1
	h := x
	for l <= h {
		m := l + (h-l)/2
		square := m * m
		if square == x {
			return m
		}

		if square < x {
			l = m + 1
		} else {
			h = m - 1
		}
	}

    return h
}
