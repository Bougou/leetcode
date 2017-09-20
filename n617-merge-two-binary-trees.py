class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def mergeTrees(self, t1, t2):
        if t1 == None and t2 == None:
            return None
        
        tree = TreeNode((t1.val if t1 else 0) + (t2.val if t2 else 0))
        tree.left = self.mergeTrees(t1 and t1.left, t2 and t2.left)
        tree.right = self.mergeTrees(t1 and t1.right, t2 and t2.right)

        return tree
