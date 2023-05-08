import re

class Solution(object):
    def isPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """
        pattern = re.compile(r'[^a-zA-Z0-9]')
        stripped = re.sub(pattern, '', s).lower()
        for i in range(0, int(len(stripped)/2)):
            mirrorIndex = len(stripped) - 1 - i 
            if stripped[i] != stripped[mirrorIndex]:
                return False
        return True


class Tests:
    def __init__(self, s, expected):
        self.s = s
        self.expected = expected


if __name__ == '__main__':
    tests = [
        Tests("A man, a plan, a canal: Panama", True),
        Tests(" ", True),
        Tests("race a car", False)
    ]
    sol = Solution()

for test in tests:
    res = sol.isPalindrome(test.s)
    if res != test.expected:
        print("fail %r; correct - %r" % (res,  test.expected))
    else:
        print("ok %r" % (test.expected))