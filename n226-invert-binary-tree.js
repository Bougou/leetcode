/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
var invertTree = function(root) {
    if (root !== null) {
        var left = invertTree(root.right)
        var right = invertTree(root.left)
        
        root.left = left
        root.right = right
    }
    
    return root
};
