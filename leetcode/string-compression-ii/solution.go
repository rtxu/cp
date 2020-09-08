/*
"bbabbbabbbbcbb"
4
"aaabcccd"
2
"aabbaa"
2
"aaaaaaaaaaa"
0
"abaabba"
5
*/

type Node struct {
    i, k, L int
    op string
}

var T [105][105][105]int
var pre [105][105][105]Node
var next [105]int

func getLengthOfOptimalCompression(s string, k int) int {
    // T(i, k, L) [i, n) 的范围内可删除 k 个字符且左边有 L 个与 s[i] 相同的字符 情况下的最优解 
    // T(0, k, 0) 为最终答案
    
    n := len(s)

    for i := 0; i < n; i++ {
        for j := 0; j <= k; j++ {
            for L := 0; L < n; L++ {
                T[i][j][L] = -1
                pre[i][j][L] = Node{}
            }
        }
    }
    last := make([]int, 26)
    for i := 0; i < 26; i++ {
        last[i] = -1
    }
    for i := 0; i < n; i++ {
        next[i] = n
        ch := s[i] - 'a'
        if last[ch] >= 0 {
            next[last[ch]] = i
        }
        last[ch] = i
    }
    
    v := solve(s, 0, k, 0)
    // dumpResult(0, k, 0)
    return v
}

func dumpResult(i, k, L int) {
    var emptyNode Node
    for pre[i][k][L] != emptyNode {
        fmt.Printf("state: {%d, %d, %d} result: %d <--- %s ---\n", i, k, L, T[i][k][L], pre[i][k][L].op)
        i, k, L = pre[i][k][L].i, pre[i][k][L].k, pre[i][k][L].L
    }
    fmt.Printf("state: {%d, %d, %d} result: %d\n", i, k, L, T[i][k][L])
}

func encodeOne(cnt int) int {
    if cnt <= 0 { 
        return 0 
    } else if cnt == 1 {
        return 1
    } else {
        digitCnt := 0
        for cnt > 0 {
            digitCnt++
            cnt /= 10
        }
        return 1 + digitCnt
    }
}

func encode(s string, L int) int {
    n := len(s)
    if n == 0 {
        return encodeOne(L)
    } else {
        s += "$"
        c := s[0]
        cnt := L+1
        result := 0
        for i := 1; i < n+1; i++ {
            if s[i] == c {
                cnt++
            } else {
                result += encodeOne(cnt)
                c = s[i]
                cnt = 1
            }
        }
        return result
    }
    
}

func solve(s string, i, k, L int) int {
    n := len(s)
    if i >= n {
        return 0
    }
    v := T[i][k][L]
    if v > 0 {
        return v
    }
    
    if k == 0 {
        v = encode(s[i:], L)
        T[i][k][L] = v
        return v
    }
    
    if n-i <= k {
        v = encodeOne(L - (k-(n-i)))
        T[i][k][L] = v
        return v
    }
    
    v = math.MaxInt32
    // delete s[i]
    var nv, nL int
    if i+1 < n {
        if s[i+1] == s[i] {
            nL = L
            nv = solve(s, i+1, k-1, nL)
            if nv < v {
                v = nv
                pre[i][k][L] = Node{i+1, k-1, nL, "delete-1"}
            }
        } else {
            nv = encodeOne(L) + solve(s, i+1, k-1, 0)
            if nv < v {
                v = nv
                pre[i][k][L] = Node{i+1, k-1, 0, "delete-2"}
            }
        }
    }
    
    // keep s[i]
    nv = encodeOne(L+1) + solve(s, i+1, k, 0)
    if nv < v {
        v = nv
        pre[i][k][L] = Node{i+1, k, 0, "merge-current"}
    }
    j := next[i]
    if j < n && j-i-1 <= k {
        nv = solve(s, j, k - (j-i-1), L+1)
        if nv < v {
            v = nv 
            pre[i][k][L] = Node{j, k-(j-i-1), L+1, "merge-next"}
        }
    }
    
    
    T[i][k][L] = v
    return v
}

