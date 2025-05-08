package main

// https://leetcode.com/problems/palindrome-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. loop head1 -> add to stack []
// loop head1 with stack popping
// 2. mid := size/2+1
// TC: O(n), SC: O(n)
func isPalindrome(head *ListNode) bool {
	stack := []int{}

	current := head
	for current != nil {
		stack = append(stack, current.Val)
		current = current.Next
	}

	i := len(stack) - 1
	current = head
	for current != nil {
		if current.Val != stack[i] {
			return false
		}
		i--
		current = current.Next
	}
	return true
}

/*
				s
	1->2->2->1->
					 f
	1->2->
			<-2<-1

				s
	1->2->3->2->1
								f
	1->2->
			<-3<-2<-1
*/
// 1 2 var prev *ListNode
// 2. mid := size/2+1
// reverse list node from mid
// more optimal than isPalindrome3
// TC: O(n), SC: O(1)
func isPalindrome2(head *ListNode) bool {
	mid := findMid(head)
	reversed := reverse(mid)
	for reversed != nil && head != nil {
		if reversed.Val != head.Val {
			return false
		}

		head = head.Next
		reversed = reversed.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func findMid(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func isPalindrome3(head *ListNode) bool {
	n := 0
	current := head
	for current != nil {
		n++
		current = current.Next
	}

	mid := n/2 - 1
	current = head // find mid node
	i := 0
	for current != nil && i <= mid {
		current = current.Next
		i++
	}

	// reverse
	var reversed *ListNode
	for current != nil {
		next := current.Next
		current.Next = reversed
		//
		reversed = current
		current = next
	}

	i = 0
	for i <= mid {
		if reversed.Val != head.Val {
			return false
		}

		i++
		head = head.Next
		reversed = reversed.Next
	}

	return true
}
