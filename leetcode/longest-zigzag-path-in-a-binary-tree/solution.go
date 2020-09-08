/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a, b int) int {
    if a < b {
        return b
    } else {
        return a
    }
}

func dfs(current *TreeNode, maxLen *int) (int, int) {
    if current == nil {
        return 0, 0
    }
    
    _, l_right := dfs(current.Left, maxLen)
    r_left, _ := dfs(current.Right, maxLen)
    
    left, right := 1+l_right, 1+r_left
    *maxLen = max(*maxLen, left-1)
    *maxLen = max(*maxLen, right-1)
    
    return left, right
}

func longestZigZag(root *TreeNode) int {
    var maxLen int
    dfs(root, &maxLen)
    return maxLen
}
