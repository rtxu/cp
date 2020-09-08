/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(current, prev *ListNode) *ListNode {
    if current == nil {
        return prev
    } else {
        head := reverse(current.Next, current)
        current.Next = prev
        return head
    }
}

func reverseList(head *ListNode) *ListNode {
    return reverse(head, nil)
}
