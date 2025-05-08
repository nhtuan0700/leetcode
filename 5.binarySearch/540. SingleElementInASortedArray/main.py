from typing import List

# TC:: O(logn), SC: O(1)
class Solution:
    def singleNonDuplicate(self, nums: List[int]) -> int:
        n = len(nums) - 1
        l, r = 0, n
        while l < r:
            m = l + (r-l)//2
            if nums[m] != nums[m+1] and  nums[m] != nums[m-1]:
                return nums[m]
            if m%2 == 0:
                if nums[m] == nums[m+1]:
                    l = m + 1
                else:
                    r = m - 1
            else:
                if nums[m] == nums[m+1]:
                    r = m - 1
                else:
                    l = m + 1
        
        return nums[l]

class Solution:
    def singleNonDuplicate(self, nums: List[int]) -> int:
        n = len(nums) - 1
        l, r = 0, n
        while l <= r:
            m = l + (r-l)//2
            equalLeft = m > 0 and nums[m] == nums[m-1]
            equalRight = m < n and nums[m] == nums[m+1]
            if not equalLeft and not equalRight:
                return nums[m]
            if equalLeft:
                if m%2 == 0:
                    r = m - 1
                else:
                    l = m + 1
            else:
                if m%2 == 0:
                    l = m + 1
                else:
                    r = m - 1
        
        return -1
