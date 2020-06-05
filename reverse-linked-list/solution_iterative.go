/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    var prev, next *ListNode
    for current := head; current != nil; current = next {
        next = current.Next
        current.Next = prev
        prev = current
    }
    return prev
}
