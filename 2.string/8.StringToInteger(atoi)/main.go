package main

// https://leetcode.com/problems/string-to-integer-atoi/

const MAX_INT = (1 << 31) - 1
const MIN_INT = -(1 << 31)

// each step: result = result * 10 + (s[i] - '0')
// make sure before handle result, we need to check result <= (MAX_INT - (s[i] - '0')) / 10
func myAtoi(s string) int {
	digitStartIndex := 0
	sign := 1
	for digitStartIndex < len(s) {
		if s[digitStartIndex] == ' ' {
			digitStartIndex++
			continue
		}
		if s[digitStartIndex] == '-' {
			sign *= -1
			digitStartIndex++
			break
		}

		if s[digitStartIndex] == '+' {
			digitStartIndex++
			break
		}
		break
	}

	result := 0
	for i := digitStartIndex; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}

		if result > (MAX_INT-int(s[i]-'0'))/10 {
			if sign == -1 {
				return MIN_INT
			}
			return MAX_INT
		}

		result = result*10 + int(s[i]-'0')
	}

	return result * sign
}
