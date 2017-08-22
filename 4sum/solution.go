package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	if len(nums) < 4 {
		return make([][]int, 0)
	}

	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0)
	for i1 := 0; i1 < n-3; i1++ {
		if i1 > 0 && nums[i1] == nums[i1-1] {
			continue
		}
		for i2 := i1 + 1; i2 < n-2; i2++ {
			if i2 > i1+1 && nums[i2] == nums[i2-1] {
				continue
			}
			j := i2 + 1
			k := n - 1
			for j < k {
				s := nums[i1] + nums[i2] + nums[j] + nums[k]
				if s == target {
					result = append(result, []int{nums[i1], nums[i2], nums[j], nums[k]})
					j++
					for j < k && nums[j] == nums[j-1] {
						j++
					}
					k--
					for j < k && nums[k] == nums[k+1] {
						k--
					}
				} else if s > target {
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
	}

	return result
}

func main() {
	fmt.Println(fourSum([]int{-1, 0, 1, 0, 2, -2}, 0))
}
