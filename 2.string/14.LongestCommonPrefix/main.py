from typing import List

# https://leetcode.com/problems/longest-common-prefix/

# TC: O(n*m), SC: O(1)
# m: min length of str in strs
class Solution:
	def longestCommonPrefix(self, strs: List[str]) -> str:
		if not strs:
			return ""
		word = strs[0]
		for i in range(0, len(word)):
			for j in range(1, len(strs)):
				if i == len(strs[j]) or strs[j][i] != word[i]:
					return word[:i]
		return word

# TC: O(nlogn), SC: O(1)
class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return strs[0]
        strs.sort()
        firstWord = strs[0]
        lastWord = strs[-1]
        i = 0
        while i < len(firstWord) and i < len(lastWord):
            if firstWord[i] != lastWord[i]:
                return firstWord[:i]
            i += 1
        
        return firstWord[:i]
