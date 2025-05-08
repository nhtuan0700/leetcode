package main

// https://leetcode.com/problems/calculate-digit-sum-of-a-string/

import "strconv"

// 11111222223
// 0, 3, 6

func digitSum(s string, k int) string {
	result := s

	for len(result) > k {
		temp := ""

		for i := 0; i < len(result); i++ {
			j := i

			sum := 0
			for {
				sum = sum + int(result[j]-'0')
				if (j+1)%k == 0 || j+1 == len(result) {
					break
				}
				j++
			}

			temp += strconv.Itoa(sum)
			i = j
		}

		result = temp

	}

	return result
}
