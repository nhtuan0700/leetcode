package main

// https://leetcode.com/problems/alien-dictionary/description/
// https://www.lintcode.com/problem/892/description

/**
 * @param words: a list of words
 * @return: a string which is correct order
 */

// ["wrt","wrf","er","ett","rftt"]
// t < f, w < e, r < t, e < r
// t [f]
// w [e]
// r [t]
// e [r]
// => w e r t f

// invalid:
// - cycle
// - aaa aa

// TC: O(W*L) SC: O(1)
// W: len(words)
// L: average length of word in words
// DFS: worst case: 26 nodes, 26*(26-1)/2 edges
func AlienOrder(words []string) string {
	n := len(words)
	adj := make(map[int][]int)

	for _, w := range words {
		for i := 0; i < len(w); i++ {
			c := int(w[i] - 'a')
			if _, ok := adj[c]; !ok {
				adj[c] = make([]int, 0)
			}
		}
	}

	for i := 0; i < n-1; i++ {
		cur, next := words[i], words[i+1]
		for j := 0; j < len(cur); j++ {
			// aaa aa => invalid
			if j >= len(next) {
				return ""
			}
			if cur[j] == next[j] {
				continue
			}
			c1 := int(cur[j] - 'a')
			c2 := int(next[j] - 'a')
			adj[c1] = append(adj[c1], c2)
			break
		}
	}

	vis := make([]bool, 26)
	path := make([]bool, 26)
	res := make([]byte, 0)
	for k := range adj {
		if vis[k] {
			continue
		}

		if !dfs(k, adj, vis, path, &res) {
			return ""
		}
	}

	l, h := 0, len(res)-1
	for l <= h {
		res[l], res[h] = res[h], res[l]
		l++
		h--
	}

	return string(res)
}

func dfs(cur int, adj map[int][]int, vis, path []bool, res *[]byte) bool {
	vis[cur] = true
	path[cur] = true

	for _, v := range adj[cur] {
		if !vis[v] {
			if !dfs(v, adj, vis, path, res) {
				return false
			}
		} else if path[v] {
			return false
		}
	}

	*res = append(*res, byte(cur)+'a')

	path[cur] = false

	return true
}
