const INF = 1e9

func umin(current *int, val int) {
    if val < *current {
        *current = val
    }
}

func ceil(n, k int) int {
    return 1 + (n-1)/k
}

func countBits(n int) int {
    var res int
    for i := n; i > 0; i -= i&-i {
        res++
    }
    return res
}

func minNumberOfSemesters(n int, deps [][]int, k int) int {
    pre := make([]int, n)
    for i := 0; i < len(deps); i++ {
        x, y := deps[i][0], deps[i][1]
        x--
        y--
        pre[y] |= (1 << x)
    }
    T := make([]int, 1<<n)
    statePre := make([]int, 1<<n)
    for s := 1; s < 1<<n; s++ {
        T[s] = INF
        for j := 0; j < n; j++ {
            if s & (1<<j) > 0 {
                statePre[s] |= pre[j]
            }
        }
    }
    for s := 1; s < 1<<n; s++ {
        for i := (s-1)&s; ; i = (i-1)&s {
            todo := s - i
            if statePre[todo] & i == statePre[todo] {
                umin(&T[s], T[i] + ceil(countBits(todo), k))
            }
            if i == 0 {
                break
            }
        }
    }
    return T[(1<<n)-1]
}

// T(s) 表示 s 状态需要的最小 semester 数，
// s 为 n-bit 二进制数，第 i bit 为 1 表示此课已上
// 显然，T(0) = 0
// T(s) = min{T(i) + ceil(countBits(s-i), k)}, i 是 s 的 submask && (s-i) 中的课 degree 均为 0
// i 表示已经上完的课，s-i 即为要上的课
