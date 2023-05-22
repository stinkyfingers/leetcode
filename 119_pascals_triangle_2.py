import re

class Solution(object):
    def getRow(self, rowIndex):
        current = []
        i = 0 
        while i <= rowIndex:
            last = list(current)
            current = []
            length = i + 1
            j = 0
            while j < length:
                if j == 0 or j == length - 1:
                    current.append(1)
                else:
                    current.append(last[j]+last[j-1])
                j += 1
            i += 1
        return current

class Tests:
    def __init__(self, s, expected):
        self.s = s
        self.expected = expected


if __name__ == '__main__':
    tests = [
        Tests(3, [1, 3, 3, 1]),
        Tests(0, [1]),
        Tests(1, [1, 1])
    ]
    sol = Solution()

for test in tests:
    res = sol.getRow(test.s)
    if res != test.expected:
        print("fail %r; correct - %r" % (res,  test.expected))
    else:
        print("ok %r" % (test.expected))