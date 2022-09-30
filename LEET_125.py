import re
class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = re.sub(r'[^a-z|0-9]| ', "", s.lower())
        if len(s)%2 != 0: rev = list(s[len(s)//2+1:])
        else: rev = list(s[len(s)//2:])
        rev.reverse()
        rev = "".join(rev)
        return s[:len(s)//2] == rev
