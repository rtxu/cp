func umax(current *int, val int) {
    if val > *current {
        *current = val
    }
}

func longestAwesome(s string) int {
    var digitOdd [10]bool
    n := len(s)
    prefix := make(map[[10]bool]int)
    prefix[digitOdd] = -1
    maxLen := 1
    for i := 0; i < n; i++ {
        d := int(s[i] - '0')
        digitOdd[d] = !digitOdd[d]
        // fmt.Println("i = ", i, "s[i] = ", s[i] - '0', "digitOdd", digitOdd)
        
        // all even
        if j, exist := prefix[digitOdd]; exist {
            umax(&maxLen, i - j)
        }
        // 1 odd
        for d := 0; d < 10; d++ {
            next := digitOdd
            next[d] = !digitOdd[d]
            if j, exist := prefix[next]; exist {
                umax(&maxLen, i - j)
            }
        }
        
        if _, exist := prefix[digitOdd]; !exist {
            prefix[digitOdd] = i
        }
    }
    return maxLen
}
