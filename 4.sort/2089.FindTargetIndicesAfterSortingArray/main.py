from typing import List

# TC: O(nlogn), SC: O(n)
class Solution:
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        nums.sort()
        res = []
        i = 0
        for (i, num) in enumerate(nums):
            if num > target:
                break
            if num == target:
                res.append(i)
        return res

# binary search
# TC: O(nlogn), Sc: O(n)
class Solution:
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        nums.sort()

        def findLeft():
            nonlocal nums
            l = 0
            h = len(nums) - 1
            pos = -1
            while l <= h:
                m = l + (h - l) // 2
                if nums[m] >= target:
                    h = m - 1
                else:
                    l = m + 1
                
                if nums[m] == target:
                    pos = m
            return pos
        def findRight():
            nonlocal nums
            l = 0
            h = len(nums) - 1
            pos = -1
            while l <= h:
                m = l + (h - l) // 2
                if nums[m] <= target:
                    l = m + 1
                else:
                    h = m - 1
                if nums[m] == target:
                    pos = m
            return pos

        left = findLeft()
        right = findRight()

        if left == -1:
            return []
        return list(range(left, right+1))
            