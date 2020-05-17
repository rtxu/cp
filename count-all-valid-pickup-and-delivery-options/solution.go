var M = make(map[int]int)

const MOD = 1e9+7

func countOrders(n int) int {
    if n == 1 {
        return 1
    }
    if v, exist := M[n]; exist {
        return v
    }
    
    v := n * (2*n-1)
    w := countOrders(n-1)
    M[n] = int(int64(v) * int64(w) % MOD)
    
    return M[n]
}
