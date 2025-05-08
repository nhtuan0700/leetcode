from typing import List

# https://leetcode.com/problems/can-make-palindrome-from-substring/

# Intuition:
# 1. How to make s to be palindrome
# - odd frequencies <= 1 or odd <= 1
# 2. We have k replacements, if odd > 1
# - For example: We have 5 odds: a,b,c,d,e.Then we need to have at least 2 replacements
#   => odd - 2*k <= 1
#     a b c d a
# a 0 1 1 1 1 2
# b 0 0 1 1 1 1
# c 0 0 0 1 1 1
# d 0 0 0 0 1 1


# Approach
# Find the odd frequencies in each query from string s
# - We could calculate frequencies of each character in string s
# - After that, We could know the odd frequencies in each queries

class Solution:
    def canMakePaliQueries(self, s: str, queries: List[List[int]]) -> List[bool]:
        n = len(s) + 1
        cnt = [[0] * n for _ in range(26)]
        for i, c in enumerate(s):
            idx = ord(c) - ord('a')
            for j in range(26):
                cnt[j][i+1] = cnt[j][i] + (j == idx)
    
        res = [False] * len(queries)
        for i, q in enumerate(queries):
            l, r, k = q
            # num&1 = num%2
            odd = sum((cnt[c][r+1] - cnt[c][l]) & 1 for c in range(26))
            res[i] = odd - 2*k <= 1
        
        return res
