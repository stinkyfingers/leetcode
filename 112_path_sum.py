
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

    def buildTree(self, arr):
        head = TreeNode(arr[0])
        arr.pop(0)
        nodes = [head]

        while len(arr) > 0:
            current = nodes.pop(0)

            left = TreeNode(arr[0])
            current.left = left
            nodes.append(left)
            arr.pop(0)
            if len(arr) == 0:
                break
            right = TreeNode(arr[0])
            current.right = right
            nodes.append(right)
            arr.pop(0)
        return head

class Solution(object):
    def hasPathSum(self, root, targetSum):
        """
        :type root: TreeNode
        :type targetSum: int
        :rtype: bool
        """
        if root is None:
            return False
        return self.findSum(root, targetSum)

    def findSum(self, tree, target, total=0):
        if tree.val is None:
            return False
        if tree.left is None and tree.right is None and total+tree.val == target:
            return True

        left = False
        right = False
        if tree.left is not None:
            left = self.findSum(tree.left, target, total + tree.val) 
        if tree.right is not None:
            right = self.findSum(tree.right, target, total + tree.val)
        return left or right


class Tests:
    def __init__(self, root, targetSum, expecteed):
        self.root = root
        self.targetSum = targetSum
        self.expected = expecteed


if __name__ == '__main__':
    tests = [
        Tests([5,4,8,11,None,13,4,7,2,None, None, None,1], 22, True),
        Tests([1,2,3], 5, False),
        Tests([1], 1, True),
        Tests([1,2], 1, False)
    ]
    sol = Solution()

for test in tests:
    root = TreeNode()
    root = root.buildTree(test.root)
    res = sol.hasPathSum(root, test.targetSum)
    if res != test.expected:
        print("fail %r; correct - %r" % (res,  test.expected))
    else:
        print("ok %r" % (test.expected))