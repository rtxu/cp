/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    var prev *ListNode
    for current := head; current != nil; current = current.Next {
        if prev != nil && prev.Val == current.Val {
            prev.Next = current.Next
        } else {
            // prev == nil || prev.Val != current.Val
            prev = current
        }
    }
    return head
}
