package main

// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/description/

import (
	"container/list"
)

// sandwich: 0 | 1
// sandwich: stack
// student: queue
// process until student is back
// if student[i] == sandwitch[i]
//  - true: remove student and sandwich
//  - false: remove student and push student to back again

// dry run:
// st: [1,1,0,0], sw: [0,1,0,1]
// st = [1,0,0,1], sw: [0,1,0,1]
// st = [0,0,1,1], sw: [0,1,0,1]
// st = [0,1,1], sw: [1,0,1]
// st = [1,1,0], sw: [1,0,1]
// st = [1,0], sw: [0,1]
// st = [0,1], sw: [0,1]
// st = [1], sw: [1]
// st = [], sw: []

// use wait, if wait == len(queue) => return
// TC: O(n), SC: O(n)
func countStudents1(students []int, sandwiches []int) int {
	topQueue := 0
	wait := 0
	queue := list.New()
	for _, st := range students {
		queue.PushBack(st)
	}

	for queue.Front() != nil {
		if wait == queue.Len() {
			break
		}
		if queue.Front().Value == sandwiches[topQueue] {
			queue.Remove(queue.Front())
			wait = 0
			topQueue++
		} else {
			queue.MoveToBack(queue.Front())
			wait++
		}
	}

	return queue.Len()
}

// not using wait
// TC: O(n), SC: O(n)
func countStudents2(students []int, sandwiches []int) int {
	topQueue := 0
	queue := list.New()
	for _, st := range students {
		queue.PushBack(st)
	}

	for queue.Front() != nil {
		originalSize := queue.Len()
		for i := 0; i < originalSize; i++ {
			if queue.Front().Value == sandwiches[topQueue] {
				queue.Remove(queue.Front())
				topQueue++
				break
			} else {
				queue.MoveToBack(queue.Front())
			}
		}

		if originalSize == queue.Len() {
			return queue.Len()
		}
	}

	return 0
}

// thứ tự sandwiches cố định -> đi lần lần
// thứ tự students không quan trọng -> chỉ cần đếm số lượng sandwich muốn ăn tương ứng của students
// ví dụ: chúng ta đang ở sandwich 0 mà student còn lại chỉ muốn ăn loại 1 [0, count] => return count
//
//	hoặc ngược lại, đang ở sandwich 1 mà student còn lại chỉ muốn ăn loại 0 [count, 0] => return count
// TC: O(n), SC: O(1)
func countStudents3(students []int, sandwiches []int) int {
	wants := [2]int{}
	for _, st := range students {
		wants[st]++
	}

	for _, sw := range sandwiches {
		if wants[sw] > 0 {
			wants[sw]--
		} else {
			break
		}
	}

	return wants[0] + wants[1]
}

// bad because:
// students = append(students, students[0])
// students = students[1:]
// TC: O(n^2), SC: O(1)
func countStudents4(students []int, sandwiches []int) int {
	topQueue := 0
	wait := 0

	for len(students) > 0 {
		if wait == len(students) {
			break
		}
		if students[0] == sandwiches[topQueue] {
			students = students[1:]
			wait = 0
			topQueue++
		} else {
			students = append(students, students[0])
			students = students[1:]
			wait++
		}
	}

	return len(students)
}

// Same as description of problem
// TC: O(n), SC: O(2n)
func countStudents5(students []int, sandwiches []int) int {
	wait := 0
	stack := list.New()
	queue := list.New()
	for _, st := range students {
		queue.PushBack(st)
	}
	for i := len(sandwiches) - 1; i >= 0; i-- {
		stack.PushBack(sandwiches[i])
	}

	for queue.Front() != nil {
		if wait == queue.Len() {
			break
		}
		if queue.Front().Value == stack.Back().Value {
			queue.Remove(queue.Front())
			stack.Remove(stack.Back())
			wait = 0
		} else {
			queue.MoveToBack(queue.Front())
			wait++
		}
	}

	return queue.Len()
}
