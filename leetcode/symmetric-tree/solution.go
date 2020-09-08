/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isEqual(l, r *TreeNode) bool {
    if l == nil && r == nil {
        return true
    }
    if l == nil || r == nil {
        return false
    }
    return l.Val == r.Val
}

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    Q := make([]*TreeNode, 0)
    Q = append(Q, root)
    for len(Q) > 0 {
        nextQ := make([]*TreeNode, 0)
        for i := 0; i < len(Q); i++ {
            if Q[i] != nil {
                nextQ = append(nextQ, Q[i].Left, Q[i].Right)
            }
        }
        for i := 0; i < len(nextQ) / 2; i++ {
            if !isEqual(nextQ[i], nextQ[len(nextQ)-1-i]) {
                return false
            }
        }
        Q = nextQ
    }
    
    return true
}
