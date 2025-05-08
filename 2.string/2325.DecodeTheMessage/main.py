import string

# https://leetcode.com/problems/decode-the-message/

# TC: O(n), SC: O(1)
class Solution:
    def decodeMessage(self, key: str, message: str) -> str:
        alphabet = string.ascii_lowercase
        mapping = {}
        for c in key:
            if c == ' ':
                continue
            if c not in mapping:
                mapping[c] = alphabet[len(mapping)]
            if len(mapping) == len(alphabet):
                break
        
        res = [''] * len(message)
        for i, c in enumerate(message):
            if c == ' ':
                res[i] = ' '
            else:
                res[i] = mapping[c]
        
        return ''.join(res)
            
