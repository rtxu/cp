

func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 {
        if divisor == 1 { 
            return dividend 
        } else if divisor == -1 {
            return math.MaxInt32 
        } else if divisor&1 == 0 {
            return divide(dividend>>1, divisor>>1)
        } else {
            return divide(dividend+1, divisor)
        }
    }
    // sign >= 0，result 为正；sign < 0，result 为负
    sign := (dividend >> 31) ^ (divisor >> 31)
    if dividend >> 31 & 1 == 1 {
        dividend = -dividend
    }
    if divisor >> 31 & 1 == 1 {
        divisor = -divisor
    }
    var i int
    for i = 0; i < 31 && (divisor << i) >> 30 == 0; i++ {}
    var result int
    for ; i >= 0; i-- {
        if (divisor << i) <= dividend {
            result += 1<<i
            dividend -= divisor<<i
        }
    }
    if sign >= 0 {
        return result
    } else {
        return -result
    }
}
