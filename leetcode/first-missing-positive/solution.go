package main

import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)
	numSet := make(map[int]int, n)
	for i := 0; i < n; i++ {
		numSet[nums[i]] = 0
	}

	for i := 1; ; i++ {
		if _, exist := numSet[i]; !exist {
			return i
		}
	}
	return -1
}

func main() {
	var nums []int
	var expected int

	nums, expected = []int{1, 2, 0}, 3
	fmt.Printf("input: %v, expected: %d, actual: %d\n", nums, expected, firstMissingPositive(nums))
	nums, expected = []int{3, 4, -1, 1}, 2
	fmt.Printf("input: %v, expected: %d, actual: %d\n", nums, expected, firstMissingPositive(nums))
}
