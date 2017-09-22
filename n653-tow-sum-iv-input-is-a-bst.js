/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} k
 * @return {boolean}
 */
var findTarget = function(root, k) {
    let data = {}
    
    // 闭包，可以直接访问外层函数的变量
    function dfs(root) {
        if (!root) return false;
        
        if (data[k-root.val]) return true;
        
        data[root.val] = 1
        return dfs(root.left) || dfs(root.right)
    }
     
    return dfs(root)
};
