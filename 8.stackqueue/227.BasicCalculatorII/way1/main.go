package main

// https://leetcode.com/problems/basic-calculator-ii/description/

/*
    i
3+2*2

step1: handle *,/
nums: [3,4]
ops: [+]

step2: handle +,-
*/

// TC: O(n), SC: O(2n) = O(n)
func calculate(s string) int {
	nums := []int{}
	ops := []byte{}

	// handle *,/
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ' ' {
			continue
		}
		if c == '+' || c == '-' || c == '*' || c == '/' {
			ops = append(ops, c)
			continue
		}

		num := 0
		for j := i; j < len(s); j++ {
			if s[j] >= '0' && s[j] <= '9' {
				num = num*10 + int(s[j]-'0')
				i = j
				continue
			}
			break
		}

		if len(ops) > 0 && (ops[len(ops)-1] == '*' || ops[len(ops)-1] == '/') {
			lastOp := ops[len(ops)-1]
			lastNum := nums[len(nums)-1]
			nums[len(nums)-1] = calc(lastNum, num, lastOp)
			ops = ops[:len(ops)-1]
			continue
		}

		nums = append(nums, num)
	}

	res := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		res = calc(res, nums[i+1], ops[i])
	}
	return res
}

func calc(num1, num2 int, op byte) int {
	switch op {
	case '+':
		return num1 + num2
	case '-':
		return num1 - num2
	case '*':
		return num1 * num2
	case '/':
		return num1 / num2
	}

	return 0
}
