func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func index(i int) int { return i-1 }

func construct(s1, s2 string, LCS [][]int, i, j int) string {
    if i == 0 || j == 0 {
        return ""
    }
    if LCS[i][j] == LCS[i-1][j] {
        return construct(s1, s2, LCS, i-1, j)
    }
    if LCS[i][j] == LCS[i][j-1] {
        return construct(s1, s2, LCS, i, j-1)
    }
    if s1[index(i)] == s2[index(j)] && LCS[i][j] == LCS[i-1][j-1] + 1 {
        return construct(s1, s2, LCS, i-1, j-1) + string(s1[index(i)])
    }
    
    panic(fmt.Sprintf("unexpected s1=%s, s2=%s, i=%d, j=%d", s1, s2, i, j))
    return ""
}

func shortestCommonSupersequence(str1 string, str2 string) string {
    // 1. compute LCS(i, j)
    n1, n2 := len(str1), len(str2)
    LCS := make([][]int, n1+1)
    for i := 0; i <= n1; i++ {
        LCS[i] = make([]int, n2+1)
    }
    for i := 1; i <= n1; i++ {
        for j := 1; j <= n2; j++ {
            LCS[i][j] = max(LCS[i-1][j], LCS[i][j-1])
            if str1[index(i)] == str2[index(j)] {
                LCS[i][j] = max(LCS[i][j], LCS[i-1][j-1]+1)
            }
        }
    }
    
    // 2. reconstrcut LCS based on LCS(i, j)    
    lcs := construct(str1, str2, LCS, n1, n2)
    fmt.Println("lcs =", lcs)
    
    // 3. label the LCS position in str1 and str2 seperately
    listLcsPosition := func(s, lcs string) []int {
        pos := make([]int, 0)
        j := 0
        for i := 0; i < len(s) && j < len(lcs); i++ {
            if s[i] == lcs[j] {
                pos = append(pos, i)
                j++
            }
        }
        // 将 end of string 作为 lcs 的最后一个字符
        // 好处：所有未出现在 lcs 中的字符，后面都有 lcs 字符，方便后续处理
        pos = append(pos, len(s))
        return pos
    }
    pos1, pos2 := listLcsPosition(str1, lcs), listLcsPosition(str2, lcs)
    
    // 4. 将 str2 中未被标记的字符按照顺序插入到 str1 中的相应位置
    // 例：str1 = abaca str2 = mamama
    // 将 str2 的第一个 m 插入到 str1 第一个 a 之前，以此类推
    var result string
    lastLcsPosition1, lastLcsPosition2 := -1, -1
    for i := 0; i < len(pos2); i++ {
        nonLcsSeq := str2[lastLcsPosition2+1:pos2[i]]
        lastLcsPosition2 = pos2[i]
        result += nonLcsSeq + str1[lastLcsPosition1+1:min(pos1[i]+1, n1)]
        lastLcsPosition1 = pos1[i]
    }
    
    
    return result
}
