/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    n := len(inorder)
    if n == 0 {
        return nil
    }
    root := &TreeNode{postorder[n-1], nil, nil}
    var i int
    for i = 0; i < n && inorder[i] != postorder[n-1]; i++ {}
    root.Left = buildTree(inorder[0:i], postorder[0:i])
    root.Right = buildTree(inorder[i+1:], postorder[i:n-1])
    return root
}
