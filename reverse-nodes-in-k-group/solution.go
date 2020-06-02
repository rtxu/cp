/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    var next *ListNode
    i := 1
    var kHead, lastKHead *ListNode
    oldHead := head
    var newHead *ListNode
    
    for head != nil {
        next = head.Next
        
        if i == 1 {
            kHead = head
        }
        if i == k {
            i = 0
            head.Next = nil
            reverse(kHead)
            if newHead == nil {
                newHead = head
            }
            kHead.Next = next
            if lastKHead != nil {
                lastKHead.Next = head
            }
            lastKHead = kHead
        }
        
        head = next
        i++
    }
    
    if newHead != nil {
        return newHead
    } else {
        return oldHead
    }
}
