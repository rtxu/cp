func minOperations(nums []int) int {
    n := len(nums)
    var result int
    var zeroCnt int
    var round int
    for zeroCnt < n {
        zeroCnt = 0
        for i := 0; i < n; i++ {
            if nums[i] % 2 == 1 {
                result++
                nums[i]--
            }
            if nums[i] == 0 {
                zeroCnt++
            }
        }
        if zeroCnt < n {
            result++
            for i := 0; i < n; i++ {
                nums[i] /= 2
            }
        }
        
        round++
        // fmt.Printf("Round #%d: result = %d, zeroCnt = %d, nums = %v\n", round, result, zeroCnt, nums)
    }
    return result
}
