/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} t
 * @return {string}
 */
var tree2str = function(t) {
    if (t === null) {
        return ''
    }
    
    let left
    if (t.left !== null || t.right !== null) {
        left = '(' + tree2str(t.left) + ')'
    } else {
        left = ''
    }
    
    let right
    if (t.right !== null) {
        right = '(' + tree2str(t.right) + ')'
    } else {
        right = ''
    }
    
    return t.val + left + right
    
};
