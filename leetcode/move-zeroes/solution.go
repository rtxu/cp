func moveZeroes(nums []int)  {
    var count int
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            count++
        } else {
            nums[i-count] = nums[i]
        }
    }
    for i := len(nums)-1; count > 0; i, count = i-1, count-1 {
        nums[i] = 0
    }
}
