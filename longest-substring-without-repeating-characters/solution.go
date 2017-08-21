func lengthOfLongestSubstring(s string) int {
	substring := make(map[rune]int)
	length := 0
	for pos, char := range s {
		if repeatedPos, exist := substring[char]; exist {
			for i := pos - len(substring); i < repeatedPos; i++ {
				delete(substring, rune(s[i]))
			}
			substring[char] = pos
		} else {
			substring[char] = pos
			newLength := len(substring)
			if newLength > length {
				length = newLength
			}
		}
	}
	return length
}
