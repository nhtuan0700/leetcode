package main

// https://leetcode.com/problems/fibonacci-number/submissions/1626008400/

// TC: O(2^n): SC: O(n)
func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

// TC: O(n): SC: O(n)
func fib2(n int) int {
	fibos := []int{0, 1}
	for i := 2; i <= n; i++ {
		fibos = append(fibos, fibos[i-1]+fibos[i-2])
	}

	return fibos[n]
}

// TC: O(n): SC: O(n)
func fib3(n int) int {
	if n < 2 {
		return n
	}
	prevNum := 1
	num := 1

	for n > 2 {
		temp := num
		num += prevNum
		prevNum = temp
		n--
	}

	return num
}
