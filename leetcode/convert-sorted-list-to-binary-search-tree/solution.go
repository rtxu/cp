
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

// Time: O(nlogn) logn is the depth of BST
// Space: O(logn) 

func toBST(head, tail *ListNode) *TreeNode {
    if head == tail {
        return nil
    }
    slow, fast := head, head
    for fast != tail && fast.Next != tail {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return &TreeNode{
        Val: slow.Val,
        Left: toBST(head, slow),
        Right: toBST(slow.Next, tail),
    }
}

func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    return toBST(head, nil)
}

