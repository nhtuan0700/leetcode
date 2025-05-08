from typing import List

# Idea: Prefix+suffix
# Pass 1 (left to right): products[i] = product to all nums to the left of i
# Pass 1 (right to left): products[i] *= product to all nums to the right of i

# TC: O(n), SC: O(n)
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        productLeft = 1
        productRight = 1
        res = [1] * n
        for i in range(0, n):
            res[i] = productLeft
            productLeft *= nums[i]
        for i in range(n - 1, -1, -1):
            res[i] *= productRight
            productRight *= nums[i]

        return res

# not recommend
# TC: O(n), SC: O(n)
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        product = 1
        n = len(nums)
        zero = 0
        for num in nums:
            if num == 0:
                zero += 1
                continue
            product *= num
        ans = [0] * n
        if zero > 1:
            return ans
        for i, num in enumerate(nums):
            if num == 0:
                ans[i] = product
            elif zero > 0:
                ans[i] = 0
            else:
                ans[i] = product // num
        return ans
