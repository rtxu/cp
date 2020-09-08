func singleNumber(nums []int) int {
    var result int
    for i := 0; i < len(nums); i++ {
        result ^= nums[i]
    }
    return result
}
