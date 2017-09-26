/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {number[]} nums
 * @return {TreeNode}
 */
var constructMaximumBinaryTree = function(nums) {
    if (nums.length === 0) {
        return null
    }
    
    var idx = 0
    for (var i = 0; i < nums.length; i++) {
        if (nums[i] > nums[idx]) {
            idx = i
        }
    }
    
    var ret = new TreeNode(nums[idx])
    ret.left = constructMaximumBinaryTree(nums.slice(0, idx))
    ret.right = constructMaximumBinaryTree(nums.slice(idx+1))
    
    return ret
};
