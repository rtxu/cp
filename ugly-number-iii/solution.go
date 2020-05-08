
func gcd(a, b int64) int64 {
    if a > b {
        return gcd(b, a)
    }
    if a == 0 {
        return b
    }
    
    return gcd(b % a, a)
}

func lcm(a, b int64) int64 {
    return a * b / gcd(a, b)
}

func max(a, b int64) int64 {
    if a < b {
        return b
    }
    return a
}


func nthUglyNumber(in int, ia int, ib int, ic int) int {
    n, a, b, c := int64(in), int64(ia), int64(ib), int64(ic)
    lcm_ab := lcm(a, b)
    lcm_bc := lcm(b, c)
    lcm_ac := lcm(a, c)
    lcm_abc := lcm(a, lcm(b, c))
    
    // x in [0, lcm_abc)
    calcOrdinal := func(x int64) int64 {
        return x / a + x / b + x / c - x / lcm_ab - x / lcm_bc - x / lcm_ac
    }
    n_lcm_abc := calcOrdinal(lcm_abc) + 1
    
    originX := n / n_lcm_abc * lcm_abc
    offsetN := n % n_lcm_abc
    
    left, right := int64(0), lcm_abc
    for left+1 < right {
        m := (left+right)/2
        ordinal := calcOrdinal(m)
        if ordinal <= offsetN {
            left = m
        } else {
            right = m
        }
    }
    
    result := originX + max(left - left%a, max(left - left%b, left-left%c))

    return int(result)

}
