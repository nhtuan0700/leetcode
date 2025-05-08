package main

import "container/list"

// https://leetcode.com/problems/time-needed-to-buy-tickets/

// TC: O(n*tikets[k]) SC: 0
func timeRequiredToBuy1(tickets []int, k int) int {
	current := 0
	count := 0
	n := len(tickets) - 1

	for tickets[k] > 0 {
		if tickets[current] > 0 {
			tickets[current]--
			count++
		}
		if current == n {
			current = 0
		} else {
			current++
		}
	}

	return count
}

// TC: O(n*tickets[k]) SC: O(n)
// worst case: [3,3,(3)] = O(n^2) when ticket[k] >= n
type Ticket struct {
	Pos int
	Val int
}

func timeRequiredToBuy2(tickets []int, k int) int {
	queue := list.New()
	for i, v := range tickets {
		queue.PushBack(&Ticket{
			Pos: i,
			Val: v,
		})
	}

	count := 0
	for queue.Front() != nil && queue.Len() > 1 {
		currTicket := queue.Front().Value.(*Ticket)
		if currTicket.Pos == k && currTicket.Val == 1 {
			return count + 1
		}
		if currTicket.Val > 0 {
			count++
			currTicket.Val--
		} else {
			queue.Remove(queue.Front())
			continue
		}
		queue.MoveToBack(queue.Front())
	}

	return queue.Front().Value.(*Ticket).Val + count
}

/*
for each i, t: tickets
- if i < k:
    - total += min(tickets[i], tickets[k])
    - They are in front of person k
    - so they will always get to buy a ticket in every round until person k finishes.
    => Add min(tickets[i], tickets[k]) to the total time.
- if i == k:
    - This is the person we're tracking.
    → Add tickets[k] to the total time.
- if i > k:
    - They are behind person k
    - so in the last round, person k will buy their final ticket before them.
    => Add min(tickets[i], tickets[k] - 1) to the total time.
*/

// TC: O(n), SC: O(1)
func timeRequiredToBuy(tickets []int, k int) int {
	count := 0
	for i, t := range tickets {
		if i <= k {
			count += min(t, tickets[k])
		} else {
			count += min(t, tickets[k]-1)
		}
	}

	return count
}
