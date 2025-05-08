package main

// https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/

// TC: O(V+E), SC: O(V+E)
// V: n (not care about m because m <= n)
// E: len(beforeItems) (not care about edges of group because they are always smaller than edges of items)
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	itemAdj := make([][]int, n)
	itemDegree := make([]int, n)

	for i := range group {
		if group[i] == -1 {
			group[i] = m
			m++
		}
	}

	groupAdj := make([][]int, m)
	groupDegree := make([]int, m)
	for i, items := range beforeItems {
		for _, item := range items {
			u := item
			v := i
			groupU := group[u]
			groupV := group[v]

			itemAdj[u] = append(itemAdj[u], v)
			itemDegree[v]++

			if groupU != groupV {
				groupAdj[groupU] = append(groupAdj[groupU], groupV)
				groupDegree[groupV]++
			}
		}
	}

	groupItems := make([][]int, m)
	q := make([]int, 0)
	for i := 0; i < n; i++ {
		if itemDegree[i] == 0 {
			q = append(q, i)
		}
	}

	head := 0
	count := 0
	for head < len(q) {
		itemCur := q[head]
		head++
		groupCur := group[itemCur]

		count++
		groupItems[groupCur] = append(groupItems[groupCur], itemCur)

		for _, v := range itemAdj[itemCur] {
			itemDegree[v]--
			if itemDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	// detect cycle for items
	if count < n {
		return nil
	}

	q = make([]int, 0)
	for i := 0; i < m; i++ {
		if groupDegree[i] == 0 {
			q = append(q, i)
		}
	}

	res := make([]int, 0)
	head = 0
	count = 0
	for head < len(q) {
		cur := q[head]
		head++

		res = append(res, groupItems[cur]...)
		count++
		for _, v := range groupAdj[cur] {
			groupDegree[v]--
			if groupDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	// detect cycle for groups
	if count < m {
		return nil
	}

	return res
}
