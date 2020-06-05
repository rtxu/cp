/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func list(head *ListNode) []int {
    var result []int
    for current := head; current != nil; current = current.Next {
        result = append(result, current.Val)
    }
    return result
}

func oddEvenList(head *ListNode) *ListNode {
    k := 1
    var next, evenHead *ListNode
    for current := head; current != nil; current = next {
        //fmt.Println(k, list(head), list(evenHead))
        if k == 2 {
            evenHead = current
        }
        next = current.Next
        if next != nil {
            current.Next = next.Next
        }
        if k % 2 == 1 && (next == nil || next.Next == nil) {
            // the last odd node
            if k == 1 {
                current.Next = next
            } else {
                current.Next = evenHead
            }
        }
        k++
    }
    return head
}
