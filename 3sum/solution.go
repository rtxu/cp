package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return make([][]int, 0)
	}

	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			if s == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if s > 0 {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, 4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
	fmt.Println(threeSum([]int{-1, -1, 2, -4, 2}))
	fmt.Println(threeSum([]int{-1, 2, -4, 2}))
	fmt.Println(threeSum([]int{-1, -1, 2, -4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}
