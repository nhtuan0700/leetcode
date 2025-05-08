from typing import List

# https://leetcode.com/problems/minimum-cost-of-buying-candies-with-discount/

# greedy
# Bởi vì phải mua tất cả candies, ko thể né giá cao, nên ko thể tối ưu mua rẻ trước
# Điều kiện: "kẹo free phải rẻ hơn hoặc bằng cái rẻ nhất trong 2 cái mua" -> nếu muốn giảm tổng tiền nhiều nhất ta sẽ mua 2 cái đắt nhất để cái free tiếp theo cũng là đắt kế tiếp

# Approach:
# - sort decrease, add 2 first elements, ignore element 3rd
# TC: O(nlogn), SC: O(1)
class Solution:
    def minimumCost(self, cost: List[int]) -> int:
        cost.sort(reverse=True)
        res = 0
        for i in range(len(cost)):
            if (i + 1) % 3 == 0:
                continue
            res += cost[i]
        return res
