const MOD = 1e9+7

const N = 105
const F = 205

func abs(n int) int {
    if n < 0 {
        return -n
    } else {
        return n
    }
}

func countRoutes(locations []int, start int, finish int, fuel int) int {
    // T(i, j) 表示从 start 到 i 耗油 j 的方案数
    // Sum{T(finish, j), 0 <= j <= fuel } 为答案
    // T(i, j) = Sum{T(k, j - |location[k] - locations[i]|), 0 <= k < n && k != i}
    
    n := len(locations)
    var T [N][F]int
    T[start][0] = 1
    for j := 1; j <= fuel; j++ {
        for i := 0; i < n; i++ {
            sum := 0
            for k := 0; k < n; k++ {
                if k == i { continue }
                if j - abs(locations[k] - locations[i]) >= 0 {
                    sum += T[k][j - abs(locations[k] - locations[i])]
                    sum %= MOD
                }
            }
            T[i][j] = sum
        }
    }
    result := 0
    for j := 0; j <= fuel; j++ {
        result += T[finish][j]
        result %= MOD
    }
    return result
}

