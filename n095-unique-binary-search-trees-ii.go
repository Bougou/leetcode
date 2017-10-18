/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func generateTrees(n int) []*TreeNode {
    return generateTreesHelper(1, n)
}

// [] => low=0, high=0
// [1] => low=1, high=1
// [1,2,3,4] => low=1, high=4
// [3,4,5,6,7] => low=3, high=7
func generateTreesHelper(low int, high int) []*TreeNode {
    res := []*TreeNode{}

	if low == 0 || high == 0 || low > high {
		return res
	}

	if low == high {
		node := &TreeNode{Val:   low}
		res = append(res, node)
		return res
	}

	for i := low; i <= high; i++ {
		leftTrees := generateTreesHelper(low, i-1)
		rightTrees := generateTreesHelper(i+1, high)

        if len(leftTrees) == 0 && len(rightTrees) == 0 {
            node := &TreeNode{Val: i}
            res = append(res, node)
        }
        
        if len(leftTrees) == 0 && len(rightTrees) != 0 {
            for _, r := range rightTrees {
                node := &TreeNode{Val: i, Right: r}
                res = append(res, node)
            }
        }
            
        if len(leftTrees) != 0 && len(rightTrees) == 0 {
            for _, l := range leftTrees {
                node := &TreeNode{Val: i, Left: l}
                res = append(res, node)
            }
        } 
        
        if len(leftTrees) != 0 && len(rightTrees) != 0 {
            for _, l := range leftTrees {
			    for _, r := range rightTrees {
				    node := &TreeNode{
					    Val:   i,
					    Left:  l,
					    Right: r,
				    }
				    res = append(res, node)
			    }
            }
		}
	}
	return res
}
