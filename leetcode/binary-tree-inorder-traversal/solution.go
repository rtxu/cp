/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    Q := make([]*TreeNode, 0)
    result := make([]int, 0)
    
    cur := root
    for {
        if cur == nil {
            if len(Q) == 0 {
                break
            }
            cur = Q[len(Q) - 1]
            cur.Left = nil
            Q = Q[:len(Q)-1]
        }
        if cur.Left != nil {
            Q = append(Q, cur)
            cur = cur.Left
        } else {
            result = append(result, cur.Val)
            cur = cur.Right
        }
    }
    
    return result
}
