# https://leetcode.com/problems/remove-duplicate-letters/

# inituation:
# - use monotonic stack to build the result
# - When encoutering the new character:
# 	+ if it is not already in result, compare it with the top of the stack
# 		+ if the top character is greater and it appears again later:
# 			=> pop it from the stack
# - use set to ensure no duplicate characters in the result

class Solution:
    def removeDuplicateLetters(self, s: str) -> str:
        stack = []
        last_occurences = {letter: i for i, letter in enumerate(s)}
        seen = set()
        for i, letter in enumerate(s):
            if letter not in seen:
                while stack and stack[-1] > letter and last_occurences[stack[-1]] > i:
                    seen.remove(stack[-1])
                    stack.pop()
                stack.append(letter)
                seen.add(letter)

        return ''.join(stack)
