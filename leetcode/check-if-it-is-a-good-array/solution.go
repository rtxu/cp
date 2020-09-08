func gcd(a, b int) int {
    if a > b {
        return gcd(b, a)
    }
    if a == 0 {
        return b
    }
    return gcd(b % a, a)
}

func isGoodArray(nums []int) bool {
    if len(nums) == 0 {
        return false
    }
    divisor := nums[0]
    for i := 1; i < len(nums); i++ {
        divisor = gcd(divisor, nums[i])
    }
    
    return divisor == 1
}
