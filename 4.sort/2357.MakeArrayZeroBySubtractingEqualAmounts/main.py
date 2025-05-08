from typing import List

# Straight forward
# TC: O(nlogn), SC: O(1)
class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        nums.sort()
        i = 0
        n = len(nums)

        while i < n and nums[i] == 0:
            i += 1

        maxVal = nums[-1]
        res = 0
        total = 0
        while i < n:
            if maxVal - total < 0:
                break
            cur = nums[i]
            while i + 1 < n and nums[i + 1] == cur:
                i += 1
            total += (nums[i] - total)
            res += 1
            i += 1

        return res

# Bởi mỗi số khác nhau > 0 đều phải ít nhất 1 lần trừ riêng => nên kết quả = số unique > 0
# TC: O(n), SC: O(n)
class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        seen = set()
        for num in nums:
            if num > 0:
                seen.add(num)
        return len(seen)
