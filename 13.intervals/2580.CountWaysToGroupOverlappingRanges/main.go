package main

import (
	"math/big"
	"sort"
)

// https://leetcode.com/problems/count-ways-to-group-overlapping-ranges/

// 1. count ranges that are not overlapping
// 2. find the total number of possible ways to split into 2 groups
// count => ways
// 1 => 2
// 2 => 4
// 3 => 8
// 4 => 16
// ...
// n => 2^n
// modulo: 2^n % 10^9+7
// 1,2,3,4 (4 ranges)
// [1,2,3,4][]
// [][1,2,3,4]
// [1,2][3,4]
// [3,4][1,2]
// [1,3][2,4]
// [2,4][1,3]
// [1,4][2,3]
// [2,3][1,4]
// [1,2,3][4]
// [4][1,2,3]
// [1,2,4][3]
// [3][1,2,4]
// [1,3,4][2]
// [2][1,3,4]
// [2,3,4][1]
// [1][2,3,4]

// TC: O(n), SC: O(1)
func countWays(ranges [][]int) int {
	mod := 1_000_000_007
	// [[1,3],[2,5],[4,8],[10,20]]
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
	count := 0
	last := -1

	for _, r := range ranges {
		// not overlap
		if last < r[0] {
			count++
		}
		last = max(last, r[1])
	}

	// mod := 1_000_000_007 // 1e9+7
	// result := 1
	// for i := 0; i < count; i++ {
	// 	result = (result * 2) % mod
	// }
	// return result

	// return powMod(2, count, mod)

	base := big.NewInt(2)
	exp := big.NewInt(int64(count))
	m := big.NewInt(int64(mod))
	result := new(big.Int).Exp(base, exp, m)

	return int(result.Int64())
}

func powMod(base, exp, mod int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1 // right shift <=> exp /= 2
	}
	return result
}
