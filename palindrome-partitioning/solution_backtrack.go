func backtrack(start int, s string, current []string, P [][]bool, result *[][]string) {
    end := len(s)
    if start == end {
        currentCopy := make([]string, len(current))
        copy(currentCopy, current)
        *result = append(*result, currentCopy)
        return 
    }
    for j := start; j < end; j++ {
        if P[start][j] {
            current = append(current, s[start:j+1])
            backtrack(j+1, s, current, P, result)
            current = current[0:len(current)-1]
        }
    }
}

func partition(s string) [][]string {
    // P(i, j) 表示 s[i:j+1] 是否为回文串, 0 <= i < n, i <= j < n
    // 显然
    // 1) P(i，i) 皆为回文
    // 2) 空串，j < i 也为回文
    
    n := len(s)
    P := make([][]bool, n)
    for i := 0; i < n; i++ {
        P[i] = make([]bool, n)
    }
    
    for i := n-1; i >= 0; i-- {
        P[i][i] = true
        for j := i+1; j < n; j++ {
            if s[i] == s[j] && (j-1 < i+1 || P[i+1][j-1]) {
                P[i][j] = true
            }
        }
        //fmt.Println(P[i])
    }
    
    var result [][]string
    backtrack(0, s, []string{}, P, &result)
    return result
}
