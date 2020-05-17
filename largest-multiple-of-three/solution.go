func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func largestMultipleOfThree(digits []int) string {
    mod := make([][]int, 3)
    mod[0] = make([]int, 0)
    mod[1] = make([]int, 0)
    mod[2] = make([]int, 0)
    
    n := len(digits)
    for i := 0; i < n; i++ {
        remainder := digits[i] % 3
        mod[remainder] = append(mod[remainder], digits[i])
    }
    n1 := len(mod[1])
    n2 := len(mod[2])
    if (n1 % 3 == 0 && n2 % 3 == 2) {
        // n1, n2 mod 3 后，共有 9 种情况，仅有一个为 0 一个为 2 时，
        // 可以通过拆开为 0 数字的三个数与为 2 数字的两个数结合，形成 4 位数
        // 大于单纯选为 0 数字的三个数组成的 3 位数
        if n1 > 0 {
            n1--
        }
    } 
    if (n1 % 3 == 2 && n2 % 3 == 0) {
        if n2 > 0 {
            n2--
        }
    }
    minN := min(n1%3, n2%3)
    n1, n2 = n1/3*3 + minN, n2/3*3 + minN
    
    var resultDigits []int
    // remainder 为 0 的全取
    resultDigits = append(resultDigits, mod[0]...)
    // 取最大的 3x+minN 个
    sort.Ints(mod[1])
    resultDigits = append(resultDigits, mod[1][len(mod[1])-n1:len(mod[1])]...)
    sort.Ints(mod[2])
    resultDigits = append(resultDigits, mod[2][len(mod[2])-n2:len(mod[2])]...)
    
    if len(resultDigits) == 0 {
        return ""
    } else {
        sort.Ints(resultDigits)
        if resultDigits[len(resultDigits)-1] == 0 {
            return "0"
        } else {
            var sb strings.Builder
            for i := len(resultDigits) - 1; i >= 0; i-- {
                sb.WriteByte(byte(resultDigits[i]) + '0')
            }
            return sb.String()
        }
    }
}
