package main

import "fmt"

func longestConsecutive(nums []int) int {
	rangeBegin := make(map[int]int) // map range begin to end
	rangeEnd := make(map[int]int)   // map range end to begin
	processed := make(map[int]bool)

	n := len(nums)
	for i := 0; i < n; i++ {
		cur := nums[i]
		if processed[cur] {
			continue
		}
		var beginExtend, endExtend bool
		if _, exist := rangeEnd[cur-1]; exist {
			rangeEnd[cur] = rangeEnd[cur-1]
			rangeBegin[rangeEnd[cur-1]]++
			endExtend = true
			delete(rangeEnd, cur-1)
		}
		if _, exist := rangeBegin[cur+1]; exist {
			rangeBegin[cur] = rangeBegin[cur+1]
			rangeEnd[rangeBegin[cur+1]]--
			beginExtend = true
			delete(rangeBegin, cur+1)
		}
		if beginExtend && endExtend {
			begin, end := rangeEnd[cur], rangeBegin[cur]
			delete(rangeBegin, rangeEnd[cur])
			delete(rangeEnd, cur)
			delete(rangeEnd, rangeBegin[cur])
			delete(rangeBegin, cur)
			rangeBegin[begin] = end
			rangeEnd[end] = begin
		} else if !beginExtend && !endExtend {
			rangeBegin[cur] = cur
			rangeEnd[cur] = cur
		}
		processed[cur] = true

		/*
			fmt.Println("Round: ", i)
			for begin, end := range rangeBegin {
				fmt.Println(begin, end)
			}
		*/
	}

	var maxLength int
	for begin, end := range rangeBegin {
		if (end - begin + 1) > maxLength {
			maxLength = end - begin + 1
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
