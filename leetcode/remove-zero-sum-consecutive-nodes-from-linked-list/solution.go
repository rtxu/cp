/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    rightMost := map[int]*ListNode{}
    sum := 0
    for current := head; current != nil; current = current.Next {
        sum += current.Val
        rightMost[sum] = current
    }
    if rightMost[0] != nil {
        head = rightMost[0].Next
    }
    sum = 0
    for current := head; current != nil; current = current.Next {
        sum += current.Val
        if rightMost[sum] != current {
            current.Next = rightMost[sum].Next
        }
    }
    return head
}
