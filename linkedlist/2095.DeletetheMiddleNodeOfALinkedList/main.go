package main

// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

// size: [1, 105]
// 
func deleteMiddle1(head *ListNode) *ListNode {
	if head.Next == nil { // edge case
		return nil
	}

	slow, fast := head, head
	prev := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}

	prev.Next = slow.Next
	slow = nil

	return head
}
