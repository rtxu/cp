func gcd(a, b int) int {
    if a > b {
        return gcd(b, a)
    }
    
    if a == 0 {
        return b
    }
    
    return gcd(b%a, a)
}

const MOD = 1e9+7

func nthMagicalNumber(N int, A int, B int) int {
    // 1*A 2*A ... x1*A
    // 1*B 2*B ... x2*B
    // 在 1 到 lcm(A, B) 之间不存在可以同时被 A和B 整除的数
    // lcm(A, B)+1 到 2*lcm(A, B) 的情况等同于 1 到 lcm(A, B)，故以 [1, lcm(A,B)] 为循环节
    // 假设 [1, lcm(A,B)] 之间的结果为 count
    // 最终答案为 N / count * lcm(A, B) + 第 N % count 个在 [1, lcm(A,B)] 之间可整除 A 或 B 的数
    
    gcdAB := gcd(A, B)
    lcmAB := A * B / gcdAB
    count := (lcmAB-1) / A + (lcmAB-1) / B + 1
    
    target := N % count
    mA, mB := 1, 1
    current := 0
    for i := 0; i < target; i++ {
        if mA * A < mB * B {
            current = mA * A
            mA++
        } else {
            current = mB * B
            mB++
        }
    }
    result := int(int64(N / count) * int64(lcmAB) % MOD) + current
    result %= MOD
    return result
}
