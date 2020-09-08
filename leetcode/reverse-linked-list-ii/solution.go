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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    k := 0
    var mPrev, mNode, nNode, nNext *ListNode
    for current := head; current != nil; current = current.Next {
        k++
        if k == m-1 {
            mPrev = current
        }
        if k == m {
            mNode = current
        }
        if k == n {
            nNode = current
            nNext = nNode.Next
            nNode.Next = nil
            reverse(mNode, nil)
            if m == 1 {
                head = nNode
            } else {
                mPrev.Next = nNode
            }
            mNode.Next = nNext
        }
    }
    return head
}
