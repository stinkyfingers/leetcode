

class Solution(object):
    def climbStairs(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n < 2:
            return n
        fibs = [None] * (n + 1)
        fibs[0] = 0
        fibs[1] = 1
        i = 2

        while i <= n:
            fibs[i] = fibs[i-1] + fibs[i-2]

            i += 1
        return fibs[n] + fibs[n-1]



if __name__ == '__main__':
    sol = Solution()
    tests = {
        3: 3,
        2: 2,
        1: 1,
        4: 5,
        5: 8,
        6: 13,
        7: 21,
        8: 34,
        9: 55,

    }
    for key, value in tests.items():
        res = sol.climbStairs(key)
        if res != value:
            print("fail %d %d; correct - %d" % (key, res, value))
        else:
            print("ok %d %d" % (key, value))