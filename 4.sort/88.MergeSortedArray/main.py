from typing import List

# https://leetcode.com/problems/merge-sorted-array/

# Vì phải modify nums1 in-place, nên ta duyệt từ cuối mảng về đầu để điền n phần tử của nums2 vào sau trước, tránh bị ghi đè lên m phần tử ban đầu của nums1.

# TC: O(m+n), SC: O(1)
class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        i, j = m - 1, n - 1
        for k in range(m + n - 1, -1, -1):
            if j < 0 or (i >= 0 and nums1[i] >= nums2[j]):
                nums1[k] = nums1[i]
                i -= 1
            else:
                nums1[k] = nums2[j]
                j -= 1
        return nums1
