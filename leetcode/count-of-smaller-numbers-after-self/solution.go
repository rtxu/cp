package main

import "fmt"

type ElementType struct {
	Value        int
	OriginalPos  int
	SmallerCount int
}

func mergeSort(array []ElementType) {
	n := len(array)
	if n == 0 || n == 1 {
		return
	}

	mergeSort(array[0 : n/2])
	mergeSort(array[n/2 : n])

	temp := make([]ElementType, n)
	tempLen := 0
	i, j := 0, n/2
	for i < n/2 && j < n {
		if array[i].Value > array[j].Value {
			array[i].SmallerCount += n - j
			temp[tempLen] = array[i]
			i++
		} else {
			temp[tempLen] = array[j]
			j++
		}
		tempLen++
	}
	for i < n/2 {
		temp[tempLen] = array[i]
		tempLen++
		i++
	}
	for j < n {
		temp[tempLen] = array[j]
		tempLen++
		j++
	}
	for i := 0; i < n; i++ {
		array[i] = temp[i]
	}
}

func countSmaller(nums []int) []int {
	n := len(nums)
	array := make([]ElementType, n)
	for i := 0; i < n; i++ {
		array[i] = ElementType{nums[i], i, 0}
	}

	mergeSort(array)

	count := make([]int, n)
	for i := 0; i < n; i++ {
		e := array[i]
		count[e.OriginalPos] = e.SmallerCount
	}
	return count
}

func main() {
	var input, expected []int
	input, expected = []int{5, 2, 6, 1}, []int{2, 1, 1, 0}
	fmt.Printf("input: %v, expected: %v, actual: %v", input, expected, countSmaller(input))
}
