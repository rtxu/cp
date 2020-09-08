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

func _isSubPath(head *ListNode, root *TreeNode, inSearch bool) bool {
    if head == nil {
        return true
    }
    if root == nil {
        return false
    }
    
    if inSearch {
        if root.Val == head.Val {
            return _isSubPath(head.Next, root.Left, true) || _isSubPath(head.Next, root.Right, true)
        } else {
            return false
        }
    } else {
        if root.Val == head.Val {
            return _isSubPath(head.Next, root.Left, true) || 
                _isSubPath(head.Next, root.Right, true) ||
                _isSubPath(head, root.Left, false) ||
                _isSubPath(head, root.Right, false)
        } else {
            return _isSubPath(head, root.Left, false) || _isSubPath(head, root.Right, false)
        }
    }
}

func isSubPath(head *ListNode, root *TreeNode) bool {
    return _isSubPath(head, root, false)
}
