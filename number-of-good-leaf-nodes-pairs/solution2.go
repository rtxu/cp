/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countPairs(root *TreeNode, distance int) int {
    var result int
    postOrder(root, distance, &result)
    return result
}

func postOrder(root *TreeNode, dist int, count *int) []int {
    d := make([]int, dist + 1)
    if root == nil {
        return d
    }
    
    leftChilds := postOrder(root.Left, dist, count)
    rightChilds := postOrder(root.Right, dist, count)
    
    for i := 1; i < dist; i++ {
        for j := 1; i + j <= dist; j++ {
            *count += leftChilds[i] * rightChilds[j]
        }
    }
    
    if root.Left == nil && root.Right == nil {
        d[1] = 1
    } else {
        for i := 1; i < dist; i++ {
            d[i+1] += leftChilds[i] + rightChilds[i]
        }
    }
    return d
}
