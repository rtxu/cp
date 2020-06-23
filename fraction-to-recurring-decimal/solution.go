func abs(n int) int {
    if n >= 0 {
        return n
    } else {
        return -n
    }
}

func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }
    sign := ""
    if numerator >= 0 && denominator < 0 || numerator < 0 && denominator >= 0 {
        sign = "-"
    }
    numerator, denominator = abs(numerator), abs(denominator)
    quotient, remainder := numerator / denominator, numerator % denominator
    var result strings.Builder
    if remainder == 0 {
        return fmt.Sprint(sign, quotient)
    } else {
        result.WriteString(fmt.Sprint(sign, quotient, "."))
    }
    M := map[int]int{}
    M[remainder] = result.Len()  // 每个 remainder 对应的 quotient 在 result 中的起始位置
    for {
        numerator = remainder * 10
        for numerator < denominator {
            numerator *= 10
            result.WriteByte('0')
        }
        next_quotient, next_remainder := numerator / denominator, numerator % denominator
        result.WriteByte(byte('0' + next_quotient))
        
        if next_remainder == 0 {
            return result.String()
        } else {
            if pos, exist := M[next_remainder]; exist {
                s1 := result.String()[:pos]
                s2 := "(" + result.String()[pos:] + ")"
                return s1 + s2
            }
        }
        remainder = next_remainder
        M[remainder] = result.Len()
    }
    return result.String()
}
