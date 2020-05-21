const MOD = 1e9+7

func sumSubseqWidths(A []int) int {
    sort.Ints(A)
    n := len(A)
    var result int
    x := 0
    pow2, pow2Sum := 1, 1
    for i := n-1; i >= 1; i-- {
        x = ((x << 1) + int(int64(A[i]-A[i-1]) * int64(pow2Sum) % MOD)) % MOD
        result += x
        if result > MOD {
            result -= MOD
        }
        pow2 <<= 1
        if pow2 > MOD {
            pow2 -= MOD
        }
        pow2Sum += pow2
        if pow2Sum > MOD {
            pow2Sum -= MOD
        }
    }
    
    return result
}
