// topDigitUnit[i]: 第 i 位数字的单位值，如 topDigitUnit[1] 就是个位数字的单位值
var topDigitUnit = []int{-1, 1, 10, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9}

func topDigit(num, pos int) (int, int) {
    if pos == -1 {
        var digit int
        for pos = 0; num > 0; pos++ {
            digit = num % 10
            num /= 10
        }
        return digit, pos
    } else {
        return num / topDigitUnit[pos] % 10, pos
    }
}

func atMostNGivenDigitSet(D []string, N int) int {
    numD := make([]int, len(D))
    for i := 0; i < len(D); i++ {
        numD[i] = int(D[i][0] - '0')
    }
    fmt.Println(numD)
    
    nDigitCount := make([]int, 10)
    totalCount := make([]int, 10)
    // nDigitCount[i]: i 位数字任意选，可组成的数的数量
    // totalCount[i]: 1...i 位数字任意选，可组成的数的数量
    nDigitCount[1] = len(D)
    totalCount[1] = nDigitCount[1]
    for i := 2; i < 10; i++ {
        nDigitCount[i] = nDigitCount[i-1] * len(D)
        totalCount[i] = totalCount[i-1] + nDigitCount[i]
    }
    nDigitCount[0] = 1 // 特殊情况：不取任何数时，也构成唯一方案，见下：`i * nDigitCount[pos-1]` 时可能被用到
    fmt.Println(totalCount)
    
    digit, pos := topDigit(N, -1)
    result := 0
    if pos > 1 {
        result += totalCount[pos - 1]
    }
    for N > 0 {
        // example: N = 1234, digit = 1, pos = 4
        digit, pos = topDigit(N, pos)
        fmt.Println(digit, pos, result)
        
        var i int
        for i = 0; i < len(numD); i++ {
            if numD[i] < digit {
            } else {
                break
            }
        }
        // i == len(numD) || numD[i] >= digit
        if i == len(numD) {
            // all available digit less than the `digit` of N
            result += nDigitCount[pos]
            break
        } else {
            // `i` available digit less than the `digit` of N
            result += i * nDigitCount[pos-1]
            if numD[i] > digit {
                break
            } else {
                // numD[i] == digit
                N = N - digit * topDigitUnit[pos]
                pos--
            }
        }
    }
    if pos == 0 {
        result ++
    }
    
    return result
}
