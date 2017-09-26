# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def constructMaximumBinaryTree(self, nums):
        """
        :type nums: List[int]
        :rtype: TreeNode
        """
        max_num = nums[0]
        max_index = 0
        for i in xrange(0, len(nums)):
            if nums[i] > max_num:
                max_num = nums[i]
                max_index = i
        
        tree = TreeNode(max_num)
        if max_index < len(nums) - 1:
            tree.right = self.constructMaximumBinaryTree(nums[max_index+1:])

        if max_index > 0:
            tree.left = self.constructMaximumBinaryTree(nums[0:max_index])
        
        return tree
