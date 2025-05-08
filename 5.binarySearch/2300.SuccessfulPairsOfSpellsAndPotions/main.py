from typing import List
from math import ceil

# TC: O(mlogn), SC: O(m)
# m: len(spells)
# n: len(potions)
class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        potions.sort()
        def getCountOfPairProductLargerThanSucess(spell):
            nonlocal success, potions
            n = len(potions)
            l, r = 0, n - 1
            res = n
            required = ceil(success/spell)
            while l <= r:
                m = l + (r - l)//2
                # product = potions[m] * spell
                # if product >= sucess:
                if potions[m] >= required:
                    res = m
                    r = m - 1
                else:
                    l = m + 1
            return n - res
        
        res = [0] * len(spells)
        for i, spell in enumerate(spells):
            res[i] = getCountOfPairProductLargerThanSucess(spell)
        return res
