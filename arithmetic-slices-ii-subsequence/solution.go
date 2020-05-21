func numberOfArithmeticSlices(A []int) int {
    // 枚举序列的前两个数(i, j)，
    // 以 seq[len] 表示以 (i, j) 为最前面两个数字，长度为 len 的子序列数
    // 初始状态 seq[2] = 1
    // 从左到右按序遍历 j 后的数字 k, 确定 A[k] 在 (i, j) 序列中的位置，如：第 5 个，如果长度为 4 的序列已经出现了 10 个，则长度为 5 的序列又增加了 10 个
    
    var result int
    n := len(A)
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            diff := A[j] - A[i]
            if diff == 0 {
                cnt := 0
                for k := j+1; k < n; k++ {
                    if A[k] == A[i] {
                        cnt++
                    }
                }
                result += ((1<<cnt) - 1)
            } else {
                seq := make([]int, n+1)
                seq[2] = 1
                for k := j+1; k < n; k++ {
                    if (A[k] - A[i]) % diff == 0 {
                        expLen := (A[k] - A[i]) / diff + 1
                        if expLen >= 3 && expLen <= n && seq[expLen-1] > 0 {
                            seq[expLen] += seq[expLen-1]
                            result += seq[expLen-1]
                            //fmt.Println("expLen=", expLen, "prefix=", A[i], A[j], A[k])
                        }
                    }
                }
            }
        }
    }
    return result
}
