package main

// https://leetcode.com/problems/valid-parentheses/description/

import "container/list"

// push the bracket into stack if the current bracket is not close
// Then we need to pop the stack and check whether the current close bracket is valid corresponding open one
// Dry run
// s := "([])"
// 1. Push '(' -> stack: ['(']
// 2. Push '[' -> stack: ['(', ']']
// 3. Pop ']' -> stack: ['(']
// 4. Pop ')' -> stack: []

// TC: O(n), SC: O(n)

// Linked list
// slice is better than linked list because push/pop in slice is O(1) and slice has more optimal memory
func isValid1(s string) bool {
	brackets := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := list.New()
	for _, c := range []byte(s) {
		// close bracket
		if open, ok := brackets[c]; ok {
			node := stack.Back()
			if node == nil || node.Value != open { // no open bracket previous
				return false
			}
			stack.Remove(node)
		} else { // open bracket
			stack.PushBack(c) // push to top of stack
		}
	}

	return stack.Len() == 0 // > 0 => still have open brackets
}

// using slice
func isValid2(s string) bool {
	brackets := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]byte, 0)
	for _, c := range []byte(s) {
		// close bracket
		if open, ok := brackets[c]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != open { // no open bracket previous
				return false
			}
			// pop
			stack = stack[:len(stack)-1]
		} else { // open bracket
			stack = append(stack, c) // push to top of stack
		}
	}

	return len(stack) == 0
}

// store opposite brackets rather than 2 above
func isValid3(s string) bool {
	brackets := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]byte, 0)
	for _, c := range []byte(s) {
		// open bracket
		if close, ok := brackets[c]; ok {
			stack = append(stack, close) // // push close bracket to top of stack
		} else { // close bracket
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			}
			// pop
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isValid4(s string) bool {
	stack := list.New()
	for _, c := range s {
		switch c {
		case '(':
			stack.PushBack(')')
		case '{':
			stack.PushBack('}')
		case '[':
			stack.PushBack(']')
		default:
			if stack.Back() == nil || stack.Remove(stack.Back()) != c {
				return false
			}
		}
	}

	return stack.Len() == 0
}

// using recursion
func isValid5(s string) bool {
	return run(s, 0, make([]rune, 0))
}

func run(s string, i int, stack []rune) bool {
	if i == len(s) {
		return len(stack) == 0
	}

	// open bracket
	if s[i] == '(' || s[i] == '[' || s[i] == '{' {
		switch s[i] {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		}

		return run(s, i+1, stack)
	}

	// close bracket
	if len(stack) == 0 || stack[len(stack)-1] != rune(s[i]) {
		return false
	}

	// pop stack
	return run(s, i+1, stack[:len(stack)-1])
}
