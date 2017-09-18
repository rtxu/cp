package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap struct {
	IntHeap
}
type MaxHeap struct {
	IntHeap
}

func (h MaxHeap) Less(i, j int) bool { return h.IntHeap[i] > h.IntHeap[j] }

type MedianFinder struct {
	left  *MaxHeap
	right *MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	m := MedianFinder{
		left: &MaxHeap{
			IntHeap: []int{},
		},
		right: &MinHeap{
			IntHeap: []int{},
		},
	}
	heap.Init(m.left)
	heap.Init(m.right)
	return m
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == this.right.Len() {
		heap.Push(this.left, num)
	} else {
		heap.Push(this.right, num)
	}

	if this.right.Len() > 0 && (*this.left).IntHeap[0] > (*this.right).IntHeap[0] {
		leftVal := heap.Pop(this.left)
		rightVal := heap.Pop(this.right)
		heap.Push(this.right, leftVal)
		heap.Push(this.left, rightVal)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() == this.right.Len() {
		return float64((*this.left).IntHeap[0]+(*this.right).IntHeap[0]) / 2
	} else {
		return float64((*this.left).IntHeap[0])
	}
}

func main() {
	min := &MinHeap{
		IntHeap: []int{},
	}
	heap.Init(min)
	heap.Push(min, 1)
	heap.Push(min, 2)
	heap.Push(min, 3)
	for min.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(min))
	}
	fmt.Println()
	max := &MaxHeap{
		IntHeap: []int{},
	}
	heap.Init(max)
	heap.Push(max, 1)
	heap.Push(max, 2)
	heap.Push(max, 3)
	for max.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(max))
	}
	fmt.Println()

	{
		m := Constructor()
		m.AddNum(1)
		m.AddNum(2)
		fmt.Println(m.FindMedian())
		m.AddNum(3)
		fmt.Println(m.FindMedian())
	}
	{
		m := Constructor()
		m.AddNum(1)
		m.AddNum(3)
		m.AddNum(7)
		m.AddNum(9)
		m.AddNum(5)
		fmt.Println(m.FindMedian())
	}
	{
		m := Constructor()
		m.AddNum(1)
		m.AddNum(3)
		m.AddNum(7)
		m.AddNum(9)
		m.AddNum(8)
		fmt.Println(m.FindMedian())
	}
	{
		m := Constructor()
		m.AddNum(1)
		m.AddNum(3)
		m.AddNum(5)
		m.AddNum(7)
		m.AddNum(9)
		m.AddNum(6)
		fmt.Println(m.FindMedian())
	}
	{
		m := Constructor()
		m.AddNum(1)
		m.AddNum(3)
		m.AddNum(5)
		m.AddNum(7)
		m.AddNum(9)
		m.AddNum(2)
		fmt.Println(m.FindMedian())
	}
}
