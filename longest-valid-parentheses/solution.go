package main

import "fmt"

func longestValidParentheses(s string) int {
	n := len(s)
	stackTop := make([]int, n+1)
	for i := 0; i <= n; i++ {
		stackTop[i] = -1 // invalid
	}
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			if stackTop[i] == -1 {
				stackTop[i+1] = 1
			} else {
				stackTop[i+1] = stackTop[i] + 1
			}
		} else {
			if stackTop[i] > 0 {
				stackTop[i+1] = stackTop[i] - 1
			}
		}
	}

	maxLength := 0
	stackTopMinPos := make([]int, n+1)
	//fmt.Print("[")
	for i := 0; i <= n; i++ {
		stackTopMinPos[i] = -1 // invalid
		//fmt.Printf("%d, ", stackTop[i])
	}
	//fmt.Println("]")

	for i := 0; i <= n; i++ {
		if stackTop[i] == -1 {
			if stackTopMinPos[0] != -1 {
				// leave a complete stack
				stackTopMinPos[0] = -1
			}
		} else {
			if stackTopMinPos[0] == -1 {
				// enter a new stack
				stackTopMinPos[0] = i - 1
			}
			if stackTopMinPos[stackTop[i]] == -1 {
				stackTopMinPos[stackTop[i]] = i
			} else {
				if i-stackTopMinPos[stackTop[i]] > maxLength {
					maxLength = i - stackTopMinPos[stackTop[i]]
					//fmt.Printf("(%d, %d]\n", stackTopMinPos[stackTop[i]], i)
				}
				stackTopMinPos[stackTop[i]+1] = -1
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
