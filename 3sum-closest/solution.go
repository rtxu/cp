package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	result := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(s-target)) < math.Abs(float64(result-target)) {
				result = s
			}
			if s == target {
				return target
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
	return result
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 0, 1, 2, -1, 4}, 1))
}
