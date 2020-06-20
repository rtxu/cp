/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Time: O(n*depth) depth=logn
func constructMaximumBinaryTree(nums []int) *TreeNode {
    n := len(nums)
    if n == 0 {
        return nil
    }
    
    maxIndex := 0
    for i := 1; i < n; i++ {
        if nums[i] > nums[maxIndex] {
            maxIndex = i
        }
    }
    left, right := constructMaximumBinaryTree(nums[0:maxIndex]), constructMaximumBinaryTree(nums[maxIndex+1:])
    root := &TreeNode{nums[maxIndex], left, right}
    return root
}

