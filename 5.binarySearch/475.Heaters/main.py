from typing import List
import bisect

# https://leetcode.com/problems/heaters

# TC: O(mlogm + nlogn + mlogn), SC: O(1)
# - m: houses length
# - n: heaters length

class Solution:
    def findRadius(self, houses: List[int], heaters: List[int]) -> int:
        houses.sort()
        heaters.sort()

        def getMinRadius(house):
            nonlocal heaters
            l, h = 0, len(heaters) - 1
            res = 0
            pos = len(heaters)-1
            while l <= h:
                m = l + (h-l)//2
                if heaters[m] < house:
                    l = m + 1
                else:
                    h = m - 1
                    pos = m
            left = abs(heaters[pos-1] - house) if pos > 0 else float('inf')
            right = abs(heaters[pos] - house) if pos < len(heaters) else float('inf')
            return min(left, right)
    
        res = 0
        for h in houses:
            res = max(res, getMinRadius(h))

        return res

class Solution:
    def findRadius(self, houses: List[int], heaters: List[int]) -> int:
        houses.sort()
        heaters.sort()

        def getMinRadius(house):
            nonlocal heaters
            l, h = 0, len(heaters) - 1
            minRadius = abs(heaters[0]-house)
            while l <= h:
                m = l + (h-l)//2
                radius = heaters[m] - house
                if radius < 0:
                    l = m + 1
                else:
                    h = m - 1
                minRadius = min(abs(radius), minRadius)
            return minRadius

        res = 0
        for house in houses:
            radius = getMinRadius(house)
            res = max(res, radius)

        return res

class Solution:
    def findRadius(self, houses: List[int], heaters: List[int]) -> int:
        houses.sort()
        heaters.sort()
    
        res = 0
        for house in houses:
            pos = bisect.bisect(heaters, house)
            left = abs(heaters[pos-1] - house) if pos > 0 else float('inf')
            right = abs(heaters[pos] - house) if pos < len(heaters) else float('inf')
            res = max(res, min(left, right))

        return res
