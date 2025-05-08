from typing import List

# https://leetcode.com/problems/can-place-flowers/

# [0,0,0,0,0] => 3
# [0,0,0,0,0,0] => 3
# nmax = (n+1)/2

# [1,0,0,0,1]
# count 0's flowerbed item

# TC: O(n), SC: O(1)
class Solution:
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        def maximumPlaceNewFlowers(flowerbed):
            n = len(flowerbed)
            res = 0
            i = 0
            while i < n:
                if flowerbed[i] == 1:
                    i += 2
                    continue
                j = i
                # flowerbed[j] = 0
                while j + 1 < n and (flowerbed[j+ 1] == 0):
                    j += 1

                # length of flowerbed[i:j] = 0
                cnt = j - i + 1
                # flowerbed[i] = 1 or i == n
                if j + 1 < n:
                    cnt -= 1
                # if i > 0:
                #     cnt -= 1
                res += (cnt + 1) // 2
                i = j + 1

            return res

        return maximumPlaceNewFlowers(flowerbed) >= n
