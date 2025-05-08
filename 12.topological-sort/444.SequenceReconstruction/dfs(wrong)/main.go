package main

// https://www.lintcode.com/problem/605/description

// TC: O(V + E), SC: O(V+E)
// V: len(org)
// E: total elm of seqs

// => it's wrong , because
// [1,2,3] [[1,2],[1,3]] => false -> it's correct
// [1,2,3] [[1,3],[1,2]] => true -> wrong
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
	vis := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if vis[i] == 0 {
			if !dfs(i, adj, vis, &res) {
				return false
			}
		}
	}

	l := 0
	h := n - 1
	for l <= h {
		if res[l] != org[h] {
			return false
		}

		l++
		h--
	}
	return true
}

func dfs(cur int, adj [][]int, vis []int, res *[]int) bool {
	vis[cur] = 1 // visiting

	for _, v := range adj[cur] {
		if vis[v] == 0 {
			if !dfs(v, adj, vis, res) {
				return false
			}
		} else if vis[v] == 1 {
			return false
		}
	}

	vis[cur] = 2 // done

	*res = append(*res, cur)

	return true
}
