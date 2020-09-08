const MOD = 1e9+7

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
    
    for i := 1; i <= n; i++ {
        lastDigit, lastOptionCount, last2OptionCount := 0, 0, 0
        if s[i-1] == '*' {
            lastOptionCount = 9
        } else {
            t := int(s[i-1] - '0')
            if t > 0 {
                lastOptionCount = 1
                lastDigit = t
            } else {
                lastOptionCount = 0
            }
        }
        // lastOptionCount 仅有三个选项：9, 1, 0
        
        if i-2 >= 0 {
            if s[i-2] == '*' {
                if lastOptionCount == 0 {
                    // 10, 20
                    last2OptionCount = 2
                } else if lastOptionCount == 1 {
                    last2OptionCount = 1 // 10 + lastDigit
                    if 20 + lastDigit <= 26 {
                        last2OptionCount++
                    }
                } else {
                    last2OptionCount = 9 + 6 // 11-19, 21-26
                }
                
            } else {
                t := int(s[i-2] - '0')
                if t >= 1 && t <= 2 {
                    if lastOptionCount == 0 {
                        // 10 or 20
                        last2OptionCount = 1
                    } else if lastOptionCount == 1 {
                        if t*10 + lastDigit <= 26 {
                            last2OptionCount++
                        }
                    } else {
                        if t == 1 {
                            last2OptionCount = 9 // 11-19
                        } else {
                            last2OptionCount = 6 // 21-26
                        }
                    }
                }
            }
        }
        
        sum[i] += int(int64(sum[i-1]) * int64(lastOptionCount) % MOD)
        if i >= 2 {
            sum[i] += int(int64(sum[i-2]) * int64(last2OptionCount) % MOD)
        }
        sum[i] %= MOD
    }
    
    return sum[n]
}
