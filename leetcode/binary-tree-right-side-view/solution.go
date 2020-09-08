/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(depth int, root *TreeNode, result *[]int) {
    if root == nil {
        return
    }
    if len(*result) < depth {
        // reach depth at the first time 
        *result = append(*result, root.Val)
    }
    dfs(depth+1, root.Right, result)
    dfs(depth+1, root.Left, result)
}

func rightSideView(root *TreeNode) []int {
    result := make([]int, 0)
    dfs(1, root, &result)
    return result
}
