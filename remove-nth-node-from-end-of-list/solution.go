package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, q := &head, &head

	for *p != nil {
		p = &((*p).Next)
		n--
		if n < 0 {
			q = &((*q).Next)
		}
	}
	if q == &head {
		*q = (*q).Next
		return *q
	} else {
		*q = (*q).Next
		return head
	}
}

func main() {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}

	p := head
	for i := 1; i < 10; i++ {
		p.Next = &ListNode{
			Val:  i,
			Next: nil,
		}
		p = p.Next
	}

	printList := func(head *ListNode) {
		fmt.Printf("[")
		for head != nil {
			fmt.Printf("%d, ", head.Val)
			head = head.Next
		}
		fmt.Printf("]")
	}

	printList(head)
	head = removeNthFromEnd(head, 1)
	fmt.Println("\nafter remove the 1st from end")
	printList(head)
	head = removeNthFromEnd(head, 2)
	fmt.Println("\nafter remove the 2nd  from end")
	printList(head)
	head = removeNthFromEnd(head, 8)
	fmt.Println("\nafter remove the 8th  from end")
	printList(head)

}
