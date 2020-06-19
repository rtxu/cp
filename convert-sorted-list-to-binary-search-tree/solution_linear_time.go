
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Time: O(n) logn is the depth of BST
// Space: O(logn) 

var node *ListNode

func inorder(start, end int) *TreeNode {
    if start > end {
        return nil
    }
    mid := (start+end)/2
    left := inorder(start, mid-1)
    tree := &TreeNode{
        Val: node.Val,
        Left: left,
    }
    node = node.Next
    tree.Right = inorder(mid+1, end)
    return tree
}


func sortedListToBST(head *ListNode) *TreeNode {
    count := 0
    for current := head; current != nil; current = current.Next {count++}
    node = head
    return inorder(0, count-1)
}

