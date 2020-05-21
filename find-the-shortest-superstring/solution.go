/*
["alex","loves","leetcode"]
["catg","ctaagt","gcta","ttca","atgcatc"]
["wmiy","yarn","rnnwc","arnnw","wcj"]
*/

func shortestSuperstring(A []string) string {
    if len(A) == 0 {
        return ""
    }
    // S(i, j) 表示 i 这个二进制状态组合下以 j 串作为最后一个串所得的最短长度
    // 关键信息：A 中不存在一个串是另一个串的子串的情况，所以「以哪个串结尾」可以作为无后效性的状态
    // 如果存在一个串是另一个串的子串，那么直接去除掉子串即可
    // 0 <= i < (1 << len(A)), 0 <= j < len(A)
    // S(i, j) = 
    // 1) 当 i & (1<<j) == 0 时，-1 不符合条件
    // 2) 当 i & (1<<j) >  0 时，说明 i 状态包含 j
    //    2.1) 将 j append 到最后，枚举 i 状态去除 j 以后的所有结果，S(i - (1<<j), k)，0 <= k < len(A)
    //    S(i, j) = min(S(i - (1<<j), k) + 额外长度)
    // 复杂度：2^12*12*12
    
    n := len(A)
    
    extra := make([][]int, n)
    for i := 0; i < n; i++ {
        extra[i] = make([]int, n)
        for j := 0; j < n; j++ {
            if i != j {
                extraLen := len(A[j])
                for k := 0; k < len(A[i]); k++ {
                    suffix := A[i][k:]
                    if len(A[j]) > len(suffix) {
                        prefix := A[j][0:len(suffix)]
                        if suffix == prefix {
                            extraLen -= len(prefix)
                            break
                        }
                    } 
                }
                extra[i][j] = extraLen
            } else {
                extra[i][i] = 0
            }
        }
    }
    
    S := make([][]int, 1<<n)
    for i := 0; i < (1<<n); i++ {
        S[i] = make([]int, n)
        for j := 0; j < n; j++ {
            S[i][j] = -1
        }
    }
    for i := 0; i < (1<<n); i++ {
        for j := 0; j < n; j++ {
            if i & (1<<j) == 0 {
                
            } else {
                prevState := i - (1<<j)
                for k := 0; k < n; k++ {
                    if prevState == 0 {
                        S[i][j] = len(A[j])
                    } else if S[prevState][k] != -1 {
                        newLen := S[prevState][k] + extra[k][j]
                        if S[i][j] == -1 || newLen < S[i][j] {
                            S[i][j] = newLen
                        }
                    }
                }
            }
        }
        // fmt.Printf("%b %v\n", i, S[i])
    }
    
    const MAXLEN = 12*20+1
    state := (1<<n) - 1
    tailIndex := -1
    resultLen := MAXLEN
    for i := 0; i < n; i++ {
        if S[state][i] != -1 && S[state][i] < resultLen {
            resultLen = S[state][i]
            tailIndex = i
        }
    }
    state -= (1<<tailIndex)
    result := A[tailIndex]
    currentResultLen := len(result)
    fmt.Printf("tailIndex = %d str = %s\n", tailIndex, A[tailIndex])
    
    getCommonLen := func (i, j int) int {
        extraLen := extra[i][j]
        commonLen := len(A[j]) - extraLen
        return commonLen
    }
    
    for state > 0 {
        for i := 0; i < n; i++ {
            commonLen := getCommonLen(i, tailIndex)
            if S[state][i] != -1 && S[state][i] + currentResultLen - commonLen == resultLen {
                prefixLen := len(A[i]) - commonLen
                result = A[i][0:prefixLen] + result
                tailIndex = i
                state -= (1<<tailIndex)
                currentResultLen = len(result)
            }
        }
        fmt.Printf("tailIndex = %d str = %s\n", tailIndex, A[tailIndex])
    }
    
    return result
}
