function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

var mergeTrees = function(t1, t2) {
    if (!t1 && !t2) {
        return null
    }

    const tree = new TreeNode(t1 ? t1.val : 0 + t2 ? t2.val : 0)

    tree.left = mergeTrees(t1 && t1.left, t2 && t2.left);
    tree.right = mergeTrees(t1 && t1.right, t2 && t2.right);
    return tree
}
