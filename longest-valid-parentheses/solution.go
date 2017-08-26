package main

import "fmt"

func longestValidParentheses(s string) int {
	n := len(s)
	maxLength := 0
	validLeft := make(map[int]int)
	currentHeight := 0
	validLeft[currentHeight] = 0
	for i := 1; i <= n; i++ {
		v := s[i-1]
		if v == '(' {
			currentHeight++
			validLeft[currentHeight] = i
		} else {
			delete(validLeft, currentHeight)
			currentHeight--
			if left, exist := validLeft[currentHeight]; exist {
				if i-left > maxLength {
					maxLength = i - left
				}
			} else {
				validLeft[currentHeight] = i
			}
		}
	}
	return maxLength
}

func main() {
	input, expected := "(())", 4
	fmt.Printf("intput: %s, expected: %d, actual: %d\n", input, expected, longestValidParentheses(input))
	input, expected = "((((())", 4
	fmt.Printf("intput: %s, expected: %d, actual: %d\n", input, expected, longestValidParentheses(input))
	input, expected = "(())))))", 4
	fmt.Printf("intput: %s, expected: %d, actual: %d\n", input, expected, longestValidParentheses(input))
	input, expected = "))))(())", 4
	fmt.Printf("intput: %s, expected: %d, actual: %d\n", input, expected, longestValidParentheses(input))
	input, expected = "((((", 0
	fmt.Printf("intput: %s, expected: %d, actual: %d\n", input, expected, longestValidParentheses(input))
	input, expected = "))))", 0
	fmt.Printf("intput: %s, expected: %d, actual: %d\n", input, expected, longestValidParentheses(input))
	input, expected = "", 0
	fmt.Printf("intput: %s, expected: %d, actual: %d\n", input, expected, longestValidParentheses(input))
	input, expected = "(((((()()()(((((", 6
	fmt.Printf("intput: %s, expected: %d, actual: %d\n", input, expected, longestValidParentheses(input))
	input, expected = ")(", 0
	fmt.Printf("intput: %s, expected: %d, actual: %d\n", input, expected, longestValidParentheses(input))
	input, expected = "()()())))((((()))))(((((", 10
	fmt.Printf("intput: %s, expected: %d, actual: %d\n", input, expected, longestValidParentheses(input))
	input, expected = "()()())))(((((((((())))(((((", 8
	fmt.Printf("intput: %s, expected: %d, actual: %d\n", input, expected, longestValidParentheses(input))
	input, expected = "(())((((((((((()()()(())", 10
	fmt.Printf("intput: %s, expected: %d, actual: %d\n", input, expected, longestValidParentheses(input))
	input, expected = ")()())()()(", 4
	fmt.Printf("intput: %s, expected: %d, actual: %d\n", input, expected, longestValidParentheses(input))
}
