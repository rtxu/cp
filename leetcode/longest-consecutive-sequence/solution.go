package main

import "fmt"

func longestConsecutive(nums []int) int {
	n := len(nums)
	numSet := make(map[int]int, n)
	for i := 0; i < n; i++ {
		numSet[nums[i]] = 1
	}

	var maxLength int
	for i, _ := range numSet {
		if _, exist := numSet[i-1]; !exist {
			// now i is the begin
			j := i + 1
			for {
				if _, exist := numSet[j]; exist {
					j++
				} else {
					break
				}
			}
			if j-i > maxLength {
				maxLength = j - i
			}
		}
	}
	return maxLength
}

func main() {
	var input []int
	var expected int
	input, expected = []int{100, 4, 200, 1, 3, 2}, 4
	fmt.Printf("input: %v, expected: %d, actual: %d\n", input, expected, longestConsecutive(input))
	input, expected = []int{-7, -1, 3, -9, -4, 7, -3, 2, 4, 9, 4, -9, 8, -7, 5, -1, -7}, 4
	fmt.Printf("input: %v, expected: %d, actual: %d\n", input, expected, longestConsecutive(input))
	input, expected = []int{7, -9, 3, -6, 3, 5, 3, 6, -2, -5, 8, 6, -4, -6, -4, -4, 5, -9, 2, 7, 0, 0}, 4
	fmt.Printf("input: %v, expected: %d, actual: %d\n", input, expected, longestConsecutive(input))
}
