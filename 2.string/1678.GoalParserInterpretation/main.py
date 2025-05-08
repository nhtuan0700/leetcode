# https://leetcode.com/problems/goal-parser-interpretation/

# TC: O(n), SC: O(n)
class Solution:
    def interpret(self, command: str) -> str:
        res = ""
        n = len(command)
        i = 0
        while i < n:
            if command[i] == "G":
                res += "G"
                i += 1
            else: # (
                if command[i+1] == ")":
                    res += "o"
                    i += 2
                else:
                    res += "al"
                    i += 4

        return res
