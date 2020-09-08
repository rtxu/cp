func numSubarrayProductLessThanK(nums []int, k int) int {
    if k <= 1 {
        return 0
    }

    prod := 1
    n := len(nums)
    j := 0
    var result int
    for i := 0; i < n; i++ {
        prod *= nums[i]
        if prod >= k {
            // j...i-1 < k
            // fmt.Println(i, j, prod)
            for prod >= k {
                result += i-j
                prod /= nums[j]
                j++
            }
        }
    }
    result += (1+n-j)*(n-j) / 2
    
    return result
}
