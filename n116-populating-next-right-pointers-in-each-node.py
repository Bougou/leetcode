# Definition for binary tree with next pointer.
# class TreeLinkNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
#         self.next = None

class Solution:
    # @param root, a tree link node
    # @return nothing
    def connect(self, root):
        self.connectHelper(root, None, True)
    
    def connectHelper(self, root, parent, isleft):
        if root == None:
            return
        
        if parent != None:
            if isleft:
                root.next = parent.right
            else:
                if parent.next != None:
                    root.next = parent.next.left
        
        self.connectHelper(root.left, root, True)
        self.connectHelper(root.right, root, False)

class Solution2:
    # @param root, a tree link node
    # @return nothing
    def connect(self, root):
        # loop for the left node on every level
        while root:
            # loop all nodes on the level 
            cur = root
            while cur and cur.left:
                cur.left.next = cur.right
                if cur.next:
                    cur.right.next = cur.next.left
                cur = cur.next
            
            root = root.left
