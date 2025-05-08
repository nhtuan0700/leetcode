package main

// https://www.lintcode.com/problem/605/description

/**
 * @param org: a permutation of the integers from 1 to n
 * @param seqs: a list of sequences
 * @return: true if it can be reconstructed only one or false
 */

// bfs
// 1. build adjacent list from seqs
// 2. make sure that having only one node with indegree-0

// Leet code make sure the sequences[i] is a subsequences of nums
// But, lint code dont

// TC: O(V + E), SC: O(V+E)
// V: len(org)
// E: total length of all seqs
func SequenceReconstruction(org []int, seqs [][]int) bool {
	n := len(org)
	adj := make([][]int, n+1)
	inDegree := make([]int, n+1)

	seen := make(map[int]bool)
	for _, sq := range seqs {
		if len(sq) == 0 {
			continue
		}

		seen[sq[0]] = true
		for i := 1; i < len(sq); i++ {
			u := sq[i-1]
			v := sq[i]

			// make sure the squence i is not exceed n
			if u > n || v > n {
				return false
			}
			adj[u] = append(adj[u], v)
			inDegree[v]++
			seen[v] = true
		}
	}
	if len(seen) != n {
		return false
	}

	res := make([]int, 0)
	q := make([]int, 0)
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}
	if len(q) > 1 {
		return false
	}

	head := 0
	for head < len(q) {
		cur := q[head]
		head++
		res = append(res, cur)
		count := 0
		for _, v := range adj[cur] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
				count++
			}
		}

		if count > 1 {
			return false
		}
	}

	if len(res) != n {
		return false
	}

	for i := 0; i < n; i++ {
		if res[i] != org[i] {
			return false
		}
	}
	return true
}
