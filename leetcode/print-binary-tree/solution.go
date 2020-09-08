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

func getDepth(current *TreeNode, depth int, maxDepth *int) {
    if current == nil {
        return
    }
    *maxDepth = max(*maxDepth, depth)
    getDepth(current.Left, depth+1, maxDepth)
    getDepth(current.Right, depth+1, maxDepth)
}

func dfs(current *TreeNode, depth, totalDepth, index int, result [][]string) {
    if current == nil {
        return 
    }
    result[depth-1][index] = fmt.Sprint(current.Val)
    offset := (1 << (totalDepth-depth)) / 2
    dfs(current.Left, depth+1, totalDepth, index-offset, result)
    dfs(current.Right, depth+1, totalDepth, index+offset, result)
}

func printTree(root *TreeNode) [][]string {
    var maxDepth int
    getDepth(root, 1, &maxDepth)
    
    col := (1 << maxDepth) - 1
    result := make([][]string, maxDepth)
    for i := 0; i < maxDepth; i++ {
        result[i] = make([]string, col)
    }
    dfs(root, 1, maxDepth, col/2, result)
    return result
}
