/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(current *TreeNode, depth int) int {
    if current == nil {
        return depth
    }
    left := dfs(current.Left, depth+1)
    right := dfs(current.Right, depth+1)
    if left < right {
        return right
    } else {
        return left
    }
}

func maxDepth(root *TreeNode) int {
    return dfs(root, 0)
}
