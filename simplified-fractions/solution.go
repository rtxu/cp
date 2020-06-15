func gcd(a, b int) int {
    if a > b {
        return gcd(b, a)
    }
    if b % a == 0 {
        return a
    } else {
        return gcd(a, b % a)
    }
}

func simplifiedFractions(n int) []string {
    var result []string
    for i := 2; i <= n; i++ {
        for j := 1; j < i; j++ {
            if gcd(i, j) == 1 {
                result = append(result, fmt.Sprintf("%d/%d", j, i))
            }
        }
    }
    return result
}
