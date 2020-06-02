/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, isOdd []bool) int {
    if root == nil {
        return 0
    } else if root.Left == nil && root.Right == nil {
        // leaf
        isOdd[root.Val] = !isOdd[root.Val]
        var oddCnt int
        for i := 1; i <= 9; i++ {
            if isOdd[i] {
                oddCnt++
            }
        }
        isOdd[root.Val] = !isOdd[root.Val]
        if oddCnt > 1 {
            return 0
        } else {
            return 1
        }
    } else {
        //fmt.Println(root.Val, isOdd)
        isOdd[root.Val] = !isOdd[root.Val]
        var result int
        result += dfs(root.Left, isOdd)
        result += dfs(root.Right, isOdd)
        isOdd[root.Val] = !isOdd[root.Val]
        return result
    }
}

func pseudoPalindromicPaths (root *TreeNode) int {
    isOdd := make([]bool, 10)
    return dfs(root, isOdd)
}
