func numDecodings(s string) int {
    // A(i, j) 表示满足 s[0:i] 以 j 结尾的方案数
    // A(i, j) = 
    // 1) 如果 j == s 的最后一位数字，即 s[i-1]，Sum(A(i-1, k))，any k
    // 2) 如果 j == s 的最后两位数字，即 s[i-2:i]，Sum(A(i-2, k)), any k
    // 令 sum[i] 表示第 i 行的和，则 sum[n] 为最终答案
    // 初始条件：sum[0] = 1，代表「一个数字不取」这一唯一方案
    
    n := len(s)
    sum := make([]int, n+1)
    sum[0] = 1
    A := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        A[i] = make([]int, 27)
    }
    
    last, last2 := 0, 0
    for i := 1; i <= n; i++ {
        last = int(s[i-1] - '0')
        last2 = (last2 * 10 + last) % 100
        
        if last > 0 {
            A[i][last] = sum[i-1]
            sum[i] += A[i][last]
        }
        if last2 >= 10 && last2 <= 26 {
            A[i][last2] = sum[i-2]
            sum[i] += A[i][last2]
        }
    }
    
    return sum[n]
}
