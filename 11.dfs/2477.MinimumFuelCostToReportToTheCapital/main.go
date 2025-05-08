package main

// https://leetcode.com/problems/minimum-fuel-cost-to-report-to-the-capital/description/

// at each city, we need to count the number of car that need to transport the people to the city that is closer to the capital city
// numberOfCar = ceil(people / seat)

// TC: O(n), SC: O(n)
func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads) + 1
	adj := make([][]int, n)

	for _, r := range roads {
		u, v := r[0], r[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	fuels := 0

	dfs(adj, 0, -1, seats, &fuels)
	return int64(fuels)
}

func dfs(adj [][]int, cur, prev int, seats int, fuels *int) int {
	people := 1
	for _, v := range adj[cur] {
		if v == prev {
			continue
		}

		people += dfs(adj, v, cur, seats, fuels)
	}

	if cur > 0 {
		// similar: ceil(people / seats)
		*fuels += (people + seats - 1) / seats
	}

	return people
}
