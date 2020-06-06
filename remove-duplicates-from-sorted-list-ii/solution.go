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
        if current.Next != nil && current.Val == current.Next.Val {
            for current.Next != nil && current.Val == current.Next.Val {
                current = current.Next
            }
            if prev == nil {
                head = current.Next
            } else {
                prev.Next = current.Next
            }
        } else {
            prev = current
        }
    }
    return head
}
