from typing import List

# SC: O(n)
class NumArray:
	# TC: O(n)
    def __init__(self, nums: List[int]):
        n = len(nums)
        prefix = [0] * (n + 1)
        for i in range(n):
            prefix[i+1] = prefix[i] + nums[i]
        self.prefix = prefix
	# TC: O(1)
    def sumRange(self, left: int, right: int) -> int:
        return self.prefix[right + 1] - self.prefix[left]
