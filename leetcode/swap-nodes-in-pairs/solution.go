/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    k := 1
    var prev, next, lastTail *ListNode
    for current := head; current != nil; current = next {
        next = current.Next
        if k % 2 == 0 {
            prev.Next = current.Next
            current.Next = prev
            
            if k == 2 {
                head = current
            }
            
            if lastTail != nil {
                lastTail.Next = current
            }
            lastTail = prev
        } else {
            prev = current
        }
        k++
    }
    return head
}
