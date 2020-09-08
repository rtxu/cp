package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type intHeap []*ListNode

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
	//fmt.Printf("Push %d\n", x.(*ListNode).Val)
}
func (h *intHeap) Pop() interface{} {
	tmp := *h
	n := len(*h)
	*h = (*h)[0 : n-1]
	//fmt.Printf("Pop %d\n", tmp[n-1].Val)
	return tmp[n-1]
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &intHeap{}
	heap.Init(h)
	for _, v := range lists {
		if v != nil {
			heap.Push(h, v)
		}
	}

	var head, tail *ListNode
	for h.Len() > 0 {
		cur := heap.Pop(h).(*ListNode)
		if cur.Next != nil {
			heap.Push(h, cur.Next)
		}
		if head == nil {
			head = cur
			tail = cur
		} else {
			tail.Next = cur
			tail = cur
		}
	}
	return head
}

func printList(head *ListNode) {
	fmt.Printf("[")
	for head != nil {
		fmt.Printf("%d, ", head.Val)
		head = head.Next
	}
	fmt.Printf("]")
}

func main() {
	h1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	h2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	head := mergeKLists([]*ListNode{h1, h2})
	printList(head)
}
