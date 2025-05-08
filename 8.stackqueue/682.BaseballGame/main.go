package main

import (
	"container/list"
	"strconv"
)

// https://leetcode.com/problems/baseball-game/
// TC: O(n), SC: O(n)
func calPoints(operations []string) int {
	stack := list.New()

	for _, o := range operations {
		switch o {
		case "+":
			num1 := stack.Back().Value.(int)
			num2 := stack.Back().Prev().Value.(int)
			stack.PushBack(num1 + num2)
		case "D":
			num := stack.Back().Value.(int)
			stack.PushBack(num * 2)
		case "C":
			stack.Remove(stack.Back())
		default:
			num, _ := strconv.Atoi(o)
			stack.PushBack(num)
		}
	}

	res := 0
	for stack.Back() != nil {
		res += stack.Back().Value.(int)
		stack.Remove(stack.Back())
	}

	return res
}
