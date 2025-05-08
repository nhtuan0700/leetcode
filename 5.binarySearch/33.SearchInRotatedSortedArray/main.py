from typing import List

# https://leetcode.com/problems/search-in-rotated-sorted-array/

# TC: O(n) SC: O(n)
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        startIndex = 0
        n = len(nums)
        for i in range(0, n-1):
            if nums[i] > nums[i+1]:
                startIndex = i + 1
                break
        nums = nums[startIndex:] + nums[:startIndex]
        l, r = 0, n - 1
        while l <= r:
            m = l + (r-l) // 2
            if nums[m] == target:
                res = m + startIndex
                return res if res < n else res - n
            if nums[m] > target:
                r = m - 1
            else:
                l = m + 1
        return -1

# nums[m] > target => thông thường ta sẽ xét "trái" nhưng đây là rotate array ta cần xét tiếp bên dưới
# - if nums[h] < nums[m] (rotate) & nums[h] >= target => nghĩa là phần tử ta cần tìm nằm bên "phải"
# - else: xét bên trái như bình thường
# nums[m] < target
# - if nums[l] > nums[m] (rotate) & nums[l] <= target => xét bên trái
# - else: xét bên phải như bình thường
# TC: O(logn), O(n)
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        n = len(nums)
        l, h = 0, n - 1
        while l <= h:
            m = l + (h - l)//2
            if nums[m] == target:
                return m
            if nums[m] > target: # left
                if nums[h] < nums[m] and nums[h] >= target:
                    l = m + 1
                else:
                    h = m - 1
            else: # right
                if nums[l] > nums[m] and nums[l] <= target:
                    h = m - 1
                else:
                    l = m + 1

        return -1

# Nếu nums[m] == target → return m.
# Xác định nửa nào sorted
# Nếu nums[l] <= nums[m] → left sorted
# - Nếu nums[l] <= target < nums[m] → target ở trái → h = m-1.
# - Ngược lại → target ở phải → l = m+1.
# Nếu nums[m] <= nums[h] → right sorted
# - Nếu nums[m] < target <= nums[h] → target ở phải → l = m+1.
# - Ngược lại → target ở trái → h = m-1.
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        n = len(nums)
        l, h = 0, n - 1
        while l <= h:
            m = l + (h - l)//2
            if nums[m] == target:
                return m
            if nums[l] <= nums[m]: # left sorted
                if nums[l] <= target and target < nums[m]:
                    h = m - 1
                else:
                    l = m + 1
            else: # right sorted
                if nums[h] >= target and target > nums[m]:
                    l = m + 1
                else:
                    h = m - 1

        return -1
