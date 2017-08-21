package main

import "fmt"

func twoSum(nums []int, target int) []int {
	existMap := make(map[int]int, len(nums))
	for i2, v := range nums {
		if i1, exist := existMap[target-v]; exist {
			return []int{i1, i2}
		} else {
			existMap[v] = i2
		}
	}
	return nil
}

func main() {
	fmt.Println("hello world")
}
