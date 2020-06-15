/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func backtrack(current *TreeNode, path []int, pathMax int) int {
    if current == nil {
        return 0
    }
    path = append(path, current.Val)
    count := 0
    if current.Val >= pathMax {
        count++
        pathMax = current.Val
    }
    count += backtrack(current.Left, path, pathMax)
    count += backtrack(current.Right, path, pathMax)
    return count
}

func goodNodes(root *TreeNode) int {
    return backtrack(root, []int{}, -99999)
}
