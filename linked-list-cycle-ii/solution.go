/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ref : https://leetcode.com/problems/linked-list-cycle-ii/discuss/44781/Concise-O(n)-solution-by-using-C%2B%2B-with-Detailed-Alogrithm-Description
func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    slow, fast, entry := head, head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            for slow != entry {
                slow = slow.Next
                entry = entry.Next
            }
            return entry
        }
    }
    return nil
}
