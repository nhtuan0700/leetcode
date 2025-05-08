package main

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

// https://leetcode.com/problems/construct-quad-tree/
// TC: O(n^2) SC: O(logn)

// n = 8 -> 64 cells
// 0 -> 64 (n^2/1)
// 1 -> 16 (n^2/4)
// 2 -> 4  (n^2/16)
// ...
// k -> (n^2 / 4^k)
// T(n) = 4*T(n/2) + 1 = 4*(4*T(n/4) + 1) + 1 = 16*T(n/4) + 5 = 4^k * T(n/2^k) + (4^k - 1) / 3
// Stop T(1) = n/2^k = 1 <=> k = log2(n)
// => 4^log2(n) * T(1) + (4^log2(n) - 1) / 3 (4^log2(n) = n^log2(4) = n^2)
// => n^2 * O(1) + (n^2 - 1) / 3 ~ O(n^2)
func construct(grid [][]int) *Node {
	n := len(grid)
	return buildNode(0, 0, n, grid)
}

func buildNode(i, j, w int, grid [][]int) *Node {
	baseVal := grid[i][j] == 1
	if w == 1 {
		return &Node{Val: baseVal, IsLeaf: true} // leaf
	}

	topLeft := buildNode(i, j, w/2, grid)
	topRight := buildNode(i, j+w/2, w/2, grid)
	botLeft := buildNode(i+w/2, j, w/2, grid)
	botRight := buildNode(i+w/2, j+w/2, w/2, grid)

	// leaf
	if topLeft.IsLeaf && topRight.IsLeaf && botLeft.IsLeaf && botRight.IsLeaf &&
		topLeft.Val == baseVal && topRight.Val == baseVal && botLeft.Val == baseVal && botRight.Val == baseVal {
		return &Node{Val: baseVal, IsLeaf: true}
	} else {
		return &Node{Val: baseVal, IsLeaf: false, TopLeft: topLeft, TopRight: topRight, BottomLeft: botLeft, BottomRight: botRight}
	}
}
