const MOD = 1e9+7

func sumSubseqWidths(A []int) int {
    sort.Ints(A)
    
    n := len(A)
    pow2 := make([]int, n)
    pow2[0] = 1
    for i := 1; i < n; i++ {
        pow2[i] = (pow2[i-1] << 1)
        if pow2[i] > MOD {
            pow2[i] -= MOD
        }
    }
    
    var result int
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if A[j] > A[i] {
                result += int(int64(A[j]-A[i]) * int64(pow2[j-i-1]) % MOD)
                if result > MOD {
                    result -= MOD
                }
            }
        }
    }
    
    return result
}
