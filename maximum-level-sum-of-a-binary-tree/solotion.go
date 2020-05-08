/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    Q := make([]*TreeNode, 1)
    Q[0] = root
    
    level, maxLevel := 1, 1
    maxSum := root.Val
    for len(Q) > 0 {
        nextQ := make([]*TreeNode, 0, 2*len(Q))
        nextSum := 0
        nextLevel := level + 1
        for i := 0; i < len(Q); i++ {
            if Q[i].Left != nil {
                nextQ = append(nextQ, Q[i].Left)
                nextSum += Q[i].Left.Val
            }
            if Q[i].Right != nil {
                nextQ = append(nextQ, Q[i].Right)
                nextSum += Q[i].Right.Val
            }
        }
        if nextSum > maxSum {
            maxSum = nextSum
            maxLevel = nextLevel
        }
        
        level++
        Q = nextQ
    }
    
    return maxLevel
}
