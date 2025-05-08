from typing import List

# https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/

# Ý tưởng là gom toàn bộ bóng về từng vị trí i bằng cách chia hai bên: trái và phải. 
# Tận dụng prefix sum để tính hiệu quả hơn thay vì duyệt toàn mảng mỗi lần.

# total(box=1), cnt(box=1)
# before_i: cnt_before * i - total_before
# after_i: total_after - cnt_after * i
#      = total - total_before - (cnt - cnt_before) * i
# totalMoves = before_i + after_i

# dry run: 001011
# cnt = 3, total = 11
# cntBefore = 0, totalBefore = 0
# i = 0: cntBefore=0, totalBefore=0
#   totalMoves = (2-0) + (4-0) + (5-0) = (2 + 4 + 5) - 3*0 = 0*0 - 0 + (11-0) - (3-0)*0 = 11
# i = 1: cntBefore=0, totalBefore=0
#   totalMoves = (2-1) + (4-1) + (5-1) = (2 + 4 + 5) - 3*1 = 0*1 - 0 + (11-0) - (3-0)*1 = 8
# i = 2: cntBefore=1, totalBefore=2
#   totalMoves = (2-2) + (4-2) + (5-2) = 1*2 - (2) + (4 + 5) - 2*2 = 1*2 - 2 + (11-2) - (3-1)*2 = 5
# i = 3: cntBefore=1, totalBefore=2 
#   totalMoves = (3-2) + (4-3) + (5-3) = 1*3 - (2) + (4 + 5) - 2*3 = 1*3 - 2 + (11-2) - (3-1)*3 = 4
# i = 4: cntBefore=2, totalBefore=6 
#   totalMoves = (4-2) + (4-4) + (5-4) = 2*4 - (2+4) + (5) - 1*4 = 2*4 - 6 + (11-6) - (3-2)*4 = 3
# i = 5: cntBefore=3, totalBefore=11
#   totalMoves = (5-2) + (5-4) + (5-5) = 3*5 - (2+4+5) = 3*5 - 11 + (11-11) + (3-3)*5 = 4
class Solution:
    def minOperations(self, boxes: str) -> List[int]:
        total = 0
        cnt = 0
        for i, box in enumerate(boxes):
            if box == '1':
                cnt += 1
                total += i
        
        totalBefore = 0
        cntBefore = 0
        res = [0] * len(boxes)
        for i, box in enumerate(boxes):
            if box == '1':
                cntBefore += 1
                totalBefore += i
            
            totalMoves = cntBefore * i - totalBefore + (total - totalBefore) - (cnt - cntBefore) * i
            res[i] = totalMoves
        
        return res


# brute force:
# - i: 0 -> n - 1
#   - In the leftside of i (0 -> i - 1):
#     If boxes[left] == 1, count += i - left
#   - In the rightside of i (i+1, n - 1):
#     If boxes[right] == 1, count += right - i
class Solution:
    def minOperations(self, boxes: str) -> List[int]:
        n = len(boxes)
        res = [0] * n
        for i in range(0, n):
            cnt = 0
            # left
            left = i - 1
            while left >= 0:
                if boxes[left] == '1':
                    cnt += (i-left)
                left -= 1

            right = i+1
            # right
            while right < n:
                if boxes[right] == '1':
                    cnt += (right - i)
                right += 1
            res[i] = cnt
        return res
