package main

// https://leetcode.com/problems/middle-of-the-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// idea2: use 2 pointer
// p1: 1 step at a time
// p2: 2 steps at a time
// <=> p2's speed = 2x p1's speed
// when p2 is at the end of the linked list => stop
// TC: O(n/2) is O(n)
func middleNode1(head *ListNode) *ListNode {
	slow, fast := head, head

	//       p1
	// 1->2->3->4->5->
	//                p2 
	//          p1
	// 1->2->3->4->5->6->
	// 						 		   p2
	// fast == nil when the size is odd
	// fast.Next == nil when the size is even
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// TC: O(n) n: size of node
func middleNode2(head *ListNode) *ListNode {
	currentNode := head
	length := 0

	for currentNode != nil {
		length++
		currentNode = currentNode.Next
	}

	mid := 0
	currentNode = head
	for mid != length/2 {
		mid++
		currentNode = currentNode.Next
	}

	return currentNode
}
