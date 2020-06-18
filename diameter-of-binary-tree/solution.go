/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorder(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }
    ld, lpath := postorder(root.Left)
    rd, rpath := postorder(root.Right)
    diameter := int(math.Max(float64(lpath+rpath), math.Max(float64(ld), float64(rd))))
    maxLenPathFromRoot := int(math.Max(float64(lpath+1), float64(rpath+1)))
    return diameter, maxLenPathFromRoot
}

func diameterOfBinaryTree(root *TreeNode) int {
    diameter, _ := postorder(root)
    return diameter
}
