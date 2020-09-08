/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func totalSum(current *TreeNode, total int, result *int64) int {
    if current == nil {
        return 0
    }
    sum := totalSum(current.Left, total, result) + totalSum(current.Right, total, result) + current.Val
    if result != nil {
        v := int64(total-sum)*int64(sum)
        if v > *result {
            *result = v
        }
    }
    return sum
}

func maxProduct(root *TreeNode) int {
    var result int64
    sum := totalSum(root, 0, nil)
    totalSum(root, sum, &result)
    return int(result % (1e9+7))
}
