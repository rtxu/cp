
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func nthUglyNumber(n int) int {
    
    ugly := make([]int, n)
    ugly[0] = 1
    
    i2, i3, i5 := 0, 0, 0
    for i := 1; i < n; i++ {
        next := min(ugly[i2]*2, min(ugly[i3]*3, ugly[i5]*5))
        if next == ugly[i2]*2 {
            i2++
        }
        if next == ugly[i3]*3 {
            i3++
        }
        if next == ugly[i5]*5 {
            i5++
        }
        
        ugly[i] = next
    }
    
    return ugly[n-1]
    
}
