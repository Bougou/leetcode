function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

var trimBST = function(root, L, R) {
    if (root === null) {
        return null
    }
    
    if (root.val < L) {
        return trimBST(root.right, L, R)
    } else if (root.val > R) {
        return trimBST(root.left, L, R)
    } else {
        root.left = trimBST(root.left, L, R)
        root.right = trimBST(root.right, L, R)
        return root
    }
}
