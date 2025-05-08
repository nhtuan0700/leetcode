from collections import defaultdict
from typing import List

# TC: O(n*mlog(m)), SC: O(n*m)
# - n: length of strs
# - m: max length of strs
# - k: number of unique strs
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        freqs = defaultdict(list)
        for (i, s) in enumerate(strs):
            freqs["".join(sorted(s))].append(i)
        res = []
        for freq in freqs.values():
            res.append([strs[f] for f in freq])
        return res
