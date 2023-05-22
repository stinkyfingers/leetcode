import re

class Solution(object):
    def singleNumber(self, nums):
        m = {}
        for num in nums:
            if num in m:
                m[num] = m[num] + 1
            else:
                m[num] = 1
        for i in m:
            if m[i] == 1:
                return i

class Tests:
    def __init__(self, s, expected):
        self.s = s
        self.expected = expected


if __name__ == '__main__':
    tests = [
        Tests([2,2,1], 1),
        Tests([4,1,2,1,2], 4),
        Tests([1], 1)
    ]
    sol = Solution()

for test in tests:
    res = sol.singleNumber(test.s)
    if res != test.expected:
        print("fail %r; correct - %r" % (res,  test.expected))
    else:
        print("ok %r" % (test.expected))