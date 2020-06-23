func sortColors(nums []int)  {
    n := len(nums)
    // next place to drop 0/2
    i0, i2 := 0, n-1
    swap := func(i, j int) {
        nums[i], nums[j] = nums[j], nums[i]
    }
    for i := 0; i <= i2; i++ {
        for (i >= i0 && i <= i2) {
            if nums[i] == 0 {
                swap(i, i0)
                i0++
            } else if nums[i] == 1 {
                break
            } else {
                swap(i, i2)
                i2--
            }
        }
    }
}
