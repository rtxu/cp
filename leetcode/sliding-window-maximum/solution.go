package main

import "fmt"

type ElementType struct {
	Value int
	Index int
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	Q := make([]ElementType, 0, n)

	max := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(Q) > 0 && nums[i] > Q[len(Q)-1].Value {
			Q = Q[:len(Q)-1]
		}
		Q = append(Q, ElementType{nums[i], i})

		if i >= k {
			if Q[0].Index == i-k {
				Q = Q[1:]
			}
		}

		if i >= k-1 {
			max = append(max, Q[0].Value)
		}
	}
	return max
}

func main() {
	var nums []int
	var k int
	var expected []int

	nums = []int{1, 3, -1, -3, 5, 3, 6, 7}
	k = 3
	expected = []int{3, 3, 5, 5, 6, 7}
	fmt.Printf("nums: %v, k: %d, expected: %v, actual: %v\n", nums, k, expected, maxSlidingWindow(nums, k))

	nums = []int{1, 3, -1, -3, 5, 3, 6, 7}
	k = 1
	expected = []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Printf("nums: %v, k: %d, expected: %v, actual: %v\n", nums, k, expected, maxSlidingWindow(nums, k))

	nums = []int{1, 3, -1, -3, 5, 3, 6, 7}
	k = len(nums)
	expected = []int{7}
	fmt.Printf("nums: %v, k: %d, expected: %v, actual: %v\n", nums, k, expected, maxSlidingWindow(nums, k))
}
