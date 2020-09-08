/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	var carry int

	addOneDigit := func(left, right int) {
		sum := left + right + carry
		node := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		carry = sum / 10
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	for l1 != nil || l2 != nil || carry > 0 {
		var left, right int
		if l1 != nil {
			left = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			right = l2.Val
			l2 = l2.Next
		}
		addOneDigit(left, right)
	}
	return head
}
