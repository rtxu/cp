/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    n := 0
    var last *ListNode
    for current := head; current != nil; current = current.Next {
        n++
        last = current
    }
    k %= n
    var newHead *ListNode
    if k == 0 {
        return head
    } else {
        var newTail *ListNode
        i := 1
        for newTail = head; i != n-k; newTail = newTail.Next {i++}
        newHead = newTail.Next
        newTail.Next = nil
        last.Next = head
    }
    
    return newHead
}
