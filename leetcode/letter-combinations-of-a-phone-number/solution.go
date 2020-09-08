package main

import "fmt"

var digit2Letters = map[uint8]string{
	'1': "",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
	'0': " ",
}
var resultSlice []string
var result []byte

func dfs(digits string) {
	if len(digits) == 0 {
		if len(result) > 0 {
			resultSlice = append(resultSlice, string(result))
		}
		return
	}

	d := digits[0]
	for i := 0; i < len(digit2Letters[d]); i++ {
		result = append(result, digit2Letters[d][i])
		dfs(digits[1:])
		result = result[:len(result)-1]
	}
}

func letterCombinations(digits string) []string {
	resultSlice = make([]string, 0)
	result = make([]byte, 0, len(digits))
	dfs(digits)
	return resultSlice
}

func main() {
	fmt.Printf("input: 3 expected: [d, e, f], actual: %v\n", letterCombinations("3"))
	fmt.Printf("input: 23 expected: [ad, ae, af, bd, be, bf, cd, ce, cf], actual: %v\n", letterCombinations("23"))
}
