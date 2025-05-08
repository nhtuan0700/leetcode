package main

import "strconv"

// https://leetcode.com/problems/base-7/

// 100 / 7 = 14 (2) => "20" + "2"
// 14 / 7 = 2 (0) => "2" + "0"
// 2 / 7 = 0 (2) => "" + "2"
// 0 => ""

// - 8 / 7 = -1 (-1) => 1
// - 1 / 7 = 0 (-1)

func convertToBase7(num int) string {
	if num/7 == 0 {
		return strconv.Itoa(num)
	}

	mod := num % 7
	if mod < 0 {
		mod *= -1
	}

	return convertToBase7(num/7) + strconv.Itoa(mod)
}
