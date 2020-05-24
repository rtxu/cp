/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    left := postorderTraversal(root.Left)
    right := postorderTraversal(root.Right)
    
    result := make([]int, len(left) + len(right) + 1)
    copy(result, left)
    copy(result[len(left):], right)
    result[len(left) + len(right)] = root.Val
    
    return result
}

func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    left := postorderTraversal(root.Left)
    right := postorderTraversal(root.Right)
    
    result := make([]int, len(left) + len(right) + 1)
    result[0] = root.Val
    copy(result[1:], left)
    copy(result[1+len(left):], right)
    
    return result
}

func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    left := postorderTraversal(root.Left)
    right := postorderTraversal(root.Right)
    
    result := make([]int, len(left) + len(right) + 1)
    
    copy(result, left)
    result[len(left)] = root.Val
    copy(result[1+len(left):], right)
    
    return result
}
