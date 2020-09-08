/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Time: O(n)
func constructMaximumBinaryTree(nums []int) *TreeNode {
    n := len(nums)
    if n == 0 {
        return nil
    }
    
    stack := make([]*TreeNode, 0)
    for i := 0; i < n; i++ {
        if len(stack) == 0 {
            stack = append(stack, &TreeNode{nums[i], nil, nil})
        } else {
            var lastPop *TreeNode
            node := &TreeNode{nums[i], nil, nil}
            for len(stack) > 0 && stack[len(stack)-1].Val <= nums[i] {
                lastPop = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            }
            if len(stack) == 0 {
                node.Left = lastPop
            } else {
                stack[len(stack)-1].Right = node
                node.Left = lastPop
            }
            stack = append(stack, node)
        }
    }
    return stack[0]
}

