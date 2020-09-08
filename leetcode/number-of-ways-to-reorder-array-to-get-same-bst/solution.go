const MOD = 1e9+7

var C [][]int

func numOfWays(nums []int) int {
    n := len(nums)
    C = make([][]int, n+1)
    for i := 0; i <= n; i++ {
        C[i] = make([]int, n+1)
    }
    C[0][0] = 1
    for i := 1; i <= n; i++ {
        C[i][0] = 1
        for j := 1; j <= i; j++ {
            C[i][j] = (C[i-1][j] + C[i-1][j-1]) % MOD
        }
    }
    res := solve(nums)
    return (res - 1 + MOD) % MOD
}


func solve(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return 1
    }
    nL, nR := 0, 0
    L, R := make([]int, 0), make([]int, 0)
    for i := 1; i < n; i++ {
        if nums[i] < nums[0] {
            nL++
            L = append(L, nums[i])
        } else {
            nR++
            R = append(R, nums[i])
        }
    }
    cntL, cntR := solve(L), solve(R)
    
    return int((((int64(cntL) * int64(cntR)) % MOD) * int64(C[nL+nR][nL])) % MOD)
}
