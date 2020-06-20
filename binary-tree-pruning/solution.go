/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorder(current *TreeNode) int {
    if current == nil {
        return 0
    }
    left, right := postorder(current.Left), postorder(current.Right)
    if left == 0 {
        current.Left = nil
    }
    if right == 0 {
        current.Right = nil
    }
    return left+right+current.Val
}

func pruneTree(root *TreeNode) *TreeNode {
    postorder(root)
    return root
}
