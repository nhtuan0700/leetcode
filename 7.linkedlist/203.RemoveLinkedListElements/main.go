package main

// https://leetcode.com/problems/remove-linked-list-elements/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
prev := Head
c := h

h
1 2 6 3 4 5 6
c

h
7 7 7
c
  - stop c == nil || c.Next == nil
    if c.Next.Val == val && current.Next.Val == Val => c.Next = c.Next.Next
*/

// - step 1: move head if head.Val == val
// - step 2: keep track current, move current
// TC: O(n), SC: O(1)
func removeElements1(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	prev := &ListNode{Next: head}
	current := head
	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
		} else {
			prev = current
		}

		current = current.Next
	}

	return head
}

/*
h
1 2 6 3 4 5 6

1->removeElements([2,6,3,4,5,6])
    2->removeElements([6,3,4,5,6])
        6->removeElements([3,4,5,6]) // remove 6
            3->removeElements([4,5,6])
                4->removeElements([5,6])
                    5->removeElements([6])
                        6->removeElements(nil) // remove 6
                            nil => return nil
                        6-> return nil
                    5-> return 5 -> nil
                4 -> return 4 -> 5
            3 -> return 3 -> 4 -> 5
        6 -> return 3 -> 4 -> 5
    2 -> return 2 -> 3 -> 4 -> 5
1 -> return 1 -> 2 -> 3 -> 4 -> 5


1.
*/

// recursive the head.Next
// TC: O(n), SC: O(n)
func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	if head.Val == val {
		return removeElements2(head.Next, val)
	}

	head.Next = removeElements2(head.Next, val)
	if head.Val == val {
		return head.Next
	}

	return head
}

// do not need to keep track prev
// current node is not equal to val, we will compare val with next node
// TC: O(n), SC: O(1)
func removeElements3(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	current := head
	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

/*
prev := Head
c := h

	h
	1 2 6 3 4 5 6
	c

	h
	7 7 7
	      c
d
p

c.Val == val
- true: p.Next = c.Next
- false: p = c
c = c.Next

head is posibly modified if head.Val == val
*/

// 1. keep dummy node to keep track the head
// 2. prev to keep track the previous of current
// 3. current node

// dummy = ListNode{Next: head}
// prev = dummy
// current = head
// TC: O(n), SC: O(1)
func removeElements4(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	cur := head
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}

		cur = cur.Next
	}

	return dummy.Next

}
