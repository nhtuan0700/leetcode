package main

// https://leetcode.com/problems/optimal-partition-of-string/

// hash table
// s = "abacaba"
// loop c: s
//  map[c] -> result++,
//  else clear map

// TC: O(n), SC: O(n)
/*
S = {a}
count = 1
   i
ab acaba

for each character, c is in not S:
1. append c into existing substring => optimal choice
2. create new substring, res++
*/
func partitionString1(s string) int {
	freqs := map[rune]bool{}

	result := 0
	for _, r := range s {
		if freqs[r] {
			freqs = map[rune]bool{}
			result++
		}
		freqs[r] = true

	}

	return result + 1
}

func partitionString2(s string) int {
    n := 'z' - 'a' + 1
    freqs := make([]int, n)

    result := 0
    for _, r := range s {
        freqs[r - 'a']++
        if freqs[r - 'a'] == 2 {
            freqs = make([]int, n)
            freqs[r - 'a']++
            result++
        }
    }

    return result + 1
}

// using bit operator
// AND bit: abc: 111, c:100 => 111 & 100 = 100 != 0 => c exist in abc
// OR bit: ab: 11 c: 100 => 11 & 100 = 111
// TC: O(n), SC: O(1)

/*
      i
ab ac aba

bit = 1
mask = 1
count = 2
*/
func partitionString3(s string) int {
	mask, count := 0, 0

	for _, c := range s {
		bit := 1 << (c - 'a') // convert c into bit
		if mask&bit != 0 {    // check if c is existed
			count++
			mask = bit // start new c
		} else {
			mask |= bit // store c
		}
	}

	return count + 1
}
