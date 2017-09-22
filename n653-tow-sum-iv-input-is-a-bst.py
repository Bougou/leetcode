# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def findTarget(self, root, k):
        """
        :type root: TreeNode
        :type k: int
        :rtype: bool
        """
        data = {}
        return self.dfs(root, k, data)
    
    def dfs(self, root, k, data):
        if not root:
            return False
        
        if k - root.val in data:
            return True
        
        data[root.val] = 1
        
        return self.dfs(root.left, k, data) or self.dfs(root.right, k, data)
        