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
    postOrder(root, 0, distance, &result)
    return result
}

func postOrder(root *TreeNode, depth, dist int, count *int) []int {
    if root == nil {
        return nil
    }
    
    leftChilds := postOrder(root.Left, depth+1, dist, count)
    rightChilds := postOrder(root.Right, depth+1, dist, count)
    l, r := len(leftChilds), len(rightChilds)
    
    for i := 0; i < l; i++ {
        for j := 0; j < r; j++ {
            d := leftChilds[i] + rightChilds[j] - 2 * depth
            if d <= dist {
                *count += 1
            }
        }
    }
    
    // fmt.Println("current", root.Val, "leftChilds", leftChilds, "rightChilds", rightChilds)
    if l + r == 0 {
        // leaf
        return []int{depth}
    } else {
        depths := make([]int, len(leftChilds) + len(rightChilds))
        copy(depths, leftChilds)
        copy(depths[len(leftChilds):], rightChilds)
        return depths
    }
}
