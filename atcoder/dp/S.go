package main
import "fmt"

/*
i 的遍历顺序从最高位到最低位
T(i, prefix_remainder, isEqual) 表示第 i 位、
更高位的数字和 mod D 为 prefix_remainder、
状态为 isEqual(true 表示前缀与 K 的相应前缀相等，false < K 的相应前缀) 的方案数
则 T(hi, 0, true) 为最终答案，其中 hi = K 的最高位对应的 i，
比如：98765，则最高位 i = 5

状态转移：
T(i, prefix_remainder, isEqual) = 
Sum{ 
    T(i-1, prefix_remainder + d_i, isEqual && d_i == d[i])
    0 <= d_i <= 9 if !isEqual = false or d[i] if isEqual
}
其中，d_i 表示在 i 位上的数字选择，d[i] 表示 K 在 i 位上的数字

边界条件：
T(0, any, any) = 1，遍历到 0 意味着个位数字都已经确定完了

下面的具体实现与上述推导不符
*/

const MAXL = 1e4+5
const MAXD = 105
const MOD = 1e9+7

var n int
var D int
var d [MAXL]int
var T [MAXL][MAXD]int

func solve(i int, prefix int, isEqual bool) int {
    if i == n {
        if prefix == 0 {
            return 1
        } else {
            return 0
        }
    }
    if !isEqual && T[i][prefix] >= 0 {
        return T[i][prefix]
    }
    maxD := 9
    if isEqual {
        maxD = d[i]
    }
    var sum int64
    for digit := 0; digit <= maxD; digit++ {
        sum += int64(solve(i+1, (prefix+digit) % D, isEqual && digit == d[i]))
    }
    sum %= MOD
    if !isEqual {
        T[i][prefix] = int(sum)
    }
    return int(sum)
}

func main() {
    var K string
    fmt.Scanln(&K)
    fmt.Scanln(&D)
    n = len(K)
    for i := 0; i < n; i++ {
        d[i] = int(K[i] - '0')
    }
    for i := 0; i < n; i++ {
        for j := 0; j < D; j++ {
            T[i][j] = -1
        }
    }
    
    fmt.Println((solve(0, 0, true) - 1 + MOD) % MOD) // minus 1 for 0
}
