
# https://www.lintcode.com/problem/169/description

from typing import (
    List,
)

# Idea: move from A to C
# 1. Move n-1 disks from A to B using C
# 2. Move last disk from A to C
# 3. Move n-1 disk from B to C using A

# TC: O(2^n), SC: O(2^n)
class Solution:
    """
    @param n: the number of disks
    @return: the order of moves
    """
    def tower_of_hanoi(self, n: int) -> List[str]:
        # write your code here
        res = []
        self.move(n, 'A', 'C', 'B', res)
        return res

    def move(self, n, fromP, toP, auxP, res):
        if n == 0:
            return res
        self.move(n-1, fromP, auxP, toP, res)
        res.append(self.printMove(fromP, toP))
        self.move(n-1, auxP, toP, fromP, res)


    def printMove(self, fromP, toP):
        return f"from {fromP} to {toP}"
